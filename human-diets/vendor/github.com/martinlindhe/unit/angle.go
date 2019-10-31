package unit

import "math"

// Angle represents a SI unit of angle (in radians, ㎭)
type Angle Unit

// ...
const (
	Yoctoradian          = Radian * 1e-24
	Zeptoradian          = Radian * 1e-21
	Attoradian           = Radian * 1e-18
	Femtoradian          = Radian * 1e-15
	Picoradian           = Radian * 1e-12
	Nanoradian           = Radian * 1e-9
	Microradian          = Radian * 1e-6
	Milliradian          = Radian * 1e-3
	Centiradian          = Radian * 1e-2
	Deciradian           = Radian * 1e-1
	Radian         Angle = 1e0
	Degree               = Radian * math.Pi / 180
	Arcminute            = Degree / 60
	Arcsecond            = Degree / 3600
	Milliarcsecond       = Arcsecond * 1e-3
	Microarcsecond       = Arcsecond * 1e-6
)

// Yoctoradians returns the angle in y㎭
func (a Angle) Yoctoradians() float64 {
	return float64(a / Yoctoradian)
}

// Zeptoradians returns the angle in z㎭
func (a Angle) Zeptoradians() float64 {
	return float64(a / Zeptoradian)
}

// Attoradians returns the angle in a㎭
func (a Angle) Attoradians() float64 {
	return float64(a / Attoradian)
}

// Femtoradians returns the angle in f㎭
func (a Angle) Femtoradians() float64 {
	return float64(a / Femtoradian)
}

// Picoradians returns the angle in p㎭
func (a Angle) Picoradians() float64 {
	return float64(a / Picoradian)
}

// Nanoradians returns the angle in n㎭
func (a Angle) Nanoradians() float64 {
	return float64(a / Nanoradian)
}

// Microradians returns the angle in µ㎭
func (a Angle) Microradians() float64 {
	return float64(a / Microradian)
}

// Milliradians returns the angle in m㎭
func (a Angle) Milliradians() float64 {
	return float64(a / Milliradian)
}

// Deciradians returns the angle in d㎭
func (a Angle) Deciradians() float64 {
	return float64(a / Deciradian)
}

// Centiradians returns the angle in c㎭
func (a Angle) Centiradians() float64 {
	return float64(a / Centiradian)
}

// Radians returns the angle in ㎭
func (a Angle) Radians() float64 {
	return float64(a / Radian)
}

// Degrees returns the angle in °
func (a Angle) Degrees() float64 {
	return float64(a / Degree)
}

// Arcminutes returns the angle in amin
func (a Angle) Arcminutes() float64 {
	return float64(a / Arcminute)
}

// Arcseconds returns the angle in asec
func (a Angle) Arcseconds() float64 {
	return float64(a / Arcsecond)
}

// Milliarcseconds returns the angle in masec
func (a Angle) Milliarcseconds() float64 {
	return float64(a / Milliarcsecond)
}

// Microarcseconds returns the angle in µasec
func (a Angle) Microarcseconds() float64 {
	return float64(a / Microarcsecond)
}
