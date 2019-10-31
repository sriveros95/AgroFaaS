package unit

// Frequency represents a SI unit of frequency (in hertz, Hz)
type Frequency Unit

// ...
const (
	// SI
	Yoctohertz           = Hertz * 1e-24
	Zeptohertz           = Hertz * 1e-21
	Attohertz            = Hertz * 1e-18
	Femtohertz           = Hertz * 1e-15
	Picohertz            = Hertz * 1e-12
	Nanohertz            = Hertz * 1e-9
	Microhertz           = Hertz * 1e-6
	Millihertz           = Hertz * 1e-3
	Centihertz           = Hertz * 1e-2
	Decihertz            = Hertz * 1e-1
	Hertz      Frequency = 1e0
	Decahertz            = Hertz * 1e1
	Hectohertz           = Hertz * 1e2
	Kilohertz            = Hertz * 1e3
	Megahertz            = Hertz * 1e6
	Gigahertz            = Hertz * 1e9
	Terahertz            = Hertz * 1e12
	Petahertz            = Hertz * 1e15
	Exahertz             = Hertz * 1e18
	Zettahertz           = Hertz * 1e21
	Yottahertz           = Hertz * 1e24
)

// Yoctohertz returns the frequency in yHz
func (f Frequency) Yoctohertz() float64 {
	return float64(f / Yoctohertz)
}

// Zeptohertz returns the frequency in zHz
func (f Frequency) Zeptohertz() float64 {
	return float64(f / Zeptohertz)
}

// Attohertz returns the frequency in aHz
func (f Frequency) Attohertz() float64 {
	return float64(f / Attohertz)
}

// Femtohertz returns the frequency in fHz
func (f Frequency) Femtohertz() float64 {
	return float64(f / Femtohertz)
}

// Picohertz returns the frequency in pHz
func (f Frequency) Picohertz() float64 {
	return float64(f / Picohertz)
}

// Nanohertz returns the frequency in nHz
func (f Frequency) Nanohertz() float64 {
	return float64(f / Nanohertz)
}

// Microhertz returns the frequency in µHz
func (f Frequency) Microhertz() float64 {
	return float64(f / Microhertz)
}

// Millihertz returns the frequency in mHz
func (f Frequency) Millihertz() float64 {
	return float64(f / Millihertz)
}

// Centihertz returns the frequency in cHz
func (f Frequency) Centihertz() float64 {
	return float64(f / Centihertz)
}

// Decihertz returns the frequency in dHz
func (f Frequency) Decihertz() float64 {
	return float64(f / Decihertz)
}

// Hertz returns the frequency in Hz
func (f Frequency) Hertz() float64 {
	return float64(f)
}

// Decahertz returns the frequency in daHz
func (f Frequency) Decahertz() float64 {
	return float64(f / Decahertz)
}

// Hectohertz returns the frequency in hHz
func (f Frequency) Hectohertz() float64 {
	return float64(f / Hectohertz)
}

// Kilohertz returns the frequency in kHz
func (f Frequency) Kilohertz() float64 {
	return float64(f / Kilohertz)
}

// Megahertz returns the frequency in MHz
func (f Frequency) Megahertz() float64 {
	return float64(f / Megahertz)
}

// Gigahertz returns the frequency in GHz
func (f Frequency) Gigahertz() float64 {
	return float64(f / Gigahertz)
}

// Terahertz returns the frequency in THz
func (f Frequency) Terahertz() float64 {
	return float64(f / Terahertz)
}

// Petahertz returns the frequency in PHz
func (f Frequency) Petahertz() float64 {
	return float64(f / Petahertz)
}

// Exahertz returns the frequency in EHz
func (f Frequency) Exahertz() float64 {
	return float64(f / Exahertz)
}

// Zettahertz returns the frequency in ZHz
func (f Frequency) Zettahertz() float64 {
	return float64(f / Zettahertz)
}

// Yottahertz returns the frequency in YHz
func (f Frequency) Yottahertz() float64 {
	return float64(f / Yottahertz)
}
