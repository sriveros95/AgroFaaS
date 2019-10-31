package unit

// Speed represents a unit of speed (in meters per second, m/s)
type Speed Unit

// ...
const (
	MetersPerSecond   Speed = 1e0
	KilometersPerHour       = MetersPerSecond * 0.277778
	FeetPerSecond           = MetersPerSecond * 0.3048
	MilesPerHour            = MetersPerSecond * 0.44704
	Knot                    = MetersPerSecond * 0.514444
	SpeedOfLight            = MetersPerSecond * 299792458
)

// MetersPerSecond returns the speed in m/s
func (s Speed) MetersPerSecond() float64 {
	return float64(s)
}

// KilometersPerHour returns the speed in km/h
func (s Speed) KilometersPerHour() float64 {
	return float64(s / KilometersPerHour)
}

// FeetPerSecond returns the speed in ft/s
func (s Speed) FeetPerSecond() float64 {
	return float64(s / FeetPerSecond)
}

// MilesPerHour returns the speed in mph
func (s Speed) MilesPerHour() float64 {
	return float64(s / MilesPerHour)
}

// Knots returns the speed in knots
func (s Speed) Knots() float64 {
	return float64(s / Knot)
}

// SpeedOfLight returns the speed in c
func (s Speed) SpeedOfLight() float64 {
	return float64(s / SpeedOfLight)
}
