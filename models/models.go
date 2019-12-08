package models

// Main captures the main data in the OpenWeatherMap data
type Main struct {
	Main TemperatureData `json:"main"`
}

// TemperatureData captures the temperature data from the OpenWeatherMap response
type TemperatureData struct {
	Temperature float64 `json:"temp"`
	Pressure    float64 `json:"pressure"`
	Humidity    float64 `json:"humidity"`
	High        float64 `json:"temp_max"`
	Low         float64 `json:"temp_min"`
}
