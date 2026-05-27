package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load env")
	}
	mainConfig := &Config{}

	mainConfig.Port = os.Getenv("PORT")
	return mainConfig
}
