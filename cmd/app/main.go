package main

import (

	// Logging
	// Web server
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"weather-app/internal/api"
)

func main() {
	client := api.NewAPIClient()

	http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		location := r.URL.Query().Get("location")

		geoResp, err := client.GeocodeLocation(location)
		if err != nil || len(geoResp.Results) == 0 {
			http.Error(w, "Location not found", http.StatusBadRequest)
			return
		}
		weatherResp, err := client.GetWeather(
			geoResp.Results[0].Latitude,
			geoResp.Results[0].Longitude,
		)
		if err != nil {
			http.Error(w, "weather data unavailable", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(weatherResp)
	})

	http.Handle("/", http.FileServer(http.Dir("./assets")))

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe("8080", nil))
}
