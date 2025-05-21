package db

import (
	"database/sql"

	"github.com/erwinprimadani/fleet-management/internal/models"
	"github.com/erwinprimadani/fleet-management/internal/repository/repoiface"
)

type LocationRepoPG struct {
	db *sql.DB
}

func NewLocationRepository(db *sql.DB) repoiface.DBRepository {
	return &LocationRepoPG{db: db}
}

func (r *LocationRepoPG) SaveLocation(loc models.Location) error {
	_, err := r.db.Exec(`INSERT INTO vehicle_locations (vehicle_id, latitude, longitude, timestamp) VALUES ($1, $2, $3, $4)`,
		loc.VehicleID, loc.Latitude, loc.Longitude, loc.Timestamp)
	return err
}

func (r *LocationRepoPG) GetLatestLocation(vehicleID string) (*models.Location, error) {
	var loc models.Location
	err := r.db.QueryRow(`SELECT vehicle_id, latitude, longitude, timestamp FROM vehicle_locations WHERE vehicle_id = $1 ORDER BY timestamp DESC LIMIT 1`, vehicleID).
		Scan(&loc.VehicleID, &loc.Latitude, &loc.Longitude, &loc.Timestamp)
	if err != nil {
		return nil, err
	}
	return &loc, nil
}

func (r *LocationRepoPG) GetLocationHistory(vehicleID, start, end string) ([]models.Location, error) {
	rows, err := r.db.Query(`SELECT vehicle_id, latitude, longitude, timestamp FROM vehicle_locations WHERE vehicle_id = $1 AND timestamp BETWEEN $2 AND $3 ORDER BY timestamp`,
		vehicleID, start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var locations []models.Location
	for rows.Next() {
		var loc models.Location
		if err := rows.Scan(&loc.VehicleID, &loc.Latitude, &loc.Longitude, &loc.Timestamp); err != nil {
			return nil, err
		}
		locations = append(locations, loc)
	}
	return locations, nil
}
