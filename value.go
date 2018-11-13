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

/*
Package stdsize provides the integer type "Value" that may be created
or displayed using human-readable strings.  For example, "2Ki" gets
parsed as 2048 while the numeric value of 5*1000*1000 (i.e. 5 million)
is displayed as "5M".
*/
package stdsize // import "toolman.org/numbers/stdsize"

// Value is a standard size value as returned by Parse or ParseBinary. It also
// provides a custom formatter for displaying human-readable values.
type Value int64

// Constants for common decimal (i.e. power of 10) values.
const (
	Kilo Value = 1000        // 10^3
	Mega Value = Kilo * Kilo // 10^6
	Giga Value = Mega * Kilo // 10^9
	Tera Value = Giga * Kilo // 10^12
	Peta Value = Tera * Kilo // 10^15
)

// Constants for common binary (i.e. power of 2) values.
const (
	Kibi Value = 1 << (10 * (iota + 1)) // 2^10
	Mibi                                // 2^20
	Gibi                                // 2^30
	Tebi                                // 2^40
	Pebi                                // 2^50
)

func (v Value) parts(bin bool) *parts {
	u := units(bin)
	f := float64(v)

	var n bool

	if f < 0 {
		n = true
		f *= -1
	}

	for _, c := range []rune("PTGMK") {
		cf := float64(u[c])
		if f > cf {
			return &parts{n, f / cf, c, bin}
		}
	}

	return &parts{n, f, 0, bin}
}
