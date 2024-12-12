package service

import (
	"TodoListApi/entities"
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	DB *sql.DB
}

func NewUserService(db *sql.DB) UserService {
	return UserService{DB: db}
}

func (c UserService) AddUser(username string, password string) (string, error) {
	var userId string

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	err = c.DB.QueryRow("INSERT INTO users (username, password) VALUES ($1, $2) RETURNING user_id", username, string(hashedPassword)).Scan(&userId)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				return "", errors.New("User already exists")
			}
		}
		return "", err
	}

	return userId, nil
}

func (c UserService) GetUser(username string) (entities.User, error) {
	var user entities.User
	rows, err := c.DB.Query("SELECT * FROM users WHERE username = $1", username)
	if err != nil {
		return user, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.Uid, &user.Username, &user.Password)
		if err != nil {
			return user, err
		}
	}

	return user, nil
}
