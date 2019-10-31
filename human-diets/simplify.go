package function

import (
	"github.com/martinlindhe/unit"
)

type MyUnit int

// If unit too small and number too big it converts it to a bigger unit
func simplify(unidad string, cantidad float64) (string, float64) {
	if unidad == "grams" && cantidad > 1000000 {
		unidad = "tonnes"
		n := MyUnit(cantidad)
		grams := unit.Mass(n) * unit.Gram
		cantidad = grams.Tonnes()
	} else if unidad == "grams" && cantidad > 1000 {
		unidad = "kilograms"
		n := MyUnit(cantidad)
		grams := unit.Mass(n) * unit.Gram
		cantidad = grams.Kilograms()
	}

	if unidad == "milliliters" && cantidad > 1000000 {
		unidad = "gallon (imp)"
		n := MyUnit(cantidad)
		mls := unit.Volume(n) * unit.Milliliter
		cantidad = mls.ImperialGallons()
	} else if unidad == "milliliters" && cantidad > 1000 {
		unidad = "liters"
		n := MyUnit(cantidad)
		mls := unit.Volume(n) * unit.Milliliter
		cantidad = mls.Liters()
	}
	return unidad, cantidad
}
