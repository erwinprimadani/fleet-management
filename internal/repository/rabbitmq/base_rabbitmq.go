package rabbitmq

import (
	"github.com/erwinprimadani/fleet-management/internal/repository/repoconf"
)

type rabbitMQConf struct {
	host     string
	port     int
	user     string
	password string
	exchange string
	queue    string
}

func NewRabbitMQConf(host, user, password, exchange, queue string, port int) repoconf.RepoConf {
	return &rabbitMQConf{
		host:     host,
		user:     user,
		password: password,
		exchange: exchange,
		queue:    queue,
		port:     port,
	}
}

func (conf *rabbitMQConf) GetRepoName() string {
	return "RabbitMQ"
}

func (conf *rabbitMQConf) Init(r *repoconf.Repository) error {

	rabbitMQClient, err := NewClient(conf)
	if err != nil {
		return err
	}
	r.RabbitMQ = rabbitMQClient

	return nil
}
