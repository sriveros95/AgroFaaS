package function

// CalculationResult structures a json response for a basic calc. experiment
type CalculationResult struct {
	Text        string       `json:"text"`
	Ingredients []Ingredient `json:"ingredients"`
}

// Ingredient is an individual Ingredient result
type Ingredient struct {
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	Unit     string  `json:"unit"`
}
