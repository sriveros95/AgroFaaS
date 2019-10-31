package function

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/spf13/viper"
)

// Handle Moon Farmer Request
func Handle(w http.ResponseWriter, r *http.Request) {
	r_receta := r.FormValue("recipe")
	rCantidad, err := strconv.ParseFloat(r.FormValue("quantity"), 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// FileDataBase Index
	fdb := viper.New()
	fdb.SetDefault("ContentDir", "data/")
	fdb.AddConfigPath("data/")
	fdb.SetDefault("Taxonomies", map[string]string{"recipe": "recipies"})
	fdb.SetConfigName("index") // name of config file (without extension)
	err = fdb.ReadInConfig()   // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		r.Header.Add("woopsydoopsy", `Fatal error recipe index`)
	}
	recetas := fdb.GetStringSlice("recipes")
	for _, receta := range recetas {
		fmt.Println(receta)
	}

	//Recipe
	receta := viper.New()
	receta.SetDefault("ContentDir", "data/")
	receta.AddConfigPath("data/")
	receta.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories", "ingredient": "ingredients"})
	receta.SetConfigName(r_receta)
	err = receta.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		r.Header.Add("woopsydoopsy", `Fatal error receta file`)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(receta.Get("name"))

	var (
		resultado CalculationResult
	)

	ingredients := receta.GetStringMap("ingredients")
	for ingredient, _ := range ingredients {
		var (
			unidad   string
			cantidad float64
		)

		if receta.IsSet(fmt.Sprintf("ingredients.%s.unit", ingredient)) {
			// If not declared we assume 1
			q := receta.GetFloat64(fmt.Sprintf("ingredients.%s.quantity", ingredient))
			if q == 0 {
				q = 1
			}
			unidad, cantidad = simplify(
				receta.GetString(fmt.Sprintf("ingredients.%s.unit", ingredient)),
				q*float64(rCantidad))
		} else if receta.IsSet(fmt.Sprintf("ingredients.%s.portions", ingredient)) {
			unidad = "portions"
			cantidad = receta.GetFloat64(fmt.Sprintf("ingredients.%s.portions", ingredient)) * float64(rCantidad)
		} else if receta.IsSet(fmt.Sprintf("ingredients.%s.quantity", ingredient)) {
			cantidad = receta.GetFloat64(fmt.Sprintf("ingredients.%s.quantity", ingredient)) * float64(rCantidad)
		}

		if receta.IsSet(fmt.Sprintf("ingredients.%s.size", ingredient)) {
			ingredient = fmt.Sprintf("%s (%s)", ingredient, receta.GetString(fmt.Sprintf("ingredients.%s.size", ingredient)))
		}

		ing := Ingredient{
			Name:     ingredient,
			Quantity: cantidad,
			Unit:     unidad,
		}
		resultado.Ingredients = append(resultado.Ingredients, ing)
	}

	//return json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultado)
}
