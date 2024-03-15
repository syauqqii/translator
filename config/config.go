package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// LoadConfig => load configuration from environment variables
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(" ! ERROR: Error loading '.env' file")
	}

	server := os.Getenv("SERVER_APP")
	if server == "" {
		log.Fatal(" ! SERVER is not set in '.env' file")
	}

	port := os.Getenv("PORT_APP")
	if port == "" {
		log.Fatal(" ! ERROR: PORT is not set in '.env' file")
	}
}
