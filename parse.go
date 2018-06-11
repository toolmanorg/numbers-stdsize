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
// If the numeric value is unparable, then Parse returns an error of type
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
