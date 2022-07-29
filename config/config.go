package config

import (
	"log"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

var config Config

func init() {
	err := godotenv.Load("./config/conf.yaml")
	e := os.Getenv("env")
	if e != "production" {
		if err != nil {
			log.Fatal("Error on load configuration file.")
		}
	}

	if err := env.Parse(&config); err != nil {
		log.Fatal("Error on parsing configuration file.", err)
	}

	log.Printf(`
		env: %s | port: %d
		token_secret_key: %s
		user_service_endpoint: %s
	`,
		config.Environment, config.Port,
		config.TokenSecretKey,
		config.UserServiceEndpoint,
	)
}

// GetConfig returns current config
func GetConfig() *Config {
	return &config
}

// Config : struct
type Config struct {
	Environment        string `json:"env" env:"env"`
	Port               int    `json:"port" env:"port"`
	TokenSecretKey     string `json:"token_secret_key" env:"token_secret_key"`
	InternalSecuredKey string `json:"internal_secured_key" env:"internal_secured_key"`

	// micro services
	UserServiceEndpoint string `json:"user_service_endpoint" env:"user_service_endpoint"`
}
