package config

import (
	"errors"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Server struct {
	Addr string `yaml:"addr" env-required:"true"`
}

type Config struct {
	Env    string `yaml:"env" env-required:"true"`
	Server Server `yaml:"server" env-required:"true"`
}

func Get() (*Config, error) {
	config := os.Getenv("CONFIG")
	if config == "" {
		return nil, errors.New("config.Get: CONFIG is not set")
	}

	configPath := fmt.Sprintf("./configs/%s.yml", config)
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config.Get: CONFIG is not valid: file does not exist: %s", configPath)
	}

	cfg := new(Config)

	if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
		return nil, fmt.Errorf("config.Get: %w", err)
	}

	return cfg, nil
}
