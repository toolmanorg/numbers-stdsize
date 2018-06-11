// Copyright © 2018 Timothy E. Peoples <eng@toolman.org>
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; either version 2
// of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package stdsize

import (
	"reflect"
	"strconv"
	"testing"
)

type parseTestcase struct {
	in   string
	want Value
	err  error
}

func (tc *parseTestcase) testParse(t *testing.T) {
	if got, err := Parse(tc.in); !reflect.DeepEqual(err, tc.err) || got != tc.want {
		t.Errorf("Parse(%q) == (%d, %v); Wanted (%d, %v)", tc.in, got, err, tc.want, tc.err)
	}
}

func (tc *parseTestcase) testParseBinary(t *testing.T) {
	if got, err := ParseBinary(tc.in); !reflect.DeepEqual(err, tc.err) || got != tc.want {
		t.Errorf("Parse(%q) == (%d, %v); Wanted (%d, %v)", tc.in, got, err, tc.want, tc.err)
	}
}

func TestParse(t *testing.T) {
	_, perr := strconv.ParseFloat("NO", 64)

	cases := []*parseTestcase{
		{"", 0, nil},
		{"  ", 0, nil},
		{"0M", 0, nil},
		{"123", 123, nil},
		{"2345", 2345, nil},

		{"234.5", 0, ErrNotInteger},
		{"234.5X", 0, ErrBadModifier},
		{"234.5KB", 0, ErrBadModifier},

		{"BAD", 0, ErrBadModifier},
		{"NOT", 0, perr},

		{".3K", 300, nil},
		{".5K", 500, nil},
		{"0.5K", 500, nil},
		{"0.75K", 750, nil},

		{".3Ki", 307, nil},
		{".33Ki", 338, nil},
		{".333Ki", 341, nil},
		{".33333333Ki", 341, nil},

		{".5Ki", 512, nil},
		{"0.5Ki", 512, nil},
		{"0.75Ki", 768, nil},

		{"2K", 2000, nil},
		{"2Ki", 2048, nil},
		{"2.3K", 2300, nil},
		{"2.5Ki", 2560, nil},
		{"2.7Ki", 2765, nil},

		{"2M", 2000000, nil},
		{"2Mi", 2097152, nil},
		{"2.3M", 2300000, nil},
		{"2.5Mi", 2621440, nil},
		{"2.7Mi", 2831155, nil},

		{"2G", 2000000000, nil},
		{"2Gi", 2147483648, nil},
		{"2.3G", 2300000000, nil},
		{"2.5Gi", 2684354560, nil},
		{"2.7Gi", 2899102925, nil},

		{"2T", 2000000000000, nil},
		{"2Ti", 2199023255552, nil},
		{"2.3T", 2300000000000, nil},
		{"2.5Ti", 2748779069440, nil},
		{"2.7Ti", 2968681394995, nil},

		{"2P", 2000000000000000, nil},
		{"2Pi", 2251799813685248, nil},
		{"2.3P", 2300000000000000, nil},
		{"2.5Pi", 2814749767106560, nil},
		{"2.7Pi", 3039929748475085, nil},
	}

	for _, tc := range cases {
		t.Run(tc.in, tc.testParse)
	}

	// p := v.parts(true)
	// t.Logf("p: %+v", p)
}

func TestParseBinary(t *testing.T) {
	cases := []*parseTestcase{
		{"2K", 2048, nil},
		{"2Ki", 2048, nil},
		{"2.5K", 2560, nil},
		{"2.5Ki", 2560, nil},

		{"2M", 2097152, nil},
		{"2Mi", 2097152, nil},
		{"2.5M", 2621440, nil},
		{"2.5Mi", 2621440, nil},

		{"2G", 2147483648, nil},
		{"2Gi", 2147483648, nil},
		{"2.5G", 2684354560, nil},
		{"2.5Gi", 2684354560, nil},

		{"2T", 2199023255552, nil},
		{"2Ti", 2199023255552, nil},
		{"2.5T", 2748779069440, nil},
		{"2.5Ti", 2748779069440, nil},

		{"2P", 2251799813685248, nil},
		{"2Pi", 2251799813685248, nil},
		{"2.5P", 2814749767106560, nil},
		{"2.5Pi", 2814749767106560, nil},
	}

	for _, tc := range cases {
		t.Run(tc.in, tc.testParseBinary)
	}
}
