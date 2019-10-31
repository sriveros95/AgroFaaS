package unit

// Energy represents a SI unit of energy (in joules, J)
type Energy Unit

// ...
const (
	// SI
	Yoctojoule        = Joule * 1e-24
	Zeptojoule        = Joule * 1e-21
	Attojoule         = Joule * 1e-18
	Femtojoule        = Joule * 1e-15
	Picojoule         = Joule * 1e-12
	Nanojoule         = Joule * 1e-9
	Microjoule        = Joule * 1e-6
	Millijoule        = Joule * 1e-3
	Centijoule        = Joule * 1e-2
	Decijoule         = Joule * 1e-1
	Joule      Energy = 1e0
	Decajoule         = Joule * 1e1
	Hectojoule        = Joule * 1e2
	Kilojoule         = Joule * 1e3
	Megajoule         = Joule * 1e6
	Gigajoule         = Joule * 1e9
	Terajoule         = Joule * 1e12
	Petajoule         = Joule * 1e15
	Exajoule          = Joule * 1e18
	Zettajoule        = Joule * 1e21
	Yottajoule        = Joule * 1e24

	// SI-derived
	YoctowattHour = WattHour * 1e-24
	ZeptowattHour = WattHour * 1e-21
	AttowattHour  = WattHour * 1e-18
	FemtowattHour = WattHour * 1e-15
	PicowattHour  = WattHour * 1e-12
	NanowattHour  = WattHour * 1e-9
	MicrowattHour = WattHour * 1e-6
	MilliwattHour = WattHour * 1e-3
	CentiwattHour = WattHour * 1e-2
	DeciwattHour  = WattHour * 1e-1
	WattHour      = Joule * 3600
	DecawattHour  = WattHour * 1e1
	HectowattHour = WattHour * 1e2
	KilowattHour  = WattHour * 1e3
	MegawattHour  = WattHour * 1e6
	GigawattHour  = WattHour * 1e9
	TerawattHour  = WattHour * 1e12
	PetawattHour  = WattHour * 1e15
	ExawattHour   = WattHour * 1e18
	ZettawattHour = WattHour * 1e21
	YottawattHour = WattHour * 1e24

	// constant from https://en.wikipedia.org/wiki/Calorie#Definitions
	Gramcalorie = Joule * 4.184
	Kilocalorie = Gramcalorie * 1e3
	Megacalorie = Gramcalorie * 1e6
)

// Yoctojoules returns the energy in yJ
func (e Energy) Yoctojoules() float64 {
	return float64(e / Yoctojoule)
}

// Zeptojoules returns the energy in zJ
func (e Energy) Zeptojoules() float64 {
	return float64(e / Zeptojoule)
}

// Attojoules returns the energy in aJ
func (e Energy) Attojoules() float64 {
	return float64(e / Attojoule)
}

// Femtojoules returns the energy in fJ
func (e Energy) Femtojoules() float64 {
	return float64(e / Femtojoule)
}

// Picojoules returns the energy in pJ
func (e Energy) Picojoules() float64 {
	return float64(e / Picojoule)
}

// Nanojoules returns the energy in nJ
func (e Energy) Nanojoules() float64 {
	return float64(e / Nanojoule)
}

// Microjoules returns the energy in µJ
func (e Energy) Microjoules() float64 {
	return float64(e / Microjoule)
}

// Millijoules returns the energy in mJ
func (e Energy) Millijoules() float64 {
	return float64(e / Millijoule)
}

// Centijoules returns the energy in cJ
func (e Energy) Centijoules() float64 {
	return float64(e / Centijoule)
}

// Decijoules returns the energy in dJ
func (e Energy) Decijoules() float64 {
	return float64(e / Decijoule)
}

// Joules returns the energy in J
func (e Energy) Joules() float64 {
	return float64(e / Joule)
}

// Decajoules returns the energy in dJ
func (e Energy) Decajoules() float64 {
	return float64(e / Decajoule)
}

