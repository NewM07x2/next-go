package config

import (
	"os"
)

// Config holds application configuration
type Config struct {
	Port     string
	DBHost   string
	DBUser   string
	DBPass   string
	DBName   string
	FrontURL string
}

// Load loads configuration from environment variables
func Load() *Config {
	return &Config{
		Port:     getEnv("PORT", "8080"),
		DBHost:   getEnv("DB_HOST", "localhost"),
		DBUser:   getEnv("DB_USER", "postgres"),
		DBPass:   getEnv("DB_PASSWORD", "postgres"),
		DBName:   getEnv("DB_NAME", "next-go-sample-db"),
		FrontURL: getEnv("FRONTEND_URL", "http://localhost:3000"),
	}
}

// getEnv gets environment variable or returns default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
