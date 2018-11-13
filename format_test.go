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

import (
	"fmt"
	"strconv"
	"testing"
)

type outs map[string]string

type fmtTestcase struct {
	in  Value
	out outs
}

func (tc *fmtTestcase) test(t *testing.T) {
	for fs, want := range tc.out {
		n := fmt.Sprintf("`%s`=>`%s`", fs, want)
		t.Run(n, func(t *testing.T) {
			if got := fmt.Sprintf(fs, tc.in); got != want {
				t.Errorf("fmt.Sprintf(%q, %d) => %q; Wanted %q", fs, int64(tc.in), got, want)
			}
		})
	}
}

func (tc *fmtTestcase) name() string {
	return strconv.Itoa(int(tc.in))
}

func TestFormat(t *testing.T) {
	cases := []*fmtTestcase{
		{2485125, outs{
			"%v":      "2.37Mi",
			"%q":      `"2.37Mi"`,
			"%S":      "2.37Mi",
			"%d":      "2485125",
			"%#d":     "2485125",
			"%.1f":    "2485125.0",
			"%#.2S":   "2.49M",
			"%#.2q":   `"2.49M"`,
			"%#6.2S":  "  2.49M",
			"%#06.2S": "002.49M",
		}},

		{3500000, outs{
			"%#v":   "3.5M",
			"%#S":   "3.5M",
			"%#.2S": "3.50M",
			"%#d":   "3500000",
			"%d":    "3500000",
			"%S":    "3.33786Mi",
			"%.2S":  "3.34Mi",
			"%+.2S": "+3.34Mi",
			"% .2S": " 3.34Mi",
		}},

		{-3500000, outs{
			"%+.2S": "-3.34Mi",
			"% .2S": "-3.34Mi",
			"%+.2q": `"-3.34Mi"`,
		}},

		{3145728, outs{"%v": "3Mi"}},
		{3153068, outs{"%v": "3.007Mi"}},
		{3153579, outs{"%v": "3.007487Mi"}},
		{3153579, outs{"%.9v": "3.007487297Mi"}},
		{3153592, outs{"%v": "3.0075Mi"}},
		{3154117, outs{"%v": "3.008Mi"}},
	}

	for _, tc := range cases {
		t.Run(tc.name(), tc.test)
	}
}
