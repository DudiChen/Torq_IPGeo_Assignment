package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"time"
)

type Config struct {
	DatabaseType string
	DatabasePath string
	RateRequests int
	RateInterval time.Duration
	Port         string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("failed to load .env file")
	}

	databaseType := os.Getenv("DATABASE_TYPE")
	if databaseType == "" {
		return nil, errors.New("DATABASE_TYPE is not set")
	}

	databasePath := os.Getenv("DATABASE_PATH")
	if databasePath == "" {
		return nil, errors.New("DATABASE_PATH is not set")
	}

	rateRequestsStr := os.Getenv("RATE_REQUESTS")
	if rateRequestsStr == "" {
		return nil, errors.New("RATE_REQUESTS is not set")
	}
	rateRequests, err := strconv.Atoi(rateRequestsStr)
	if err != nil {
		return nil, errors.New("RATE_REQUESTS is not a valid integer")
	}

	rateIntervalStr := os.Getenv("RATE_INTERVAL")
	if rateIntervalStr == "" {
		return nil, errors.New("RATE_INTERVAL is not set")
	}
	rateInterval, err := time.ParseDuration(rateIntervalStr)
	if err != nil {
		return nil, errors.New("RATE_INTERVAL is not a valid duration")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	config := &Config{
		DatabaseType: databaseType,
		DatabasePath: databasePath,
		RateRequests: rateRequests,
		RateInterval: rateInterval,
		Port:         port,
	}

	return config, nil
}
