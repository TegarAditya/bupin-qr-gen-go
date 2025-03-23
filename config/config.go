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

	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	appPort, _ := strconv.Atoi(os.Getenv("PORT"))

	return Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     dbPort,
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		Port:       appPort,
	}
}
