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
	"errors"
	"strconv"
	"strings"
	"unicode"
)

type parts struct {
	neg bool
	val float64
	suf rune
	bin bool
}

var (
	// ErrBadModifier is returned by Parse and ParseBinary if the
	// given size modifier is unrecognized.
	ErrBadModifier = errors.New("invalid size modifier")

	// ErrNotInteger is returned by Parse and ParseBinary if no
	// size modifier is provided and the given numeric value is
	// not an integer.
	ErrNotInteger = errors.New("bare size values must be integers")
)

func dissect(s string) (*parts, error) {
	s = strings.TrimSpace(s)

	var (
		p  = new(parts)
		li = len(s) - 1
	)

	if li < 0 {
		return p, nil
	}

	if s[li:] == "i" {
		p.bin = true
		li--
	}

	sr := rune(s[li])

	if unicode.IsDigit(sr) {
		li++
	} else {
		if !strings.ContainsRune("KMGTP", sr) {
			return nil, ErrBadModifier
		}
		p.suf = sr
	}

	v, err := strconv.ParseFloat(s[:li], 64)
	if err != nil {
		return nil, err
	}

	if v < 0 {
		p.neg = true
		v *= -1
	}

	if p.suf == 0 && v != float64(int64(v)) {
		return nil, ErrNotInteger
	}

	p.val = v

	return p, nil
}

func (p *parts) value() Value {
	v, s := p.factor()

	if s == "" {
		return Value(p.val)
	}

	return Value(v*float64(units(p.bin)[rune(s[0])]) + 0.5)
}

func (p *parts) factor() (float64, string) {
	var (
		v float64
		s string
	)

	v = p.val
	if p.neg {
		v *= -1
	}

	if p.suf != 0 {
		s = string(p.suf)
	}

	if s != "" && p.bin {
		s += "i"
	}

	return v, s
}
