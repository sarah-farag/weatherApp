package api

import (
	"encoding/json"
	"fmt"
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

func (c *APIClient) GetWeather(lat, lon float64) (*models.WeatherResponse, error) {
	url := fmt.Sprintf("%s?latitude=%.4f&longitude=%.4f&hourly=temperature_2m", WeatherAPI, lat, lon)

	resp, err := c.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result models.WeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	return &result, err
}
