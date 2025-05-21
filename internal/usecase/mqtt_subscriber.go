package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/erwinprimadani/fleet-management/internal/models"
	"github.com/erwinprimadani/fleet-management/internal/pkg/utils"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func (uc *LocationUsecase) StartMQTTSubscriber(broker string) {
	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID("fleet-subscriber")

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to connect to MQTT broker: %v", token.Error())
	}

	landmarks, err := uc.Service.GetAllLandmarks(context.Background())
	if err != nil {
		log.Printf("Error fetching landmarks: %v", err)
		return
	}

	topic := "/fleet/vehicle/+/location"
	client.Subscribe(topic, 1, func(client mqtt.Client, msg mqtt.Message) {
		var loc models.Location
		if err := json.Unmarshal(msg.Payload(), &loc); err != nil {
			log.Println("Invalid JSON:", err)
			return
		}

		if err := utils.ValidateLocation(loc); err != nil {
			log.Println("Validation error:", err)
			return
		}

		if err := uc.SaveLocation(loc); err != nil {
			log.Println("Failed to save location:", err)
		}

		for _, landmark := range landmarks {
			if utils.IsInGeofence(loc.Latitude, loc.Longitude, landmark.Latitude, landmark.Longitude, 50) {
				fmt.Println("Vehicle", loc.VehicleID, "entered geofence", landmark)
				err := uc.PublishGeofenceEvent(loc)
				if err != nil {
					log.Println("RabbitQ Failed to publish geofence event:", err)
				}
			}
		}
	})

	log.Println("MQTT subscriber started on topic:", topic)
}
