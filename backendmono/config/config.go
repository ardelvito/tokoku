package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadConfig() (*Config, error) {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file")
	}

	// Ambil value variabel env
	host := getEnv("DB_HOST", "localhost")
	port, err := getEnvAsInt("DB_PORT", 5432)
	if err != nil {
		return nil, err
	}
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "root")
	dbname := getEnv("DB_NAME", "tokokudb")

	// Logging debugging
	fmt.Printf("Loaded config: host=%s, port=%d, user=%s, password=%s, dbname=%s\n", host, port, user, password, dbname)

	return &Config{
		DBHost:     host,
		DBPort:     port,
		DBUser:     user,
		DBPassword: password,
		DBName:     dbname,
	}, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	fmt.Printf("Environment variable %s not found, using fallback: %s\n", key, fallback)
	return fallback
}

func getEnvAsInt(key string, fallback int) (int, error) {
	value, exists := os.LookupEnv(key)
	if !exists {
		fmt.Printf("Environment variable %s not found, using fallback: %d\n", key, fallback)
		return fallback, nil
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		fmt.Printf("Invalid integer value for environment variable %s, using fallback: %d\n", key, fallback)
		return fallback, fmt.Errorf("invalid integer value for environment variable %s: %v", key, err)
	}
	return intValue, nil
}
