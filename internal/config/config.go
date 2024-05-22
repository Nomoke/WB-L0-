package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env"`
	PostgresUrl string `yaml:"postgres_url" env:"POSTGRES_URL,required"`
	NustURL     string `yaml:"nats_url" env:"NATS_URL,required"`
	Address string `yaml:"http_server"`
}

func LoadConfig() *Config {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		configPath = "/Users/air/Desktop/wb-test-app/config.yaml"
	}

	//chek if exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
