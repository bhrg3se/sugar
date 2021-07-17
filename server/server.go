package server

import (
	"flag"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sugar/migrations"
	"sugar/models"
	"sugar/routes"
	"sugar/store"
)

func StartServer() {
	path := flag.String("c", "/etc/sugar/config", "config file location")
	writeToFile := flag.Bool("f", false, "write logs to file")
	flag.Parse()

	config := ParseConfig(*path)
	level, err := logrus.ParseLevel(config.Logging.Level)
	if err != nil {
		level = logrus.ErrorLevel
	}
	logrus.SetLevel(level)
	logrus.SetReportCaller(true)

	if *writeToFile {
		f, err := os.OpenFile("/var/log/sugar.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
		if err != nil {
			panic(err)
		}
		logrus.SetOutput(f)
	}

	store.State = store.NewRealStore(config)

	err = migrations.Migrate(store.State.DB)
	if err != nil {
		logrus.Fatal(err)
	}

	r := routes.Routes()
	logrus.Infof("starting server at: %s", config.Server.Listen)
	logrus.Error(http.ListenAndServe(config.Server.Listen, r))
}

// ParseConfig uses viper to parse config file.
func ParseConfig(path string) models.Config {
	var config models.Config
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic("config file not found in " + filepath.Join(path))
	}

	viper.SetConfigName("config")
	viper.AddConfigPath(absPath)

	if err = viper.ReadInConfig(); err != nil {
		if strings.Contains(err.Error(), "Not Found in") {
			erro := os.MkdirAll(absPath, os.ModePerm)
			if erro != nil {
				logrus.Fatalf("could not create directory %s: %v", absPath, erro)
			}
			f, erro := os.Create(filepath.Join(absPath, "config.toml"))
			if erro != nil {
				logrus.Fatalf("could not create config file: %v", erro)
			}
			f.Close()

			err = viper.WriteConfig()
			if err != nil {
				logrus.Fatal(err)
			}
			logrus.Fatalf("could not find log file, created one at %s ", absPath)
		}
		logrus.Fatal(err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		logrus.Fatalf("config file invalid: %v", err)
	}

	return config
}
