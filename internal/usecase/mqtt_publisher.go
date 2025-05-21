package usecase

import (
	"encoding/json"
	"log"

	"github.com/erwinprimadani/fleet-management/internal/models"
)

func (uc *LocationUsecase) SendDataToMQTT(location models.Location) error {
	payload, err := json.Marshal(location)
	if err != nil {
		return err
	}

	err = uc.Service.PublishLocationToMQTT(location.VehicleID, payload)
	if err != nil {
		log.Println("Publish to MQTT got err:", err)
		return err
	}

	return nil
}
