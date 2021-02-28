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

package rf

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
)

// Hz represents a specific frequency, in cycles per second.
type Hz float64

// UnmarshalJSON will parse a string as a frequency, and convert it into
// Hz. This can be used to transmit frequency data via JSON.
func (h *Hz) UnmarshalJSON(data []byte) error {
	var el string
	var err error

	if err := json.Unmarshal(data, &el); err != nil {
		return err
	}
	*h, err = ParseHz(el)
	return err
}

// MarshalJSON will convert the frequency in Hz to a string.
// This can be used to transmit frequency data via JSON.
func (h *Hz) MarshalJSON() ([]byte, error) {
	return json.Marshal(h.String())
}

// MarshalYAML will convert the frequency in Hz to a string.
// This can be used to transmit frequency data via YAML.
func (h *Hz) MarshalYAML() (interface{}, error) {
	return h.String(), nil
}

// UnmarshalYAML will parse a string as a frequency, and convert it into
// Hz. This can be used to transmit frequency data via YAML.
func (h *Hz) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var (
		err error
		hz  string
	)
	if err := unmarshal(&hz); err != nil {
		return err
	}
	*h, err = ParseHz(hz)
	return err
}

var (
	// KHz represents one kilohertz, or 1,000 Hz
	KHz Hz = Hz(1e+3)

	// MHz represents one megahertz, or 1,000,000 Hz
	MHz Hz = Hz(1e+6)

	// GHz represents one gigahertz, or 1,000,000,000 Hz
	GHz Hz = Hz(1e+9)

	// THz represents one terrahertz, or 1,000,000,000,000 Hz
	THz Hz = Hz(1e+12)

	// KHzBand represents the Kilohertz band, from 1KHz up to 1MHz.
	KHzBand Allocation = Allocation{Name: "KHz", Range: Range{KHz, MHz - 1}}

	// MHzBand represents the Megahertz band, from 1MHz up to 1GHz.
	MHzBand Allocation = Allocation{Name: "MHz", Range: Range{MHz, GHz - 1}}

	// GHzBand represents the Gigahertz band, from 1GHz up to 1THz.
	GHzBand Allocation = Allocation{Name: "GHz", Range: Range{GHz, THz - 1}}

	// SIBands represents the Hz-based allocations (KHz, MHz, GHz)
	SIBands Allocations = Allocations{KHzBand, MHzBand, GHzBand}
)

// String will convert the frequency into a string, able to be re-parsed as
// a frequency, or displayed to a user.
func (h Hz) String() string {
	var (
		frequency float64 = float64(h)
		fkhz      float64 = float64(KHz)
		steps     uint    = 0
		neg       bool    = h < 0
		sign      string  = ""
	)

	if neg {
		frequency = -frequency
		sign = "-"
	}

	names := []string{"Hz", "KHz", "MHz", "GHz", "THz"}

	for frequency > fkhz {
		frequency = frequency / fkhz
		steps++
	}

	return fmt.Sprintf(
		"%s%s%s",
		sign,
		strconv.FormatFloat(frequency, 'f', -1, 64),
		names[steps],
	)
}

// SIBandName will return the name of the SI frequency range (KHz, MHz, GHz)
func (h Hz) SIBandName() string {
	for _, band := range SIBands.ContainingFrequency(h) {
		return band.Name
	}
	return ""
}

// MustParseHz will run the string through ParseHz, and on error, panic. This
// is very useful for hardcoded const strings, or places where invalid
// input is actually fatal.
func MustParseHz(freq string) Hz {
	hz, err := ParseHz(freq)
	if err != nil {
		panic(err)
	}
	return hz
}

// ParseHz will take a frequency as a string, and return it as an rf.Hz.
//
// Examples of valid frequencies:
//
// -10MHz
// 2GHz
// 2000Hz
//
// Valid Hz units are 'Hz', 'KHz', 'MHz', 'GHz', 'THz'
func ParseHz(freq string) (Hz, error) {
	pattern := "(?P<sign>[-+])?((?P<freq>[0-9]*(\\.[0-9]*)?)(?P<unit>[A-Za-z]+))+"

	r := regexp.MustCompile(pattern)

	parts := map[string]string{}

	values := r.FindStringSubmatch(freq)
	keys := r.SubexpNames()

	if len(values) != len(keys) {
		return Hz(0), fmt.Errorf("rf: invalid frequency: %s", freq)
	}

	for i, key := range keys {
		parts[key] = values[i]
	}

	return parseHzFromParts(parts)
}

// While this looks hacky, looking at other similar code, for instance,
// time.Parse, this is a bit more maintainable for a single person. The
// downsides in doing this are worth it for my very specific use-case.
func parseHzFromParts(parts map[string]string) (Hz, error) {
	var scale Hz = 0

	switch parts["unit"] {
	case "Hz", "hz":
		scale = Hz(1)
		break
	case "KHz", "khz", "kHz":
		scale = KHz
		break
	case "MHz", "mhz":
		scale = MHz
		break
	case "GHz", "ghz":
		scale = GHz
		break
	case "THz", "thz":
		scale = THz
		break
	default:
		return Hz(0), fmt.Errorf("rf: unknown unit: %s", parts["unit"])
	}

	fscale := float64(scale)
	ffreq, err := strconv.ParseFloat(parts["freq"], 64)
	if err != nil {
		return Hz(0), err
	}

	ffreqInHz := fscale * ffreq
	switch parts["sign"] {
	case "+", "":
		break
	case "-":
		ffreqInHz = -ffreqInHz
	default:
		return Hz(0), fmt.Errorf("rf: Unknown prefix: %s", parts["sign"])
	}

	freqInHz := Hz(int64(ffreqInHz))

	return freqInHz, nil
}

// vim: foldmethod=marker
