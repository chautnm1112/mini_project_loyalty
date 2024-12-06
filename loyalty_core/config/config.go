package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DBName   string
	Password string
	Username string
	Port     string
	DbHost   string
	Host     string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	c := &Config{
		DBName:   os.Getenv("DB_DATABASE"),
		Password: os.Getenv("DB_PASSWORD"),
		Username: os.Getenv("DB_USERNAME"),
		Port:     os.Getenv("DB_PORT"),
		DbHost:   os.Getenv("DB_HOST_LOYALTY_CORE"),
		Host:     os.Getenv("HOST_LOYALTY_CORE"),
	}

	return c
}
