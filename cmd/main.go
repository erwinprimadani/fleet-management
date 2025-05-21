package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/erwinprimadani/fleet-management/config"
	"github.com/erwinprimadani/fleet-management/internal/handler"
	"github.com/erwinprimadani/fleet-management/internal/repository"
	"github.com/erwinprimadani/fleet-management/internal/service"
	"github.com/erwinprimadani/fleet-management/internal/usecase"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	godotenv.Load()
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	repository.LoadRepository(cfg)

	repo := repository.GetRepo()
	svc := service.NewLocationService(repo.DB, repo.MQTT, repo.RabbitMQ)
	uc := usecase.NewLocationUsecase(svc)
	h := handler.NewLocationHandler(uc)

	r := gin.Default()
	r.GET("/vehicles/:vehicle_id/location", h.GetLatestLocation)
	r.GET("/vehicles/:vehicle_id/history", h.GetLocationHistory)
	r.POST("/vehicles/sent/location", h.SendLatestLocation)
	r.GET("/healthcheck", h.Healthcheck)

	broker := fmt.Sprintf("tcp://%s:%d", cfg.MQTT.Broker, cfg.MQTT.Port)
	log.Println("StartMQTTSubscriber is running...", broker)
	go uc.StartMQTTSubscriber(broker)

	go uc.GeofenceMessage()

	log.Println("Server running at :3000")
	http.ListenAndServe(":3000", r)
}
