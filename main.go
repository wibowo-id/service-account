// main.go
package main

import (
	"fmt"
	"log"
	"os"
	conf "service-account/config"
	"service-account/helper"
	"service-account/model"
	"service-account/route"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	config := conf.LoadConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)

	var err error
	db, err = conf.ConnectDBWithRetry(dsn, 5, 3*time.Second)
	if err != nil {
		helper.LogCritical("Failed to connect to database after retries", map[string]any{"dsn": dsn, "error": err.Error()})
		os.Exit(1)
	}

	helper.LogInfo("Database connection established", map[string]any{"host": config.DBHost, "db": config.DBName})

	err = db.AutoMigrate(&model.Nasabah{}, &model.Rekening{}, &model.Mutasi{})
	if err != nil {
		helper.LogError("AutoMigrate failed", err)
	} else {
		helper.LogInfo("Database migration successful", "Nasabah")
	}

	app := fiber.New()
	helper.LogInfo("Fiber app initialized")

	route.SetupRoutes(app, db)

	addr := fmt.Sprintf("%s:%d", config.APIHost, config.APIPort)
	helper.LogInfo("Starting server", map[string]any{"address": addr})

	log.Fatal(app.Listen(addr))
}
