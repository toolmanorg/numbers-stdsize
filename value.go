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
