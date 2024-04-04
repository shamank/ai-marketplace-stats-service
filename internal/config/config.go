package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

const defaultConfigPath = "./configs/prod.yaml" // Дефолтная папка конфигурации

type (
	//Конфигурация GRPC-сервера
	GRPCConfig struct {
		Port    int           `yaml:"port"`
		Timeout time.Duration `yaml:"timeout"`
	}

	//Конфигурация PostgresSQL
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

// LoadConfig - загрузка конфигурации из файла
func LoadConfig(configPath string) (*Config, error) {

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
