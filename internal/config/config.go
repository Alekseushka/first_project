package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env		string `yaml:"env" env-required:"true"`
	HTTPServer 	   `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-required:"true"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
} 

func MustLoad() *Config {
	err := godotenv.Load("go.env") 
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH not set in .env file")
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Can't read config: %s", err)
	}

	return &cfg
}
