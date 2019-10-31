package unit

// AmountOfSubstance represents a SI unit of amount of substance (in mole, mol)
type AmountOfSubstance Unit

// ...
const (
	Mole AmountOfSubstance = 1e0 // SI
)

// Moles returns the amount of substance in mol
func (a AmountOfSubstance) Moles() float64 {
	return float64(a)
}
