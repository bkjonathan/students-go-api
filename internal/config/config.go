package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string `yaml:"storage_path" env:"STORAGE_PATH" env-required:"true" env-default:"storage/storage.db"`
	Server      struct {
		Host string `yaml:"host" env:"SERVER_HOST" env-default:"localhost"`
		Port int    `yaml:"port" env:"SERVER_PORT" env-required:"true" env-default:"8082"`
	} `yaml:"server"`
}

func LoadConfig() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "path to the configuration file")
		flag.Parse()
		configPath = *flags

		if configPath == "" {
			log.Fatal("Configuration file path must be provided via CONFIG_PATH environment variable or --config flag")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Configuration file not found at path: %s", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("Error reading configuration file: %v", err)
	}

	return &cfg
}
