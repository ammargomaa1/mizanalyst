package config

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

// Config holds all environment-driven configuration values.
type Config struct {
	AppPort string

	// Database
	DBHost               string
	DBPort               string
	DBUser               string
	DBPassword           string
	DBName               string
	DBMaxOpenConns       int
	DBMaxIdleConns       int
	DBConnMaxLifetimeMin int

	// JWT
	AccessTokenSecret   string
	AccessTokenTTLMin   int
	RefreshTokenSecret  string
	RefreshTokenTTLDays int

	// Encryption
	EncryptionKey string
}

var (
	instance *Config
	once     sync.Once
)

// GetConfig returns the singleton Config instance.
// It loads the .env file and parses all required environment variables on first call.
func GetConfig() *Config {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Println("[CONFIG] Warning: .env file not found, falling back to OS environment variables")
		}

		instance = &Config{
			AppPort: getEnv("APP_PORT", "8080"),

			DBHost:               getEnv("DB_HOST", "localhost"),
			DBPort:               getEnv("DB_PORT", "5432"),
			DBUser:               getEnv("DB_USER", "postgres"),
			DBPassword:           getEnv("DB_PASSWORD", "postgres"),
			DBName:               getEnv("DB_NAME", "mizanalyst"),
			DBMaxOpenConns:       getEnvAsInt("DB_MAX_OPEN_CONNS", 25),
			DBMaxIdleConns:       getEnvAsInt("DB_MAX_IDLE_CONNS", 10),
			DBConnMaxLifetimeMin: getEnvAsInt("DB_CONN_MAX_LIFETIME_MIN", 30),

			AccessTokenSecret:   getEnv("ACCESS_TOKEN_SECRET", ""),
			AccessTokenTTLMin:   getEnvAsInt("ACCESS_TOKEN_TTL_MIN", 15),
			RefreshTokenSecret:  getEnv("REFRESH_TOKEN_SECRET", ""),
			RefreshTokenTTLDays: getEnvAsInt("REFRESH_TOKEN_TTL_DAYS", 7),

			EncryptionKey: getEnv("ENCRYPTION_KEY", ""),
		}
	})

	return instance
}

// getEnv reads an environment variable or returns a default value.
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// getEnvAsInt reads an environment variable as an integer or returns a default value.
func getEnvAsInt(key string, fallback int) int {
	strValue := getEnv(key, "")
	if strValue == "" {
		return fallback
	}

	value, err := strconv.Atoi(strValue)
	if err != nil {
		log.Printf("[CONFIG] Warning: invalid integer for %s, using default %d", key, fallback)
		return fallback
	}

	return value
}
