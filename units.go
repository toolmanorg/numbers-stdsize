package stdsize

var (
	binUnits = map[rune]Value{
		'K': Kibi,
		'M': Mibi,
		'G': Gibi,
		'T': Tebi,
		'P': Pebi,
	}

	decUnits = map[rune]Value{
		'K': Kilo,
		'M': Mega,
		'G': Giga,
		'T': Tera,
		'P': Peta,
	}
)

func units(bin bool) map[rune]Value {
	if bin {
		return binUnits
	}
	return decUnits
}
