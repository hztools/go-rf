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

var ituELF Hz = Hz(3)
var ituSLF Hz = Hz(3e+1)
var ituULF Hz = Hz(3e+2)
var ituVLF Hz = Hz(3e+3)
var ituLF Hz = Hz(3e+4)
var ituMF Hz = Hz(3e+5)
var ituHF Hz = Hz(3e+6)
var ituVHF Hz = Hz(3e+7)
var ituUHF Hz = Hz(3e+8)
var ituSHF Hz = Hz(3e+9)
var ituEHF Hz = Hz(3e+10)
var ituTHF Hz = Hz(3e+11)

var (
	// ELFBand or Extremely Low Frequency, is a slice of RF space defined by the ITU
	// as being between 3Hz and 30Hz
	ELFBand Allocation = Allocation{Name: "ELF", Range: Range{ituELF, ituSLF - 1}}

	// SLFBand or Super Low Frequency, is a slice of RF space defined by the ITU
	// as being between 30Hz and 300Hz
	SLFBand Allocation = Allocation{Name: "SLF", Range: Range{ituSLF, ituULF - 1}}

	// ULFBand or Ultra Low Frequency, is a slice of RF space defined by the ITU
	// as being between 300Hz and 3KHz
	ULFBand Allocation = Allocation{Name: "ULF", Range: Range{ituULF, ituVLF - 1}}

	// VLFBand or Very Low Frequency, is a slice of RF space defined by the ITU
	// as being between 3KHz and 30KHz
	VLFBand Allocation = Allocation{Name: "VLF", Range: Range{ituVLF, ituLF - 1}}

	// LFBand or Low Frequency, is a slice of RF space defined by the ITU
	// as being between 30KHz and 300KHz
	LFBand Allocation = Allocation{Name: "LF", Range: Range{ituLF, ituMF - 1}}

	// MFBand or Medium Frequency, is a slice of RF space defined by the ITU
	// as being between 300KHz and 3MHz
	MFBand Allocation = Allocation{Name: "MF", Range: Range{ituMF, ituHF - 1}}

	// HFBand or High Frequency, is a slice of RF space defined by the ITU
	// as being between 3MHz and 30MHz
	HFBand Allocation = Allocation{Name: "HF", Range: Range{ituHF, ituVHF - 1}}

	// VHFBand or Very High Frequency, is a slice of RF space defined by the ITU
	// as being between 30MHz and 300MHz
	VHFBand Allocation = Allocation{Name: "VHF", Range: Range{ituVHF, ituUHF - 1}}

	// UHFBand or Ultra High Frequency, is a slice of RF space defined by the ITU
	// as being between 300MHz and 3GHz
	UHFBand Allocation = Allocation{Name: "UHF", Range: Range{ituUHF, ituSHF - 1}}

	// SHFBand or Super High Frequency, is a slice of RF space defined by the ITU
	// as being between 3GHz and 30GHz
	SHFBand Allocation = Allocation{Name: "SHF", Range: Range{ituSHF, ituEHF - 1}}

	// EHFBand or Extremely High Frequency, is a slice of RF space defined by the ITU
	// as being between 30GHz and 300GHz
	EHFBand Allocation = Allocation{Name: "EHF", Range: Range{ituEHF, ituTHF - 1}}

	// ITUBands represents all the ITU allocated RF bands.
	//
	// This is likely most useful to amateur radio applications, where specific
	// individuals are using the ITU names frequently.
	ITUBands Allocations = []Allocation{
		ELFBand, SLFBand, ULFBand, VLFBand,
		LFBand, MFBand, HFBand,
		VHFBand, UHFBand, SHFBand, EHFBand,
	}
)

// ITUBandName will return the string verify of the ITU band name the frequency
// is contained in.
func (h Hz) ITUBandName() string {
	for _, band := range ITUBands.ContainingFrequency(h) {
		return band.Name
	}
	return ""
}

// vim: foldmethod=marker
