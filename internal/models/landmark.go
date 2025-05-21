package models

type Landmark struct {
	ID           int64   `json:"id" db:"id"`
	Code         string  `json:"code" db:"code"`
	LandmarkName string  `json:"landmark_name" db:"landmark_name"`
	Latitude     float64 `json:"latitude" db:"latitude"`
	Longitude    float64 `json:"longitude" db:"longitude"`
	CreatedDate  int64   `json:"created_date" db:"created_date"`
	CreatedBy    string  `json:"created_by" db:"created_by"`
}

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
