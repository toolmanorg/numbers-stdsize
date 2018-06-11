package stdsize

import (
	"fmt"
	"strconv"
	"strings"
)

// Format implements fmt.Formatter to provide custom, printf style formatting
// for Values. It is not intended to be called directly.
func (v Value) Format(f fmt.State, c rune) {
	var cust, quot bool
	switch c {
	case 'q':
		quot = true
		fallthrough
	case 'S', 'v':
		cust = true
	}

	var (
		fs  = "%"
		bin = true
	)

	for _, fl := range "+-# 0" {
		if !f.Flag(int(fl)) {
			continue
		}

		switch fl {
		case '#':
			if cust {
				bin = false
				break
			}
			fallthrough

		default:
			fs += string(fl)
		}
	}

	var wp bool

	if w, ok := f.Width(); ok {
		fs += fmt.Sprintf("%d", w)
		wp = true
	}

	if p, ok := f.Precision(); ok {
		fs += fmt.Sprintf(".%d", p)
		wp = true
	}

	if !cust {
		cs := string(c)
		if strings.IndexAny(cs, "eEfFgG") == -1 {
			fmt.Fprintf(f, fs+cs, int64(v))
		} else {
			fmt.Fprintf(f, fs+cs, float64(v))
		}
		return
	}

	fs += "f"

	pv, ps := v.parts(bin).factor()

	var rv string
	if wp {
		rv = fmt.Sprintf(fs+ps, pv)
	} else {
		rv = strings.TrimRight(fmt.Sprintf(fs, pv), "0")
		if li := len(rv) - 1; rv[li] == '.' {
			rv = rv[:li]
		}
		rv += ps
	}

	if quot {
		rv = strconv.Quote(rv)
	}

	fmt.Fprint(f, rv)
}
