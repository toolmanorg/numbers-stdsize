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
