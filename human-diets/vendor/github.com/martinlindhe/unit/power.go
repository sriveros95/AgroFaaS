package unit

// Power represents a SI unit of power (in watts, W)
type Power Unit

// ...
const (
	// SI
	Yoctowatt       = Watt * 1e-24
	Zeptowatt       = Watt * 1e-21
	Attowatt        = Watt * 1e-18
	Femtowatt       = Watt * 1e-15
	Picowatt        = Watt * 1e-12
	Nanowatt        = Watt * 1e-9
	Microwatt       = Watt * 1e-6
	Milliwatt       = Watt * 1e-3
	Centiwatt       = Watt * 1e-2
	Deciwatt        = Watt * 1e-1
	Watt      Power = 1e0
	Decawatt        = Watt * 1e1
	Hectowatt       = Watt * 1e2
	Kilowatt        = Watt * 1e3
	Megawatt        = Watt * 1e6
	Gigawatt        = Watt * 1e9
	Terawatt        = Watt * 1e12
	Petawatt        = Watt * 1e15
	Exawatt         = Watt * 1e18
	Zettawatt       = Watt * 1e21
	Yottawatt       = Watt * 1e24

	// non-SI
	Pferdestarke = Watt * 735.49875
)

// Yoctowatts returns the power in yW
func (p Power) Yoctowatts() float64 {
	return float64(p / Yoctowatt)
}

// Zeptowatts returns the power in zW
func (p Power) Zeptowatts() float64 {
	return float64(p / Zeptowatt)
}

// Attowatts returns the power in aW
func (p Power) Attowatts() float64 {
	return float64(p / Attowatt)
}

// Femtowatts returns the power in fW
func (p Power) Femtowatts() float64 {
	return float64(p / Femtowatt)
}

// Picowatts returns the power in pW
func (p Power) Picowatts() float64 {
	return float64(p / Picowatt)
}

// Nanowatts returns the power in nW
func (p Power) Nanowatts() float64 {
	return float64(p / Nanowatt)
}

// Microwatts returns the power in µW
func (p Power) Microwatts() float64 {
	return float64(p / Microwatt)
}

// Milliwatts returns the power in mW
func (p Power) Milliwatts() float64 {
	return float64(p / Milliwatt)
}

// Centiwatts returns the power in cW
func (p Power) Centiwatts() float64 {
	return float64(p / Centiwatt)
}

// Deciwatts returns the power in dW
func (p Power) Deciwatts() float64 {
	return float64(p / Deciwatt)
}

// Watts returns the power in W
func (p Power) Watts() float64 {
	return float64(p)
}

// Decawatts returns the power in daW
func (p Power) Decawatts() float64 {
	return float64(p / Decawatt)
}

// Hectowatts returns the power in hW
func (p Power) Hectowatts() float64 {
	return float64(p / Hectowatt)
}

// Kilowatts returns the power in kW
func (p Power) Kilowatts() float64 {
	return float64(p / Kilowatt)
}

// Megawatts returns the power in MW
func (p Power) Megawatts() float64 {
	return float64(p / Megawatt)
}

// Gigawatts returns the power in GW
func (p Power) Gigawatts() float64 {
	return float64(p / Gigawatt)
}

// Petawatts returns the power in PW
func (p Power) Petawatts() float64 {
	return float64(p / Petawatt)
}

// Exawatts returns the power in EW
func (p Power) Exawatts() float64 {
	return float64(p / Exawatt)
}

// Zettawatts returns the power in ZW
func (p Power) Zettawatts() float64 {
	return float64(p / Zettawatt)
}

// Yottawatts returns the power in YW
func (p Power) Yottawatts() float64 {
	return float64(p / Yottawatt)
}

// Terawatts returns the power in tW
func (p Power) Terawatts() float64 {
	return float64(p / Terawatt)
}

// Pferdestarke returns the power in PS
func (p Power) Pferdestarke() float64 {
	return float64(p / Pferdestarke)
}
