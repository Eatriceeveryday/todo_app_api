package service

import (
	"TodoListApi/entities"
	"database/sql"
	"fmt"
)

type TodoService struct {
	DB *sql.DB
}

func NewTodoService(db *sql.DB) TodoService {
	return TodoService{DB: db}
}

func (c *TodoService) AddTodo(title string, userId string) error {
	_, err := c.DB.Query("INSERT INTO todo (title , user_id) VALUES ($1 , $2) ", title, userId)
	if err != nil {
		return err
	}
	return nil
}

func (c *TodoService) DeleteTodo(todoId string) error {
	_, err := c.DB.Query("DELETE FROM todo WHERE todo_id =  $1", todoId)
	if err != nil {
		return err
	}
	return nil
}

func (c *TodoService) GetTodo(userId string) ([]entities.Todo, error) {
	todos := []entities.Todo{}
	rows, err := c.DB.Query("SELECT todo_id, title FROM todo WHERE user_id = $1", userId)
	if err != nil {
		return todos, err
	}

	defer rows.Close()
	for rows.Next() {
		var todo entities.Todo

		err := rows.Scan(&todo.TodoId, &todo.Title)
		if err != nil {
			return todos, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (c *TodoService) CheckOwner(todoId string, userId string) (bool, error) {
	todo := entities.Todo{}

	rows, err := c.DB.Query("SELECT user_id FROM todo WHERE todo_id = $1", todoId)
	if err != nil {
		return false, err
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&todo.UserId)
		if err != nil {
			return false, err
		}
	}

	fmt.Println("user id todo:", todo.UserId)
	fmt.Println("user id", userId)
	if todo.UserId != userId {
		return false, nil
	}

	return true, nil
}
