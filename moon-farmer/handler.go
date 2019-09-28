package function

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// Handle Moon Farmer Request
func Handle(w http.ResponseWriter, r *http.Request) {
	var input []byte

	farmingData := FarmingData

	if r.Body != nil {
		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)
		input = body
	}

	moonFuncURL := "https://sriveros95.o6s.io/moon-phase"
	earthTime := time.Now()
	var err interface{}
	if len(string(input)) != 0 {
		earthTime, err = time.Parse(time.RFC3339, string(input))
		if err != nil {
			// todo log error
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	resp, err := http.Post(moonFuncURL, "string", bytes.NewReader([]byte(earthTime.Format(time.RFC3339))))
	//resp, err := http.Get(moonFuncURL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed http.post"))
		return
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	moonData := MoonData
	err = decoder.Decode(&moonData)

	if err != nil {
		panic(err)
	}

	json.NewDecoder(resp.Body).Decode(&moonData)
	farmingData.Date = moonData.Date
	farmingData.MoonString = moonData.Text
	farmingData.MoonLight = moonData.Light
	if len(moonData.Code) == 2 {
		switch moonData.Code {
		case "NW":
			// "new (totally dark)":
			farmingData.FarmingString = "Growth season is coming!. Rest, Celebrate, Meditate."

		case "XG":
			// "waxing gibbous (increasing to full)":
			farmingData.FarmingString = "The waxing period tends to be a very productive time of period for not only planters, but all of society as a whole."
			farmingData.PlantingTime = true
			farmingData.What2Plant = "Above ground annuals, especially Leaf plants also Cereals, Herbs, Cucumbers"

		case "FQ":
			// "in its first quarter (increasing to full)":
			farmingData.FarmingString = "Growth season, - Mow lawns (to increase growth) - Graft & Prune (to increase growth) All plants producing above ground growth, fruits, or flowers benefit from a waxing moon planting."
			farmingData.PlantingTime = true
			farmingData.What2Plant = "Above ground annuals, especially Leaf plants also Cereals, Herbs, Cucumbers"

		case "XC":
			// "waxing crescent (increasing to full)":
			farmingData.FarmingString = "The waxing period tends to be a very productive time of period for not only planters, but all of society as a whole."
			farmingData.What2Plant = "Above ground annuals, especially Leaf plants also Cereals, Herbs, Cucumbers"
			farmingData.PlantingTime = true

		case "FL":
			// "full (full light)":
			farmingData.FarmingString = "Time to rest, celebrate and meditate"

		case "NG":
			// "waning gibbous (decreasing from full)", "in its last quarter (decreasing from full)":
			farmingData.FarmingString = "Harvest all crops, Fertilize, Transplant, Mow lawns & Prune"
			farmingData.PlantingTime = true
			farmingData.What2Plant = "Below ground plants, especially Root plants, Plant trees, shrubs and perennials."

		case "LQ":
			// "in its last quarter (decreasing from full)":
			farmingData.FarmingString = "Harvest all crops, Fertilize, Transplant, Mow lawns & Prune"

		case "NC":
			// "waning crescent (decreasing from full)":
			farmingData.FarmingString = "Barren phase: Time to rest. Avoid seed sowing. Harvest and store crops, Fertilize, Transplant, Destroy weed, Mow lawns & Prune"

		default:
			farmingData.FarmingString = "Unknown phase"
		}
	}

	//return json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(farmingData)
}
