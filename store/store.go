package store

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"strings"
	"sugar/models"
	"time"
)

var State Store

type Store struct {
	DB     *sql.DB
	Config models.Config
}

func NewRealStore(config models.Config) Store {
	config.Acme.ApiURL = strings.TrimRight(config.Acme.ApiURL, "/")
	return Store{
		DB:     dbConn(config),
		Config: config,
	}
}

func dbConn(config models.Config) *sql.DB {

	var str string
	if config.Database.SSL {
		caCert, _ := filepath.Abs(config.Database.CaCertPath)
		userCert, _ := filepath.Abs(config.Database.UserCertPath)
		userKey, _ := filepath.Abs(config.Database.UserKeyPath)

		str = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=verify-full&sslrootcert=%s&sslcert=%s&sslkey=%s",
			config.Database.User,
			config.Database.Password,
			config.Database.Host,
			config.Database.Port,
			config.Database.Name,
			caCert,
			userCert,
			userKey)

	} else {
		str = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", config.Database.User,
			config.Database.Password,
			config.Database.Host,
			config.Database.Port,
			config.Database.Name)
	}

	db, err := sql.Open("postgres", str)
	if err != nil {
		panic(err.Error())
	}

	//Check if the connection is successful by establishing a connection.
	//Retry upto 10 times if connection is not successful
	for retryCount := 0; retryCount < 10; retryCount++ {
		err := db.Ping()
		if err == nil {
			fmt.Println("database connection successful")
			return db
		}
		logrus.Error(err)
		logrus.Error("could not connect to database: retrying...")
		time.Sleep(time.Second)
	}

	logrus.Fatal("could not connect to database")
	return nil
}
