package function

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

// Handle Moon Farmer Request
func Handle(w http.ResponseWriter, r *http.Request) {
	var (
		reqIn CalculationRequest
	)

	err := json.NewDecoder(r.Body).Decode(reqIn)
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
		panic(fmt.Errorf("Fatal error recipe index file: %s \n", err))
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
	receta.SetConfigName(reqIn.Recipe)
	err = receta.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error receta file: %s \n", err))
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
				q*float64(reqIn.Quantity))
		} else if receta.IsSet(fmt.Sprintf("ingredients.%s.portions", ingredient)) {
			unidad = "portions"
			cantidad = receta.GetFloat64(fmt.Sprintf("ingredients.%s.portions", ingredient)) * float64(reqIn.Quantity)
		} else if receta.IsSet(fmt.Sprintf("ingredients.%s.quantity", ingredient)) {
			cantidad = receta.GetFloat64(fmt.Sprintf("ingredients.%s.quantity", ingredient)) * float64(reqIn.Quantity)
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
