// Copyright © 2018 Timothy E. Peoples <eng@toolman.org>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.

package stdsize

// Parse converts the size designation s to an integer Value. A size
// designation is specified as any (real) numeric value, either positive
// or negative, followed by a size modifier.
//
// The following modifiers are supported:
//
//			Decimal     Binary
//			--------    --------
//			K: 10^3     Ki: 2^10
//			M: 10^6     Mi: 2^20
//			G: 10^9     Gi: 2^30
//			T: 10^12    Ti: 2^40
//			P: 10^15    Pi: 2^50
//
// For example, "1K" is 1000, "1Ki" is 1024, and "2.5Ki" is 2048 + 512 or 2560.
//
// If no size modifier is provided then the numeric value must be an integer,
// otherwise Parse returns ErrNotInteger.
//
// If the given size modifier is unrecognized, Parse returns ErrBadModifier.
//
// If the numeric value is unparsable, then Parse returns an error of type
// *strconv.NumError.
func Parse(s string) (Value, error) {
	return parse(s, false)
}

// ParseBinary is similar to Parse except all size modifiers are assumed to be
// binary. In other words, both "K" and "Ki" are interpreted as 2^10.
func ParseBinary(s string) (Value, error) {
	return parse(s, true)
}

func parse(s string, force bool) (Value, error) {
	p, err := dissect(s)
	if err != nil {
		return 0, err
	}

	if force {
		p.bin = true
	}

	return p.value(), nil
}
