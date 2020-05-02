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

type Config struct {
	DBCfg DBConfig
	Debug bool
}

func GetConfig() (Config, error) {
	// Get environment config parameters
	debugEnv := true

	dbCfg, err := getDBConfig()
	if err != nil {
		return Config{}, err
	}
	return Config{
		DBCfg: dbCfg,
		Debug: debugEnv,
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
