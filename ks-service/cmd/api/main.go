package main

import (
	"log"
	"time"

	"go.uber.org/zap"
)

// @TODO: Move it to constants
const version = "0.0.1" // Pre release

func main() {
	// Config for the application
	// @TODO: Pick from env
	cfg := AppConfig{
		Addr: "8080",
		DB: DBConfig{
			URI:            "mongodb://localhost:27017",
			Database:       "knowledgeshare",
			ConnectTimeout: 10 * time.Second,
			ServerTimeout:  5 * time.Second,
		},
	}

	// Logger Setup
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	appInstance := &KSapplication{
		Config: cfg,
		Logger: logger,
	}

	router := appInstance.mountRoutes()

	if err := appInstance.run(router); err != nil {
		log.Fatal(err) // @TODO: Use the logger
	}
}
