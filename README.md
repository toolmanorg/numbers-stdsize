
[![GoDoc](https://godoc.org/toolman.org/numbers/stdsize?status.svg)](https://godoc.org/toolman.org/numbers/stdsize) [![Go Report Card](https://goreportcard.com/badge/toolman.org/numbers/stdsize)](https://goreportcard.com/report/toolman.org/numbers/stdsize) [![Build Status](https://travis-ci.org/toolmanorg/numbers-stdsize.svg?branch=master)](https://travis-ci.org/toolmanorg/numbers-stdsize)

# stdsize
`import "toolman.org/numbers/stdsize"`

* [Overview](#pkg-overview)
* [Install](#pkg-install)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package stdsize provides the integer type "Value" that may be created
or displayed using human-readable strings.  For example, "2Ki" gets
parsed as 2048 while the numeric value of 5\*1000\*1000 (i.e. 5 million)
is displayed as "5M".

## <a name="pkg-install">Install</a>

``` sh
  go get toolman.org/numbers/stdsize
```

## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [type Value](#Value)
  * [func Parse(s string) (Value, error)](#Parse)
  * [func ParseBinary(s string) (Value, error)](#ParseBinary)
  * [func (v Value) Format(f fmt.State, c rune)](#Value.Format)


#### <a name="pkg-files">Package files</a>
[format.go](/src/toolman.org/numbers/stdsize/format.go) [parse.go](/src/toolman.org/numbers/stdsize/parse.go) [parts.go](/src/toolman.org/numbers/stdsize/parts.go) [units.go](/src/toolman.org/numbers/stdsize/units.go) [value.go](/src/toolman.org/numbers/stdsize/value.go) 


## <a name="pkg-variables">Variables</a>
``` go
var (
    // ErrBadModifier is returned by Parse and ParseBinary if the
    // given size modifier is unrecognized.
    ErrBadModifier = errors.New("invalid size modifier")

    // ErrNotInteger is returned by Parse and ParseBinary if no
    // size modifier is provided and the given numeric value is
    // not an integer.
    ErrNotInteger = errors.New("bare size values must be integers")
)
```

## <a name="Value">type</a> [Value](/src/target/value.go?s=1156:1172#L16)
``` go
type Value int64
```
Value is a standard size value as returned by Parse or ParseBinary. It also
provides a custom formatter for displaying human-readable values.


``` go
const (
    Kilo Value = 1000        // 10^3
    Mega Value = Kilo * Kilo // 10^6
    Giga Value = Mega * Kilo // 10^9
    Tera Value = Giga * Kilo // 10^12
    Peta Value = Tera * Kilo // 10^15
)
```
Constants for common decimal (i.e. power of 10) values.


``` go
const (
    Kibi Value = 1 << (10 * (iota + 1)) // 2^10
    Mibi                                // 2^20
    Gibi                                // 2^30
    Tebi                                // 2^40
    Pebi                                // 2^50
)
```
Constants for common binary (i.e. power of 2) values.

### <a name="Parse">func</a> [Parse](/src/target/parse.go?s=1531:1566#L31)
``` go
func Parse(s string) (Value, error)
```
Parse converts the size designation s to an integer Value. A size
designation is specified as any (real) numeric value, either positive
or negative, followed by a size modifier.

The following modifiers are supported:


	Decimal     Binary
	--------    --------
	K: 10^3     Ki: 2^10
	M: 10^6     Mi: 2^20
	G: 10^9     Gi: 2^30
	T: 10^12    Ti: 2^40
	P: 10^15    Pi: 2^50

For example, "1K" is 1000, "1Ki" is 1024, and "2.5Ki" is 2048 + 512 or 2560.

If no size modifier is provided then the numeric value must be an integer,
otherwise Parse returns ErrNotInteger.

If the given size modifier is unrecognized, Parse returns ErrBadModifier.

If the numeric value is unparsable, then Parse returns an error of type
*strconv.NumError.


### <a name="ParseBinary">func</a> [ParseBinary](/src/target/parse.go?s=1745:1786#L37)
``` go
func ParseBinary(s string) (Value, error)
```
ParseBinary is similar to Parse except all size modifiers are assumed to be
binary. In other words, both "K" and "Ki" are interpreted as 2^10.


### <a name="Value.Format">func</a> (Value) [Format](/src/target/format.go?s=907:949#L16)
``` go
func (v Value) Format(f fmt.State, c rune)
```
Format implements fmt.Formatter to provide custom, printf style formatting
for Values. It is not intended to be called directly.

