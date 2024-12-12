package config

import (
	"TodoListApi/entities"
	"github.com/joho/godotenv"
	"os"
)

var Cfg *entities.Config

func LoadConfig() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	Cfg = &entities.Config{
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		AccessKey:  os.Getenv("ACCESS_KEY"),
	}
	return nil
}
