package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"time"
)

const (
	targetLat  = -6.2430015
	targetLong = 106.8246234
)

type VehicleLocation struct {
	VehicleID string  `json:"vehicle_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timestamp string  `json:"timestamp"`
}

func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371000 // Radius bumi dalam meter
	dLat := (lat2 - lat1) * math.Pi / 180
	dLon := (lon2 - lon1) * math.Pi / 180
	lat1 = lat1 * math.Pi / 180
	lat2 = lat2 * math.Pi / 180

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(lat1)*math.Cos(lat2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}

func vehicleLocationHandler(w http.ResponseWriter, r *http.Request) {
	var loc VehicleLocation
	if err := json.NewDecoder(r.Body).Decode(&loc); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	distance := haversine(loc.Latitude, loc.Longitude, targetLat, targetLong)
	fmt.Printf("Received location: %+v | Distance to target: %.2f meters\n", loc, distance)
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/vehicles/sent/location", vehicleLocationHandler)
	go func() {
		log.Println("Starting server on :8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	startLat := targetLat + 0.001
	startLong := targetLong + 0.001
	lat := startLat
	long := startLong
	steps := 40

	for i := 0; i < steps; i++ {
		timestamp := time.Now().Format(time.RFC3339)
		data := VehicleLocation{
			VehicleID: "AA123QWE",
			Latitude:  lat,
			Longitude: long,
			Timestamp: timestamp,
		}
		body, _ := json.Marshal(data)
		_, err := http.Post("http://localhost:3000/vehicles/sent/location", "application/json", bytes.NewBuffer(body))
		if err != nil {
			log.Println("Failed to send:", err)
		}

		deltaLat := (targetLat - lat) / float64(steps-i)
		deltaLong := (targetLong - long) / float64(steps-i)

		lat += deltaLat + (rand.Float64()-0.5)*0.00001
		long += deltaLong + (rand.Float64()-0.5)*0.00001

		if haversine(lat, long, targetLat, targetLong) < 50 {
			log.Println("Target location reached within 50 meters.")
		}

		time.Sleep(2 * time.Second)
	}

}
