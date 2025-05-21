package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/erwinprimadani/fleet-management/internal/models"
)

func (r *LocationRepoPG) GetAllLandmarks(ctx context.Context) ([]models.Landmark, error) {
	query := `SELECT id, code, landmark_name, latitude, longitude, created_date, created_by FROM location_landmark`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query landmarks: %w", err)
	}
	defer rows.Close()

	var landmarks []models.Landmark
	for rows.Next() {
		var landmark models.Landmark
		var createdDate time.Time

		err := rows.Scan(
			&landmark.ID,
			&landmark.Code,
			&landmark.LandmarkName,
			&landmark.Latitude,
			&landmark.Longitude,
			&createdDate,
			&landmark.CreatedBy,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan landmark row: %w", err)
		}

		landmark.CreatedDate = createdDate.Unix()
		landmarks = append(landmarks, landmark)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after iterating landmark rows: %w", err)
	}

	return landmarks, nil
}

func (r *LocationRepoPG) GetLandmarkByCode(ctx context.Context, code string) (*models.Landmark, error) {
	query := `SELECT id, code, landmark_name, latitude, longitude, created_date, created_by FROM location_landmark WHERE code = $1`

	var landmark models.Landmark
	var createdDate time.Time

	err := r.db.QueryRowContext(ctx, query, code).Scan(
		&landmark.ID,
		&landmark.Code,
		&landmark.LandmarkName,
		&landmark.Latitude,
		&landmark.Longitude,
		&createdDate,
		&landmark.CreatedBy,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Or a specific "not found" error
		}
		return nil, fmt.Errorf("failed to get landmark by code: %w", err)
	}

	landmark.CreatedDate = createdDate.Unix()
	return &landmark, nil
}
