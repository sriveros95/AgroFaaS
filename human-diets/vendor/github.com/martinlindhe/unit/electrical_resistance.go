package unit

// ElectricalResistance represents a SI derived unit of electrical resistance (in ohm, Ω)
type ElectricalResistance Unit

// ...
const (
	Ohm ElectricalResistance = 1e0 // SI
)

// Ohms returns the electrical resistance in Ω
func (e ElectricalResistance) Ohms() float64 {
	return float64(e)
}
