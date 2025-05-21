package usecase

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/erwinprimadani/fleet-management/internal/models"
)

func (uc *LocationUsecase) PublishGeofenceEvent(loc models.Location) error {
	log.Println("####PublishGeofenceEvent message", loc)
	msg := map[string]interface{}{
		"vehicle_id": loc.VehicleID,
		"event":      "geofence_entry",
		"location": map[string]float64{
			"latitude":  loc.Latitude,
			"longitude": loc.Longitude,
		},
		"timestamp": loc.Timestamp,
	}

	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = uc.Service.PublishGeofenceEventToRabbitMQ(fmt.Sprintf("geofence.%s", loc.VehicleID), body)

	return err
}
