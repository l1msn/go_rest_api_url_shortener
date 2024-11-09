package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string `yaml:"env" env-default:"local"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func EnvLoad() {
	err := godotenv.Load("local.env")

	if err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		log.Panic("CONFIG_PATH environment variable not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Panicf("CONFIG_PATH does not exist: %s", err.Error())
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Error reading config file: %s", err.Error())
	}

	return &cfg
}
