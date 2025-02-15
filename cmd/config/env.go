package config

import (
	"fmt"
	"os"
)

func LoadEnv() (*Config, error) {
	cfg := &Config{}

	var err error
	cfg.ServerPort, err = getEnv("SERVER_PORT")
	if err != nil {
		return nil, fmt.Errorf("SERVER_PORT: %w", err)
	}

	cfg.DBURL, err = getEnv("DB_URL")
	if err != nil {
		return nil, fmt.Errorf("DB_URL: %w", err)
	}

	cfg.AddressServiceURL, err = getEnv("ADDRESS_SERVICE_URL")
	if err != nil {
		return nil, fmt.Errorf("ADDRESS_SERVICE_URL: %w", err)
	}

	cfg.PostgresUser, err = getEnv("POSTGRES_USER")
	if err != nil {
		return nil, fmt.Errorf("POSTGRES_USER: %w", err)
	}

	cfg.PostgresPassword, err = getEnv("POSTGRES_PASSWORD")
	if err != nil {
		return nil, fmt.Errorf("POSTGRES_PASSWORD: %w", err)
	}

	cfg.PostgresDB, err = getEnv("POSTGRES_DB")
	if err != nil {
		return nil, fmt.Errorf("POSTGRES_DB: %w", err)
	}

	return cfg, nil
}

func getEnv(key string) (string, error) {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		return "", fmt.Errorf("enviroment variable not found: %s", key)
	}
	return value, nil
}
