package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	APP_PORT string
	DB_URL   string
	DB_NAME  string
)

func LoadEnv() {
	godotenv.Load()

	databaseURI := os.Getenv("DB_URL")
	if databaseURI == "" {
		log.Fatal("DB_URL is required")
	}

	databaseName := os.Getenv("DB_NAME")
	if databaseName == "" {
		log.Fatal("DB_NAME is required")
	}

	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		port = 8080
	}

	APP_PORT = fmt.Sprintf("%d", port)
	DB_URL = databaseURI
	DB_NAME = databaseName
}
