package mqtt

import (
	"github.com/erwinprimadani/fleet-management/internal/repository/repoconf"
	_ "github.com/lib/pq"
)

type mqttConf struct {
	broker   string
	port     int
	clientID string
	username string
	password string
	topic    string
}

func NewMQTTConf(broker, clientID, username, password, topic string, port int) repoconf.RepoConf {
	return &mqttConf{
		broker:   broker,
		port:     port,
		clientID: clientID,
		username: username,
		password: password,
		topic:    topic,
	}
}

func (conf *mqttConf) GetRepoName() string {
	return "MQTT"
}

func (conf *mqttConf) Init(r *repoconf.Repository) error {

	mqttClient, err := NewClient(conf)
	if err != nil {

	}
	r.MQTT = mqttClient

	return nil
}
