package function

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	var input []byte

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
			return
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Failed http.Post"))
			return
		}
	}
	resp, err := http.Post(moonFuncURL, "string", bytes.NewReader([]byte(earthTime.Format(time.RFC3339))))
	//resp, err := http.Get(moonFuncURL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed http.Get"))
		return
	}
	defer resp.Body.Close()
	moonString, err := ioutil.ReadAll(resp.Body)
	fmt.Println(moonString)
	moonData := strings.Split(string(moonString), ", ")
	gardenString := "failed to process moon data"
	if len(moonData) == 3 {
		switch moonData[1] {
		case " new (totally dark)", " in its first quarter (increasing to full)":
			gardenString = "'growth' season, considered fertile and wet. All plants producing above ground growth, fruits, or flowers benefit from a waxing moon planting."

		case " waxing gibbous (increasing to full)":
			gardenString = "The waxing period tends to be a very productive time of period for not only planters, but all of society as a whole. Sow/Plant: Above ground annuals, especially Fruit plants also Cereals and Flowers"

		case " waxing crescent (increasing to full)":
			gardenString = "Sow/Plant: Above ground annuals, especially Leaf plants also Cereals, Herbs, Cucumbers"

		case " full (full light)":
			gardenString = "Time to rest, celebrate and meditate"

		case " waning gibbous (decreasing from full)", " in its last quarter (decreasing from full)":
			gardenString = "Sow/Plant: Below ground plants, especially Root plants, Plant trees, shrubs and perennials. Harvest all crops, Fertilize, Transplant, Mow lawns & Prune"

		case " waning crescent (decreasing from full)":
			gardenString = "Barren phase: Time to rest. Avoid seed sowing, Harvest and store crops, Fertilize, Transplant, Destroy weed, Mow lawns & Prune"

		default:
			gardenString = "Unknown phase"
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string(moonString)))
	w.Write([]byte(string(gardenString)))
}
