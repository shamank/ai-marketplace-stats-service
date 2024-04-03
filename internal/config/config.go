package config

import (
	"errors"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"os"
	"time"
)

const defaultConfigPath = "./configs/prod.yaml"

var ErrNoConfig = errors.New("config file not found")

type (
	GRPCConfig struct {
		Port    int           `yaml:"port"`
		Timeout time.Duration `yaml:"timeout"`
	}

	PostgresConfig struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `env:"POSTGRES_PASSWORD" env-required:"true"`
		DBName   string `yaml:"dbname"`
		SSLMode  string `yaml:"ssl-mode"`
	}

	Config struct {
		GRPC     GRPCConfig     `yaml:"grpc"`
		Postgres PostgresConfig `yaml:"postgres"`
	}
)

func LoadConfig(configPath string) (*Config, error) {

	// Почему-то без этого вызова не работает, хотя раньше cleanvenv вроде сам справлялся  :(
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	if configPath == "" {
		if configPathEnv := os.Getenv("CONFIG_PATH"); configPathEnv != "" {
			configPath = configPathEnv
		} else {
			configPath = defaultConfigPath
		}
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
