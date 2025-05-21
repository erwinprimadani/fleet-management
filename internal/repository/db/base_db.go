package db

import (
	"fmt"

	"database/sql"

	"github.com/erwinprimadani/fleet-management/internal/repository/repoconf"
	_ "github.com/lib/pq"
)

type databaseConf struct {
	Host         string
	Username     string
	Password     string
	SslMode      string
	Port         int
	DatabaseName string
}

func NewDatabaseConf(host, username, password, sslMode, databaseName string, port int) repoconf.RepoConf {
	return &databaseConf{
		Host:         host,
		Username:     username,
		Password:     password,
		SslMode:      sslMode,
		Port:         port,
		DatabaseName: databaseName,
	}
}

func (conf *databaseConf) GetRepoName() string {
	return "Database"
}

func (conf *databaseConf) Init(r *repoconf.Repository) error {

	connString := getConnectionString(conf.Host, conf.Username, conf.Password, conf.SslMode, conf.DatabaseName, conf.Port)

	dbConn, err := sql.Open("postgres", connString)
	if err != nil {
		return err
	}

	r.DB = NewLocationRepository(dbConn)

	return nil
}

func getConnectionString(host, userName, password, sslMode, databaseName string, port int) string {
	connectionStringTemplate := "host=%s port=%d sslmode=%s user=%s password='%s' dbname=%s "

	return fmt.Sprintf(
		connectionStringTemplate,
		host,
		port,
		sslMode,
		userName,
		password,
		databaseName)
}
