package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetServerPort() string {
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		log.Fatal("Port Server has not been configured")
	}
	return serverPort
}

func GetDBURL() string {
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		log.Fatal("Database User has not been configured")
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		log.Fatal("Database Host has not been configured")
	}

	dbPortStr := os.Getenv("DB_PORT")
	if dbPortStr == "" {
		log.Fatal("Database Port has not been configured")
	}
	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		log.Fatal("Invalid Database Port configuration")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal("Database Name has not been configured")
	}

	dbPass := os.Getenv("DB_PASS")

	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	return dbURL
}