// Hectojoules returns the energy in hJ
func (e Energy) Hectojoules() float64 {
	return float64(e / Hectojoule)
}

// Kilojoules returns the energy in kJ
func (e Energy) Kilojoules() float64 {
	return float64(e / Kilojoule)
}

// Megajoules returns the energy in MJ
func (e Energy) Megajoules() float64 {
	return float64(e / Megajoule)
}

// Gigajoules returns the energy in GJ
func (e Energy) Gigajoules() float64 {
	return float64(e / Gigajoule)
}

// Terajoules returns the energy in TJ
func (e Energy) Terajoules() float64 {
	return float64(e / Terajoule)
}

// Petajoules returns the energy in PJ
func (e Energy) Petajoules() float64 {
	return float64(e / Petajoule)
}

// Exajoules returns the energy in EJ
func (e Energy) Exajoules() float64 {
	return float64(e / Exajoule)
}

// Zettajoules returns the energy in ZJ
func (e Energy) Zettajoules() float64 {
	return float64(e / Zettajoule)
}

// Yottajoules returns the energy in YJ
func (e Energy) Yottajoules() float64 {
	return float64(e / Yottajoule)
}

// AttowattHours returns the energy in aWh
func (e Energy) AttowattHours() float64 {
	return float64(e / AttowattHour)
}

// YoctowattHours returns the energy in yWh
func (e Energy) YoctowattHours() float64 {
	return float64(e / YoctowattHour)
}

// ZeptowattHours returns the energy in zWh
func (e Energy) ZeptowattHours() float64 {
	return float64(e / ZeptowattHour)
}

// FemtowattHours returns the energy in fWh
func (e Energy) FemtowattHours() float64 {
	return float64(e / FemtowattHour)
}

// PicowattHours returns the energy in pWh
func (e Energy) PicowattHours() float64 {
	return float64(e / PicowattHour)
}

// NanowattHours returns the energy in nWh
func (e Energy) NanowattHours() float64 {
	return float64(e / NanowattHour)
}

// MicrowattHours returns the energy in µWh
func (e Energy) MicrowattHours() float64 {
	return float64(e / MicrowattHour)
}

// MilliwattHours returns the energy in mWh
func (e Energy) MilliwattHours() float64 {
	return float64(e / MilliwattHour)
}

// CentiwattHours returns the energy in cWh
func (e Energy) CentiwattHours() float64 {
	return float64(e / CentiwattHour)
}

// DeciwattHours returns the energy in dWh
func (e Energy) DeciwattHours() float64 {
	return float64(e / DeciwattHour)
}

// WattHours returns the energy in Wh
func (e Energy) WattHours() float64 {
	return float64(e / WattHour)
}

// DecawattHours returns the energy in daWh
func (e Energy) DecawattHours() float64 {
	return float64(e / DecawattHour)
}

// HectowattHours returns the energy in hWh
func (e Energy) HectowattHours() float64 {
	return float64(e / HectowattHour)
}

// KilowattHours returns the energy in kWh
func (e Energy) KilowattHours() float64 {
	return float64(e / KilowattHour)
}

// MegawattHours returns the energy in MWh
func (e Energy) MegawattHours() float64 {
	return float64(e / MegawattHour)
}

// GigawattHours returns the energy in GWh
func (e Energy) GigawattHours() float64 {
	return float64(e / GigawattHour)
}

// TerawattHours returns the energy in TWh
func (e Energy) TerawattHours() float64 {
	return float64(e / TerawattHour)
}

// PetawattHours returns the energy in PWh
func (e Energy) PetawattHours() float64 {
	return float64(e / PetawattHour)
}

// ExawattHours returns the energy in EWh
func (e Energy) ExawattHours() float64 {
	return float64(e / ExawattHour)
}

// ZettawattHours returns the energy in ZWh
func (e Energy) ZettawattHours() float64 {
	return float64(e / ZettawattHour)
}

// YottawattHours returns the energy in YWh
func (e Energy) YottawattHours() float64 {
	return float64(e / YottawattHour)
}
