package unit

// Illuminance represents a SI unit for illuminance (in lux, lx)
type Illuminance Unit

// constants
const (
	Lux Illuminance = 1e0 // SI
)

// Lux returns the illuminance in lx
func (i Illuminance) Lux() float64 {
	return float64(i)
}
