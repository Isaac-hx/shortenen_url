package config

import (
	"log"
	"os"
	"strconv"
)

// Type struct of app config
type AppConfig struct {
	AppName string
	AppPort int
}

// Type struct of DB config
type DbConfig struct {
	DbHost     string
	DbName     string
	DbPort     int
	DbUser     string
	DbPassword string
}

// Type struct of Config
type Config struct {
	AppConfig *AppConfig
	DbConfig  *DbConfig
}

// Constructur for object Config
func NewConfig() *Config {
	appPort, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatalf("[ERROR] App port is invalid :%v", err)
	}
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("[ERROR] Db port is invalid :%v", err)
	}

	//construct object config
	cfg := &Config{
		&AppConfig{
			AppName: os.Getenv("APP_NAME"),
			AppPort: appPort,
		},
		&DbConfig{
			DbHost:     os.Getenv("DB_HOST"),
			DbPort:     dbPort,
			DbName:     os.Getenv("DB_NAME"),
			DbUser:     os.Getenv("DB_USER"),
			DbPassword: os.Getenv("DB_PASSWORD"),
		},
	}
	return cfg
}
