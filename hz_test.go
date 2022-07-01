// {{{ Copyright (c) Paul R. Tagliamonte <paul@k3xec.com>, 2021
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE. }}}

package rf_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"hz.tools/rf"
)

func TestHzParse(t *testing.T) {
	frequency, err := rf.ParseHz("144.39MHz")
	assert.NoError(t, err)
	assert.Equal(t, rf.Hz(144390000), frequency)
}

func TestHzParseCase(t *testing.T) {
	for _, freq := range []string{
		"10Hz", "10hz",
		"100KHz", "10kHz", "10khz",
		"100MHz", "100mhz",
		"2.5GHz", "5ghz",
		"10THz", "100thz",
	} {
		_, err := rf.ParseHz(freq)
		assert.NoError(t, err)
	}
}

func TestHzParseNeg(t *testing.T) {
	frequency, err := rf.ParseHz("-10Hz")
	assert.NoError(t, err)
	assert.Equal(t, rf.Hz(-10), frequency)
}

func TestHzParseEmpty(t *testing.T) {
	_, err := rf.ParseHz("")
	assert.Error(t, err)
}

func TestRange(t *testing.T) {
	frequency := rf.MustParseHz("144.39MHz")
	assert.True(t, rf.VHFBand.Range.ContainsFrequency(frequency))
	vhf := rf.ITUBands.ContainingFrequency(frequency)
	assert.Equal(t, 1, len(vhf))
	assert.Equal(t, "VHF", vhf[0].Name)
}

func TestITUBandName(t *testing.T) {
	frequency := rf.MustParseHz("144.39MHz")
	assert.Equal(t, "VHF", frequency.ITUBandName())
}

func TestSIBandName(t *testing.T) {
	frequency := rf.MustParseHz("144.39MHz")
	assert.Equal(t, "MHz", frequency.SIBandName())
}

func TestFreqWavelength(t *testing.T) {
	frequency := rf.MustParseHz("144.39MHz")
	assert.Equal(t, 2.076268841332502, frequency.Wavelength())
}

func TestFreqName(t *testing.T) {
	frequency := rf.MustParseHz("144.39MHz")
	assert.Equal(t, "144.39MHz", frequency.String())
	frequency = rf.MustParseHz("1440.39MHz")
	assert.Equal(t, "1.44039GHz", frequency.String())
	assert.Equal(t, "-1.44039GHz", (-frequency).String())

	frequency = rf.MustParseHz("10KHz")
	assert.Equal(t, "10kHz", frequency.String())
}

func FuzzParseHz(f *testing.F) {
	f.Add("144.39MHz")
	f.Add("145.39kHz")
	f.Add("144Hz")
	f.Add("10GHz")
	f.Add("-1GHz")
	f.Fuzz(func(t *testing.T, f string) {
		// We want panics
		rf.ParseHz(f)
	})
}

func TestFreqMath(t *testing.T) {
	frequency := rf.MustParseHz("144.39MHz")
	r := rf.Range{-rf.KHz * 3, rf.KHz * 3}
	bandwidth := r.Add(frequency)
	assert.True(t, rf.Range{rf.Hz(144387000), rf.Hz(144393000)}.Equal(bandwidth))
}

func TestRangeMath(t *testing.T) {
	threeKHz := rf.Range{-rf.KHz * 3, rf.KHz * 3}
	twoKHz := rf.Range{-rf.KHz * 2, rf.KHz * 2}

	assert.False(t, twoKHz.ContainsRange(threeKHz))
	assert.True(t, threeKHz.ContainsRange(twoKHz))
	assert.True(t, threeKHz.Overlaps(twoKHz))
	assert.True(t, twoKHz.Overlaps(threeKHz))

	left := rf.Range{rf.Hz(0), rf.Hz(199)}
	right := rf.Range{rf.Hz(200), rf.Hz(300)}
	middle := rf.Range{rf.Hz(100), rf.Hz(250)}

	assert.False(t, left.Overlaps(right))
	assert.True(t, left.Overlaps(middle))
	assert.True(t, middle.Overlaps(left))
}

func TestRangeMathCenter(t *testing.T) {
	assert.Equal(t, rf.Hz(150), rf.Range{100, 200}.Center())
}

func TestRangeIntersection(t *testing.T) {
	left := rf.Range{rf.Hz(0), rf.Hz(199)}
	right := rf.Range{rf.Hz(200), rf.Hz(300)}
	middle := rf.Range{rf.Hz(100), rf.Hz(250)}
	assert.Equal(t, rf.Range{100, 199}, middle.Intersection(left))
	assert.Equal(t, rf.Range{200, 250}, middle.Intersection(right))
	assert.Equal(t, rf.Range{0, 0}, left.Intersection(right))
}

// vim: foldmethod=marker
