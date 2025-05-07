package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"weather-app/internal/models"
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

func (c *APIClient) GeocodeLocation(location string) (*models.GeoResponse, error) {
	url := fmt.Sprintf("%s?name=%s&count=1", GeoAPI, location)

	resp, err := c.Client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("geocoding API failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("geocoding API returned %d", resp.StatusCode)
	}
	var result models.GeoResponse
	if err := json.NewDecoder(resp.Body).Decode(&resp); err != nil {
		return nil, fmt.Errorf("failed to decode geocoding response: %w", err)
	}
	return &result, nil
}
