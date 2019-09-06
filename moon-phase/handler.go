// Adaptation of vegaseat's "Moon Phase at a given date (Python)" (02may2013)
// https://www.daniweb.com/programming/software-development/code/453788/moon-phase-at-a-given-date-python

package function

import (
	"fmt"
	"math"
	"time"
)

// Handle a serverless request
func Handle(req []byte) string {

	earthTime := time.Now()
	var err interface{}
	if len(string(req)) != 0 {
		earthTime, err = time.Parse(time.RFC3339, string(req))
		if err != nil {
			// todo log error
			return "Failed"
		}
	}
	var ages = [...]int32{18, 0, 11, 22, 3, 14, 25, 6, 17, 28, 9, 20, 1, 12, 23, 4, 15, 26, 7}
	var offsets = [...]int32{-1, 1, 0, 1, 2, 3, 4, 5, 7, 7, 9, 9}
	var descriptions = [...]string{"new (totally dark)",
		"waxing crescent (increasing to full)",
		"in its first quarter (increasing to full)",
		"waxing gibbous (increasing to full)",
		"full (full light)",
		"waning gibbous (decreasing from full)",
		"in its last quarter (decreasing from full)",
		"waning crescent (decreasing from full)"}
	var months = [...]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

	var (
		day   int32
		month int32
		year  int32
	)

	day = int32(earthTime.Day())
	year = int32(earthTime.Year())
	month = int32(earthTime.Month())

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
	status := descriptions[index]
	// light should be 100% 15 days into phase
	var light float64
	light = float64(2 * daysIntoPhase * 100 / 29)
	if light > 100 {
		light = math.Abs(light - 200)
	}

	dateFstr := fmt.Sprintf("%d %s %d", day, months[month-1], year)

	response := fmt.Sprintf("%v, %v, %v", dateFstr, status, light)

	//fmt.Println(response)
	return (response)
}
