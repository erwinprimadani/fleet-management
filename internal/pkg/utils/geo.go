package utils

import (
	"errors"
	"log"
	"math"
	"regexp"

	"github.com/erwinprimadani/fleet-management/internal/models"
)

var (
	pattern = `^[A-Z]+\d+[A-Z]+$`
)

func ValidateLocation(loc models.Location) error {
	if loc.VehicleID == "" {
		return errors.New("missing required fields")
	}

	if loc.Latitude == 0 || loc.Longitude == 0 {
		return errors.New("missing required fields")
	}

	matched, err := regexp.MatchString(pattern, loc.VehicleID)
	if err != nil || !matched {
		return errors.New("format vehicle_id not allowed")
	}

	return nil
}

func IsInGeofence(carLat, carLng, geofenceLat, geofenceLng, radiusMeters float64) bool {
	distance := calculateDistance(carLat, carLng, geofenceLat, geofenceLng)
	log.Println("####Distance from target = ", distance)
	return distance <= radiusMeters
}

// Haversine formula
func calculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371000 // Earth radius in meters
	dLat := (lat2 - lat1) * math.Pi / 180
	dLon := (lon2 - lon1) * math.Pi / 180
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1*math.Pi/180)*math.Cos(lat2*math.Pi/180)*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}
