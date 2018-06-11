package stdsize // import "toolman.org/numbers/stdsize"

// Value is a standard size value as returned by Parse or ParseBinary. It also
// provides a custom formatter for displaying human-readable values.
type Value int64

// Constants for common decimal (i.e. power of 10) values.
const (
	Kilo Value = 1000
	Mega Value = Kilo * Kilo
	Giga Value = Mega * Kilo
	Tera Value = Giga * Kilo
	Peta Value = Tera * Kilo
)

// Constants for common binary (i.e. power of 2) values.
const (
	Kibi Value = 1 << (10 * (iota + 1))
	Mibi
	Gibi
	Tebi
	Pebi
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
