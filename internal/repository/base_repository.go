package repository

import (
	"log"

	"github.com/erwinprimadani/fleet-management/config"
	"github.com/erwinprimadani/fleet-management/internal/repository/db"
	"github.com/erwinprimadani/fleet-management/internal/repository/mqtt"
	"github.com/erwinprimadani/fleet-management/internal/repository/rabbitmq"
	"github.com/erwinprimadani/fleet-management/internal/repository/repoconf"
)

var repo *repoconf.Repository

func LoadRepository(cfg *config.Config) {
	repoList, err := repoconf.NewRepository([]repoconf.RepoConf{
		// add repository initialization here
		db.NewDatabaseConf(
			cfg.Database.Host,
			cfg.Database.User,
			cfg.Database.Password,
			cfg.Database.SSLMode,
			cfg.Database.DBName,
			cfg.Database.Port,
		),

		mqtt.NewMQTTConf(
			cfg.MQTT.Broker,
			cfg.MQTT.ClientID,
			cfg.MQTT.Username,
			cfg.MQTT.Password,
			cfg.MQTT.Topic,
			cfg.MQTT.Port,
		),
		rabbitmq.NewRabbitMQConf(
			cfg.RabbitMQ.Host,
			cfg.RabbitMQ.User,
			cfg.RabbitMQ.Password,
			cfg.RabbitMQ.Exchange,
			cfg.RabbitMQ.Queue,
			cfg.RabbitMQ.Port,
		),
	})
	if err != nil {
		log.Fatalf("cannot initiate repository, with error: %v", err)
	}
	repo = repoList
}

func GetRepo() *repoconf.Repository {
	return repo
}
