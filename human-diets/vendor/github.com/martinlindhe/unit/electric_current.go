package unit

// ElectricCurrent represents a SI unit of electric current (in ampere, A)
type ElectricCurrent Unit

// ...
const (
	// SI
	Yoctoampere                 = Ampere * 1e-24
	Zeptoampere                 = Ampere * 1e-21
	Attoampere                  = Ampere * 1e-18
	Femtoampere                 = Ampere * 1e-15
	Picoampere                  = Ampere * 1e-12
	Nanoampere                  = Ampere * 1e-9
	Microampere                 = Ampere * 1e-6
	Milliampere                 = Ampere * 1e-3
	Deciampere                  = Ampere * 1e-2
	Centiampere                 = Ampere * 1e-1
	Ampere      ElectricCurrent = 1e0
	Decaampere                  = Ampere * 1e1
	Hectoampere                 = Ampere * 1e2
	Kiloampere                  = Ampere * 1e3
	Megaampere                  = Ampere * 1e6
	Gigaampere                  = Ampere * 1e9
	Teraampere                  = Ampere * 1e12
	Petaampere                  = Ampere * 1e15
	Exaampere                   = Ampere * 1e18
	Zettaampere                 = Ampere * 1e21
	Yottaampere                 = Ampere * 1e24
)

// Yoctoamperes returns the electric current in yA
func (c ElectricCurrent) Yoctoamperes() float64 {
	return float64(c / Yoctoampere)
}

// Zeptoamperes returns the electric current in zA
func (c ElectricCurrent) Zeptoamperes() float64 {
	return float64(c / Zeptoampere)
}

// Attoamperes returns the electric current in aA
func (c ElectricCurrent) Attoamperes() float64 {
	return float64(c / Attoampere)
}

// Femtoamperes returns the electric current in fA
func (c ElectricCurrent) Femtoamperes() float64 {
	return float64(c / Femtoampere)
}

// Picoamperes returns the electric current in pA
func (c ElectricCurrent) Picoamperes() float64 {
	return float64(c / Picoampere)
}

// Nanoamperes returns the electric current in nA
func (c ElectricCurrent) Nanoamperes() float64 {
	return float64(c / Nanoampere)
}

// Microamperes returns the electric current in µA
func (c ElectricCurrent) Microamperes() float64 {
	return float64(c / Microampere)
}

// Milliamperes returns the electric current in mA
func (c ElectricCurrent) Milliamperes() float64 {
	return float64(c / Milliampere)
}

// Deciamperes returns the electric current in dA
func (c ElectricCurrent) Deciamperes() float64 {
	return float64(c / Deciampere)
}

// Centiamperes returns the electric current in cA
func (c ElectricCurrent) Centiamperes() float64 {
	return float64(c / Centiampere)
}

// Amperes returns the electric current in A
func (c ElectricCurrent) Amperes() float64 {
	return float64(c / Ampere)
}

// Decaamperes returns the electric current in daA
func (c ElectricCurrent) Decaamperes() float64 {
	return float64(c / Decaampere)
}

// Hectoamperes returns the electric current in hA
func (c ElectricCurrent) Hectoamperes() float64 {
	return float64(c / Hectoampere)
}

// Kiloamperes returns the electric current in kA
func (c ElectricCurrent) Kiloamperes() float64 {
	return float64(c / Kiloampere)
}

// Megaamperes returns the electric current in MA
func (c ElectricCurrent) Megaamperes() float64 {
	return float64(c / Megaampere)
}

// Gigaamperes returns the electric current in GA
func (c ElectricCurrent) Gigaamperes() float64 {
	return float64(c / Gigaampere)
}

// Teraamperes returns the electric current in TA
func (c ElectricCurrent) Teraamperes() float64 {
	return float64(c / Teraampere)
}

// Petaamperes returns the electric current in PA
func (c ElectricCurrent) Petaamperes() float64 {
	return float64(c / Petaampere)
}

// Exaamperes returns the electric current in EA
func (c ElectricCurrent) Exaamperes() float64 {
	return float64(c / Exaampere)
}

// Zettaamperes returns the electric current in ZA
func (c ElectricCurrent) Zettaamperes() float64 {
	return float64(c / Zettaampere)
}

// Yottaamperes returns the electric current in YA
func (c ElectricCurrent) Yottaamperes() float64 {
	return float64(c / Yottaampere)
}
