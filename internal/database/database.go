package database

import (
	"TodoListApi/entities"
	"database/sql"
	"fmt"
)

var DB *sql.DB

func ConnectDatabase(config entities.Config) error {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUsername, config.DBPassword, config.DBName)
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	fmt.Println("Connected to database")
	return nil
}
