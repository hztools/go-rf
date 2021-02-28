// {{{ Copyright (c) Paul R. Tagliamonte <paul@kc3nwj.com>, 2020
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
	"strings"
	"testing"

	"kc3nwj.com/rf"
)

func TestHzParse(t *testing.T) {
	frequency, err := rf.ParseHz("144.39MHz")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	if frequency != rf.Hz(144390000) {
		t.Log("Frequency is mis-calculated")
		t.FailNow()
	}
}

func TestHzParseNeg(t *testing.T) {
	frequency, err := rf.ParseHz("-10Hz")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	if frequency != rf.Hz(-10) {
		t.Log("Frequency isn't being parsed correctly")
		t.FailNow()
	}
}

func TestRange(t *testing.T) {
	frequency := rf.MustParseHz("144.39MHz")

	if !rf.VHFBand.Range.ContainsFrequency(frequency) {
		t.Log("Frequency thinks it's outside of VHF range!")
		t.FailNow()
	}

	vhf := rf.ITUBands.ContainingFrequency(frequency)
	if len(vhf) != 1 {
		t.Log("ITU lookup didn't return VHF")
		t.FailNow()
	}

	if strings.Compare(vhf[0].Name, "VHF") != 0 {
		t.Log("ITU lookup returned something other than VHF")
		t.FailNow()
	}
}

func TestITUBandName(t *testing.T) {
	frequency := rf.MustParseHz("144.39MHz")

	if strings.Compare(frequency.ITUBandName(), "VHF") != 0 {
		t.Log("ITU lookup returned something other than VHF")
		t.FailNow()
	}
}

func TestSIBandName(t *testing.T) {
	frequency := rf.MustParseHz("144.39MHz")

	if strings.Compare(frequency.SIBandName(), "MHz") != 0 {
		t.Log("SI lookup returned something other than MHz")
		t.FailNow()
	}
}

func TestFreqWavelength(t *testing.T) {
	frequency := rf.MustParseHz("144.39MHz")

	if frequency.Wavelength() != 2.076268841332502 {
		t.Log("Wavelength doesn't match what it should")
		t.FailNow()
	}
}

func TestFreqName(t *testing.T) {
	frequency := rf.MustParseHz("144.39MHz")

	if strings.Compare(frequency.String(), "144.39MHz") != 0 {
		t.Log("Stringification went very poorly")
		t.FailNow()
	}

	frequency = rf.MustParseHz("1440.39MHz")

	if strings.Compare(frequency.String(), "1.44039GHz") != 0 {
		t.Log("Stringification went very poorly")
		t.FailNow()
	}
}

func TestFreqMath(t *testing.T) {
	frequency := rf.MustParseHz("144.39MHz")

	r := rf.Range{-rf.KHz * 3, rf.KHz * 3}

	bandwidth := r.Add(frequency)

	if !(rf.Range{rf.Hz(144387000), rf.Hz(144393000)}.Equal(bandwidth)) {
		t.Log("Bandwidth isn't aligning")
		t.FailNow()
	}
}

func TestRangeMath(t *testing.T) {
	threeKHz := rf.Range{-rf.KHz * 3, rf.KHz * 3}
	twoKHz := rf.Range{-rf.KHz * 2, rf.KHz * 2}

	if twoKHz.ContainsRange(threeKHz) {
		t.Log("Two contains three!")
		t.FailNow()
	}

	if !threeKHz.ContainsRange(twoKHz) {
		t.Log("Three doesn't contain two!")
		t.FailNow()
	}

	if !threeKHz.Overlaps(twoKHz) {
		t.Log("Three isn't overlapping with two")
		t.FailNow()
	}

	if !twoKHz.Overlaps(threeKHz) {
		t.Log("Two isn't overlapping with three")
		t.FailNow()
	}

	left := rf.Range{rf.Hz(0), rf.Hz(199)}
	right := rf.Range{rf.Hz(200), rf.Hz(300)}
	middle := rf.Range{rf.Hz(100), rf.Hz(250)}

	if left.Overlaps(right) {
		t.Log("Disjointed range is overlapping")
		t.FailNow()
	}

	if !left.Overlaps(middle) {
		t.Log("Left isn't overlapping with middle")
		t.FailNow()
	}

	if !middle.Overlaps(left) {
		t.Log("Middle isn't overlapping with left")
		t.FailNow()
	}
}

func TestRangeMathCenter(t *testing.T) {
	center := rf.Range{100, 200}.Center()
	if center != rf.Hz(150) {
		t.Log("Middle isn't middle")
		t.FailNow()
	}
}

// vim: foldmethod=marker
