package unit

// Voltage represents a unit of voltage (in volt, V)
type Voltage Unit

// ...
const (
	// SI
	Yoctovolt         = Volt * 1e-24
	Zeptovolt         = Volt * 1e-21
	Attovolt          = Volt * 1e-18
	Femtovolt         = Volt * 1e-15
	Picovolt          = Volt * 1e-12
	Nanovolt          = Volt * 1e-9
	Microvolt         = Volt * 1e-6
	Millivolt         = Volt * 1e-3
	Centivolt         = Volt * 1e-2
	Decivolt          = Volt * 1e-1
	Volt      Voltage = 1e0
	Decavolt          = Volt * 1e1
	Hectovolt         = Volt * 1e2
	Kilovolt          = Volt * 1e3
	Megavolt          = Volt * 1e6
	Gigavolt          = Volt * 1e9
	Teravolt          = Volt * 1e12
	Petavolt          = Volt * 1e15
	Exavolt           = Volt * 1e18
	Zettavolt         = Volt * 1e21
	Yottavolt         = Volt * 1e24
)

// Yoctovolts returns the voltage in yV
func (v Voltage) Yoctovolts() float64 {
	return float64(v / Yoctovolt)
}

// Zeptovolts returns the voltage in zV
func (v Voltage) Zeptovolts() float64 {
	return float64(v / Zeptovolt)
}

// Attovolts returns the voltage in aV
func (v Voltage) Attovolts() float64 {
	return float64(v / Attovolt)
}

// Femtovolts returns the voltage in fV
func (v Voltage) Femtovolts() float64 {
	return float64(v / Femtovolt)
}

// Picovolts returns the voltage in pV
func (v Voltage) Picovolts() float64 {
	return float64(v / Picovolt)
}

// Nanovolts returns the voltage in nV
func (v Voltage) Nanovolts() float64 {
	return float64(v / Nanovolt)
}

// Microvolts returns the voltage in µV
func (v Voltage) Microvolts() float64 {
	return float64(v / Microvolt)
}

// Millivolts returns the voltage in mV
func (v Voltage) Millivolts() float64 {
	return float64(v / Millivolt)
}

// Centivolts returns the voltage in cV
func (v Voltage) Centivolts() float64 {
	return float64(v / Centivolt)
}

// Decivolts returns the voltage in dV
func (v Voltage) Decivolts() float64 {
	return float64(v / Decivolt)
}

// Volts returns the voltage in V
func (v Voltage) Volts() float64 {
	return float64(v)
}

// Decavolts returns the voltage in daV
func (v Voltage) Decavolts() float64 {
	return float64(v / Decavolt)
}

// Hectovolts returns the voltage in hV
func (v Voltage) Hectovolts() float64 {
	return float64(v / Hectovolt)
}

// Kilovolts returns the voltage in kV
func (v Voltage) Kilovolts() float64 {
	return float64(v / Kilovolt)
}

// Megavolts returns the voltage in MV
func (v Voltage) Megavolts() float64 {
	return float64(v / Megavolt)
}

// Gigavolts returns the voltage in GV
func (v Voltage) Gigavolts() float64 {
	return float64(v / Gigavolt)
}

// Teravolts returns the voltage in TV
func (v Voltage) Teravolts() float64 {
	return float64(v / Teravolt)
}

// Petavolts returns the voltage in PV
func (v Voltage) Petavolts() float64 {
	return float64(v / Petavolt)
}

// Exavolts returns the voltage in EV
func (v Voltage) Exavolts() float64 {
	return float64(v / Exavolt)
}

// Zettavolts returns the voltage in ZV
func (v Voltage) Zettavolts() float64 {
	return float64(v / Zettavolt)
}

// Yottavolts returns the voltage in YV
func (v Voltage) Yottavolts() float64 {
	return float64(v / Yottavolt)
}
