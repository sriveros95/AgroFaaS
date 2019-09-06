package function

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Handle a serverless request
func Handle(req []byte) string {
	moonFuncURL := "http://127.0.0.1:8080/function/sriveros95/moon-phase"
	earthTime := time.Now()
	var err interface{}
	if len(string(req)) != 0 {
		earthTime, err = time.Parse(time.RFC3339, string(req))
		if err != nil {
			// todo log error
			return "Failed lunar-agriculture time.Parse"
		}
	}
	resp, err := http.Post(moonFuncURL, "string", bytes.NewReader([]byte(earthTime.Format(time.RFC3339))))
	if err != nil {
		fmt.Println(err)
		return "Failed http.Post"
	}
	defer resp.Body.Close()
	moonString, err := ioutil.ReadAll(resp.Body)
	fmt.Println(moonString)
	moonData := strings.Split(string(moonString), ",")
	switch moonData[1] {
	case " new (totally dark)", " in its first quarter (increasing to full)":
		return "'growth' season, considered fertile and wet. All plants producing above ground growth, fruits, or flowers benefit from a waxing moon planting."

	case " waxing gibbous (increasing to full)":
		return "The waxing period tends to be a very productive time of period for not only planters, but all of society as a whole. Sow/Plant: Above ground annuals, especially Fruit plants also Cereals and Flowers"

	case " waxing crescent (increasing to full)":
		return "Sow/Plant: Above ground annuals, especially Leaf plants also Cereals, Herbs, Cucumbers"

	case " full (full light)":
		return "Time to rest, celebrate and meditate"

	case " waning gibbous (decreasing from full)", " in its last quarter (decreasing from full)":
		return "Sow/Plant: Below ground plants, especially Root plants, Plant trees, shrubs and perennials. Harvest all crops, Fertilize, Transplant, Mow lawns & Prune"

	case " waning crescent (decreasing from full)":
		return "Barren phase: Time to rest. Avoid seed sowing, Harvest and store crops, Fertilize, Transplant, Destroy weed, Mow lawns & Prune"

	default:
		return "Unknown phase"
	}
}
