package function

import "time"

// MoonData json response
type MoonData struct {
	Date  time.Time `json:"date"`
	Text  string    `json:"text"`
	Code  string    `json:"code"`
	Light float64   `json:"light"`
}

// FarmingData json response
type FarmingData struct {
	Date          time.Time `json:"date"`
	MoonString    string    `json:"moonString"`
	MoonLight     float64   `json:"moonLight"`
	MoonCode      string    `json:"moonCode"`
	FarmingString string    `json:"farmingString"`
	PlantingTime  bool      `json:"plantingTime"`
	What2Plant    string    `json:"what2Plant"`
}
