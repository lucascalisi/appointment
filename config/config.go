package config

import (
	"fmt"
	"os"
	"strconv"
)

type DBConfig struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     int
	DBName     string
}

type EmailSenderCfg struct {
	User             string
	MySecretPassword string
	SMTPServer       string
	Port             string
}

type Config struct {
	DBCfg               DBConfig
	Debug               bool
	AppointmentDuration int
	EmailSenderCfg      EmailSenderCfg
}

func GetConfig() (Config, error) {
	// Get environment config parameters
	debugEnv := true

	dbCfg, err := getDBConfig()
	if err != nil {
		return Config{}, err
	}

	esCfg, err := getEmailSenderConfig()
	if err != nil {
		return Config{}, err
	}
	return Config{
		DBCfg:               dbCfg,
		AppointmentDuration: 30,
		Debug:               debugEnv,
		EmailSenderCfg:      esCfg,
	}, nil
}

func getDBConfig() (DBConfig, error) {
	dbCfg := DBConfig{}
	var dbPortEnvStr string
	dbPortDefault := 3306

	dbCfg.DBUser = os.Getenv("DBUSER")
	dbCfg.DBPassword = os.Getenv("DBPASSWORD")
	dbCfg.DBName = os.Getenv("DBNAME")
	dbCfg.DBHost = os.Getenv("DBHOST")
	dbPortEnvStr = os.Getenv("DBPORT")

	if dbPortEnvStr != "" {
		var err error
		dbCfg.DBPort, err = strconv.Atoi(dbPortEnvStr)
		if err != nil {
			return DBConfig{}, fmt.Errorf("got invalid DB port from environment")
		}
	} else {
		dbCfg.DBPort = dbPortDefault
	}

	return dbCfg, nil
}

func getEmailSenderConfig() (EmailSenderCfg, error) {
	esCfg := EmailSenderCfg{}

	esCfg.MySecretPassword = os.Getenv("SMTP_PASSWORD")
	esCfg.User = os.Getenv("SMTP_USER")
	esCfg.SMTPServer = os.Getenv("SMTP_SERVER")
	esCfg.Port = os.Getenv("SMTP_PORT")

	return esCfg, nil
}
