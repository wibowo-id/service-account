package config

import (
	"flag"
	"fmt"
	"os"
	"service-account/helper"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	APIHost    string
	APIPort    int
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		helper.LogCritical("Failed to load .env file", err)
	}
	helper.LogInfo(".env file loaded")

	var host string
	var port int
	flag.StringVar(&host, "host", "0.0.0.0", "API host")
	flag.IntVar(&port, "port", 3000, "API port")
	flag.Parse()

	return Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
		APIHost:    host,
		APIPort:    port,
	}
}

func ConnectDBWithRetry(dsn string, maxAttempts int, delay time.Duration) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			helper.LogInfo("Connected to database", nil)
			return db, nil
		}
		helper.LogWarning(fmt.Sprintf("Attempt %d: Failed to connect to DB, retrying in %s", attempt, delay), err)
		time.Sleep(delay)
	}
	return nil, err
}
