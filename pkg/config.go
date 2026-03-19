// Package config handles application configuration
package pkg

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ENV                  string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found, using system environment variables")
	}

	return &Config{
		ENV:                  getEnv("ENV", "LOCAL"),
	}
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}