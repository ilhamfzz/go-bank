package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Get() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error When load environment %s", err.Error())
	}

	return &Config{
		Server: Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Database: Database{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
		Email: Email{
			Host:     os.Getenv("MAIL_HOST"),
			Port:     os.Getenv("MAIL_PORT"),
			Username: os.Getenv("MAIL_USER"),
			Password: os.Getenv("MAIL_PASSWORD"),
		},
		Redis: Redis{
			Addr: os.Getenv("REDIS_ADDR"),
			Pass: os.Getenv("REDIS_PASS"),
		},
	}
}
