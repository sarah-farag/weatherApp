package api

import (
	"net/http"
	"time"
)

const (
	WeatherAPI = "https://api.open-meteo.com/v1/forecast"
	GeoAPI     = "https://geocoding-api.open-meteo.com/v1/search"
)

type APIClient struct {
	Client *http.Client
}

func NewAPIClient() *APIClient {
	return &APIClient{
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
