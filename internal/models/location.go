package models

type Location struct {
	VehicleID string  `json:"vehicle_id" db:"vehicle_id"`
	Latitude  float64 `json:"latitude" db:"latitude"`
	Longitude float64 `json:"longitude" db:"longitude"`
	Timestamp string  `json:"timestamp" db:"timestamp"`
}
