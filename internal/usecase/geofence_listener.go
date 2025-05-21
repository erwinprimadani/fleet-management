package usecase

import (
	"encoding/json"
	"log"

	"github.com/erwinprimadani/fleet-management/internal/models"
)

func (u *LocationUsecase) GeofenceMessage() {
	msgs, err := u.Service.RabbitMQConsume()
	if err != nil {
		log.Fatalf("Failed to consume geofence_alerts: %v", err)
	}
	log.Println("Geofence worker is running...")
	for msg := range msgs {
		var event models.GeofenceEvent
		if err := json.Unmarshal(msg.Body, &event); err != nil {
			log.Printf("Invalid geofence event: %v", err)
			continue
		}

		log.Printf("####GEOFENCE ALERT: Kendaraan %s masuk radius di lokasi (%f, %f) pada %s",
			event.VehicleID,
			event.Location.Latitude,
			event.Location.Longitude,
			event.Timestamp,
		)
	}
}
