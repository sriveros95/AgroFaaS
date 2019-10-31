package unit

// Duration represents a SI unit of time (in seconds, s)
type Duration Unit

// ...
const (
	// SI
	Yoctosecond          = Second * 1e-24
	Zeptosecond          = Second * 1e-21
	Attosecond           = Second * 1e-18
	Femtosecond          = Second * 1e-15
	Picosecond           = Second * 1e-12
	Nanosecond           = Second * 1e-9
	Microsecond          = Second * 1e-6
	Millisecond          = Second * 1e-3
	Centisecond          = Second * 1e-2
	Decisecond           = Second * 1e-1
	Second      Duration = 1e0
	Decasecond           = Second * 1e1
	Hectosecond          = Second * 1e2
	Kilosecond           = Second * 1e3
	Megasecond           = Second * 1e6
	Gigasecond           = Second * 1e9
	Terasecond           = Second * 1e12
	Petasecond           = Second * 1e15
	Exasecond            = Second * 1e18
	Zettasecond          = Second * 1e21
	Yottasecond          = Second * 1e24

	// non-SI
	Minute         = Second * 60
	Hour           = Minute * 60
	Day            = Hour * 24
	Week           = Day * 7
	ThirtyDayMonth = Day * 30
	JulianYear     = Day * 365.25
)

// Yoctoseconds returns the time in ys
func (t Duration) Yoctoseconds() float64 {
	return float64(t / Yoctosecond)
}

// Zeptoseconds returns the time in zs
func (t Duration) Zeptoseconds() float64 {
	return float64(t / Zeptosecond)
}

// Attoseconds returns the time in as
func (t Duration) Attoseconds() float64 {
	return float64(t / Attosecond)
}

// Femtoseconds returns the time in fs
func (t Duration) Femtoseconds() float64 {
	return float64(t / Femtosecond)
}

// Picoseconds returns the time in ps
func (t Duration) Picoseconds() float64 {
	return float64(t / Picosecond)
}

// Nanoseconds returns the time in ns
func (t Duration) Nanoseconds() float64 {
	return float64(t / Nanosecond)
}

// Microseconds returns the time in µs
func (t Duration) Microseconds() float64 {
	return float64(t / Microsecond)
}

// Milliseconds returns the time in ms
func (t Duration) Milliseconds() float64 {
	return float64(t / Millisecond)
}

// Centiseconds returns the time in cs
func (t Duration) Centiseconds() float64 {
	return float64(t / Centisecond)
}

// Deciseconds returns the time in ds
func (t Duration) Deciseconds() float64 {
	return float64(t / Decisecond)
}

// Seconds returns the time in s
func (t Duration) Seconds() float64 {
	return float64(t / Second)
}

// Decaseconds returns the time in das
func (t Duration) Decaseconds() float64 {
	return float64(t / Decasecond)
}

// Hectoseconds returns the time in hs
func (t Duration) Hectoseconds() float64 {
	return float64(t / Hectosecond)
}

// Kiloseconds returns the time in ks
func (t Duration) Kiloseconds() float64 {
	return float64(t / Kilosecond)
}

// Megaseconds returns the time in Ms
func (t Duration) Megaseconds() float64 {
	return float64(t / Megasecond)
}

// Gigaseconds returns the time in Gs
func (t Duration) Gigaseconds() float64 {
	return float64(t / Gigasecond)
}

// Teraseconds returns the time in Ts
func (t Duration) Teraseconds() float64 {
	return float64(t / Terasecond)
}

// Petaseconds returns the time in Ps
func (t Duration) Petaseconds() float64 {
	return float64(t / Petasecond)
}

// Exaseconds returns the time in Es
func (t Duration) Exaseconds() float64 {
	return float64(t / Exasecond)
}

// Zettaseconds returns the time in Zs
func (t Duration) Zettaseconds() float64 {
	return float64(t / Zettasecond)
}

// Yottaseconds returns the time in Ys
func (t Duration) Yottaseconds() float64 {
	return float64(t / Yottasecond)
}

// Minutes returns the time in m
func (t Duration) Minutes() float64 {
	return float64(t / Minute)
}

// Hours returns the time in h
func (t Duration) Hours() float64 {
	return float64(t / Hour)
}

// Days returns the time in d
func (t Duration) Days() float64 {
	return float64(t / Day)
}

// Weeks returns the time in w
func (t Duration) Weeks() float64 {
	return float64(t / Week)
}

// ThirtyDayMonths returns the time in M
func (t Duration) ThirtyDayMonths() float64 {
	return float64(t / ThirtyDayMonth)
}

// JulianYears returns the time in Y
func (t Duration) JulianYears() float64 {
	return float64(t / JulianYear)
}
