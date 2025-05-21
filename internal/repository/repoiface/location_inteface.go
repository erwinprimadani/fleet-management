package repoiface

import (
	"context"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/erwinprimadani/fleet-management/internal/models"
	amqp "github.com/rabbitmq/amqp091-go"
)

type DBRepository interface {
	SaveLocation(loc models.Location) error
	GetLatestLocation(vehicleID string) (*models.Location, error)
	GetLocationHistory(vehicleID, start, end string) ([]models.Location, error)
	GetAllLandmarks(ctx context.Context) ([]models.Landmark, error)
	GetLandmarkByCode(ctx context.Context, code string) (*models.Landmark, error)
}

type MQTTRepository interface {
	PublishLocation(vehicleID string, payload []byte) error
	Subscribe(topic string, qos byte, callback mqtt.MessageHandler) mqtt.Token
}

type RabbitMQRepository interface {
	Publish(routingKey string, message []byte) error
	Consume() (<-chan amqp.Delivery, error)
}
