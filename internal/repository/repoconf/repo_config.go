package repoconf

import (
	"github.com/erwinprimadani/fleet-management/internal/repository/repoiface"
)

var repo *Repository

type Repository struct {
	DB       repoiface.DBRepository
	MQTT     repoiface.MQTTRepository
	RabbitMQ repoiface.RabbitMQRepository
}

type RepoConf interface {
	Init(*Repository) error
	GetRepoName() string
}

func NewRepository(rf []RepoConf) (*Repository, error) {
	if repo != nil {
		return repo, nil
	}

	repo = &Repository{}
	for _, rc := range rf {
		err := rc.Init(repo)
		if err != nil {
			return nil, err
		}
	}

	return repo, nil
}
