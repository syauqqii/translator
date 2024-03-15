package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	SERVER_APP string
	PORT_APP   string
)

// LoadConfig => load configuration from environment variables
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(" ! ERROR: Error loading '.env' file")
	}

	SERVER_APP = os.Getenv("SERVER_APP")
	if SERVER_APP == "" {
		log.Fatal(" ! SERVER is not set in '.env' file")
	}

	PORT_APP = os.Getenv("PORT_APP")
	if PORT_APP == "" {
		log.Fatal(" ! ERROR: PORT is not set in '.env' file")
	}
}
