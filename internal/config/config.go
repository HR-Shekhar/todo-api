package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Database DBConfig
}

type DBConfig struct {
	Name string
	Host string
	User string
	Password string
	Port string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load env")
	}
	mainConfig := &Config{}
	dbconfig := &DBConfig{}

	dbconfig.Name = os.Getenv("DB_NAME")
	if dbconfig.Name == "" {
		log.Fatal("DB_NAME is empty")
	}
	dbconfig.Host = os.Getenv("DB_HOST")
	if dbconfig.Host == "" {
		log.Fatal("DB_HOST is empty")
	}
	dbconfig.User = os.Getenv("DB_USER")
	if dbconfig.User == "" {
		log.Fatal("DB_USER is empty")
	}
	dbconfig.Password = os.Getenv("DB_PASSWORD")
	if dbconfig.Password == "" {
		log.Fatal("DB_PASSWORD is empty")
	}
	dbconfig.Port = os.Getenv("DB_PORT")
	if dbconfig.Port == "" {
		log.Fatal("DB_PORT is empty")
	}

	mainConfig.Port = os.Getenv("PORT")
	if mainConfig.Port == "" {
		log.Fatal("PORT is empty")
	}
	mainConfig.Database = *dbconfig
	return mainConfig
}
