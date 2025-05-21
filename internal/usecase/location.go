package usecase

import (
	"github.com/erwinprimadani/fleet-management/internal/models"
	"github.com/erwinprimadani/fleet-management/internal/service"
)

type LocationUsecase struct {
	Service service.LocationService
}

func NewLocationUsecase(service service.LocationService) *LocationUsecase {
	return &LocationUsecase{Service: service}
}

func (u *LocationUsecase) SaveLocation(loc models.Location) error {
	return u.Service.SaveLocation(loc)
}

func (u *LocationUsecase) GetLatest(vehicleID string) (*models.Location, error) {
	return u.Service.GetLatestLocation(vehicleID)
}

func (u *LocationUsecase) GetHistory(vehicleID, start, end string) ([]models.Location, error) {
	return u.Service.GetLocationHistory(vehicleID, start, end)
}
