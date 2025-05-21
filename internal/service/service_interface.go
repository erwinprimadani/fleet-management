package service

import (
	"context"

	"github.com/erwinprimadani/fleet-management/internal/models"
	amqp "github.com/rabbitmq/amqp091-go"
)

type LocationService interface {
	SaveLocation(loc models.Location) error
	GetLatestLocation(vehicleID string) (*models.Location, error)
	GetLocationHistory(vehicleID, start, end string) ([]models.Location, error)
	PublishLocationToMQTT(vehicleID string, payload []byte) error
	PublishGeofenceEventToRabbitMQ(vehicleID string, payload []byte) error
	GetAllLandmarks(ctx context.Context) ([]models.Landmark, error)
	GetLandmarkByCode(ctx context.Context, code string) (*models.Landmark, error)
	RabbitMQConsume() (<-chan amqp.Delivery, error)
}
