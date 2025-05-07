package models

type WeatherResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Hourly    struct {
		Time          []string  `json:"time"`
		Temperature2m []float64 `json:"temperature_2m"`
	} `json:"hourly"`
}

type GeoResponse struct {
	Results []struct{
		Latitude  float64 `json:"latitude"`
        Longitude float64 `json:"longitude"`
        Name      string  `json:"name"`
    } `json:"results"`	
	}
}