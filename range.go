// {{{ Copyright (c) Paul R. Tagliamonte <paul@k3xec.com>, 2020
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

// Range is a tuple of Hz values, Range[0] is the lowest value of the range, and
// Range[1] is the highest value of the range.
type Range [2]Hz

// String will turn the range into a string.
func (r Range) String() string {
	return fmt.Sprintf("%s->%s", r[0].String(), r[1].String())
}

// ContainsFrequency will check to see if a given Frequency is contained inside
// this Range.
func (r Range) ContainsFrequency(freq Hz) bool {
	return freq >= r[0] && freq <= r[1]
}

// Add the provided frequency in Hz to both the lower and upper side
// of the Range, shifting the Range by the provided Hz.
func (r Range) Add(freq Hz) Range {
	return Range{r[0] + freq, r[1] + freq}
}

// ContainsRange will return true if r1 a subset of the range defined by the
// Range.
func (r Range) ContainsRange(r1 Range) bool {
	return r1[0] >= r[0] && r1[1] <= r[1]
}

// Overlaps will return true if r1 overlaps with the Range.
func (r Range) Overlaps(r1 Range) bool {
	return r[0] <= r1[1] && r[1] >= r1[0]
}

// Equal will check to see if the Range is specifically the same as another
// Range.
func (r Range) Equal(r1 Range) bool {
	return r[0] == r1[0] && r[1] == r1[1]
}

// Intersection will return the intersection of the range this method is
// bound to and the provided range.
func (r Range) Intersection(r1 Range) Range {
	low := r[0]
	high := r[1]

	if r1[0] > low {
		low = r1[0]
	}

	if r1[1] < high {
		high = r1[1]
	}

	if low >= high {
		return Range{Hz(0), Hz(0)}
	}

	return Range{low, high}
}

// Center will return the center of a range (perhaps to get the center of a
// channel to tune to).
func (r Range) Center() Hz {
	return (r[1] + r[0]) / 2
}

// vim: foldmethod=marker
