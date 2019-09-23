// Adaptation of vegaseat's "Moon Phase at a given date (Python)" (02may2013)
// https://www.daniweb.com/programming/software-development/code/453788/moon-phase-at-a-given-date-python

package function

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"
	"time"
)

// Handle a serverless request
func Handle(w http.ResponseWriter, r *http.Request) {
	var input []byte

	moonData := MoonData

	if r.Body != nil {
		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)
		input = body
	}

	moonData.Date = time.Now()
	var err interface{}

	if len(string(input)) != 0 {
		moonData.Date, err = time.Parse(time.RFC3339, string(input))
		if err != nil {
			// todo log error
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	var (
		ages         = [...]int32{18, 0, 11, 22, 3, 14, 25, 6, 17, 28, 9, 20, 1, 12, 23, 4, 15, 26, 7}
		offsets      = [...]int32{-1, 1, 0, 1, 2, 3, 4, 5, 7, 7, 9, 9}
		descriptions = [...]string{"new (totally dark)",
			"waxing crescent (increasing to full)",
			"in its first quarter (increasing to full)",
			"waxing gibbous (increasing to full)",
			"full (full light)",
			"waning gibbous (decreasing from full)",
			"in its last quarter (decreasing from full)",
			"waning crescent (decreasing from full)"}
		codes = [...]string{"NW", "XC", "FQ", "XG", "FL", "NG", "LQ", "NC"}
		day   = int32(moonData.Date.Day())
		year  = int32(moonData.Date.Year())
		month = int32(moonData.Date.Month())
	)

	if day == 31 {
		day = 1
	}

	daysIntoPhase := ((ages[(year+1)%19] +
		((day + offsets[month-1]) % 30) +
		(year - 1900)) % 30)

	index := int((daysIntoPhase + 2) * 16 / 59.0)
	//print(index)  // test
	if index > 7 {
		index = 7
	}
	moonData.Text = descriptions[index]
	moonData.Code = codes[index]
	// light should be 100% 15 days into phase
	moonData.Light = float64(2 * daysIntoPhase * 100 / 29)
	if moonData.Light > 100 {
		moonData.Light = math.Abs(moonData.Light - 200)
	}

	//return json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(moonData)
}
