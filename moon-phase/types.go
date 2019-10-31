package function

import "time"

// MoonData json response
type MoonData struct {
	Date  time.Time `json:"date"`
	Text  string    `json:"text"`
	Code  string    `json:"code"`
	Light float64   `json:"light"`
}
