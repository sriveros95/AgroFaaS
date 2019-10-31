package unit

// LuminousIntensity represents a SI unit for luminous intensity (in candela, cd)
type LuminousIntensity Unit

// constants
const (
	Candela LuminousIntensity = 1e0
)

// Candela returns the luminous intensity in cd
func (l LuminousIntensity) Candela() float64 {
	return float64(l)
}
