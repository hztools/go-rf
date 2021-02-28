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
	"fmt"
)

// Allocation is a range of Frequency, allocated a name,
// and perhaps a purpose. Some examples of this would be
// the 'KU' radar band, 'VHF' range or 'WiFi Channel 11'.
type Allocation struct {
	// Name describing the band
	Name string

	// Range of frequency that this Allocation covers
	Range Range
}

// String will output a human readable string representing the Allocation of
// frequency.
func (r Allocation) String() string {
	return fmt.Sprintf("name=%s, range=%s", r.Name, r.Range)
}

// Allocations is a slice that represents grouped frequency allocations, which
// allow for easy querying.
type Allocations []Allocation

// ContainingFrequency will return all Allocations that contain this Frequency.
func (a Allocations) ContainingFrequency(freq Hz) Allocations {
	ret := Allocations{}
	for _, allocation := range a {
		if allocation.Range.ContainsFrequency(freq) {
			ret = append(ret, allocation)
		}
	}
	return ret
}

// First will return the first Allocation in the Allocations slice.
func (a Allocations) First() Allocation {
	for _, el := range a {
		return el
	}
	return Allocation{}
}

// vim: foldmethod=marker
