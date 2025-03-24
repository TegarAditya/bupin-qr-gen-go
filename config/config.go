package config

import (
	"os"
	"strconv"
)

type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	Port       int
}

func LoadConfig() Config {

	dbPort, _ := strconv.Atoi(coalesce(os.Getenv("DB_PORT"), "5432"))
	appPort, _ := strconv.Atoi(coalesce(os.Getenv("PORT"), "8080"))

	return Config{
		DBHost:     coalesce(os.Getenv("DB_HOST"), "localhost"),
		DBPort:     dbPort,
		DBUser:     coalesce(os.Getenv("DB_USER"), "user"),
		DBPassword: coalesce(os.Getenv("DB_PASSWORD"), "password"),
		DBName:     coalesce(os.Getenv("DB_NAME"), "dbname"),
		Port:       appPort,
	}
}

func coalesce(value, fallback string) string {
	if value == "" {
		return fallback
	}
	return value
}
