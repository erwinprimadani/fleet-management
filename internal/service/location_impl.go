package service

import (
	"context"

	"github.com/erwinprimadani/fleet-management/internal/models"
	"github.com/erwinprimadani/fleet-management/internal/repository/repoiface"
	amqp "github.com/rabbitmq/amqp091-go"
)

type locationService struct {
	db       repoiface.DBRepository
	mqtt     repoiface.MQTTRepository
	rabbitMQ repoiface.RabbitMQRepository
}

func NewLocationService(db repoiface.DBRepository, mqtt repoiface.MQTTRepository, rabbitMQ repoiface.RabbitMQRepository) LocationService {
	return &locationService{
		db:       db,
		mqtt:     mqtt,
		rabbitMQ: rabbitMQ,
	}
}

func (s *locationService) SaveLocation(loc models.Location) error {
	return s.db.SaveLocation(loc)
}

func (s *locationService) GetLatestLocation(vehicleID string) (*models.Location, error) {
	return s.db.GetLatestLocation(vehicleID)
}

func (s *locationService) GetLocationHistory(vehicleID, start, end string) ([]models.Location, error) {
	return s.db.GetLocationHistory(vehicleID, start, end)
}

func (s *locationService) GetAllLandmarks(ctx context.Context) ([]models.Landmark, error) {
	return s.db.GetAllLandmarks(ctx)
}

func (s *locationService) GetLandmarkByCode(ctx context.Context, code string) (*models.Landmark, error) {
	return s.db.GetLandmarkByCode(ctx, code)
}

func (s *locationService) PublishLocationToMQTT(vehicleID string, payload []byte) error {
	return s.mqtt.PublishLocation(vehicleID, payload)
}

func (s *locationService) PublishGeofenceEventToRabbitMQ(vehicleID string, payload []byte) error {
	return s.rabbitMQ.Publish(vehicleID, payload)
}

func (s *locationService) RabbitMQConsume() (<-chan amqp.Delivery, error) {
	return s.rabbitMQ.Consume()
}
