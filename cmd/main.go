package main

import (
	"TodoListApi/api"
	"TodoListApi/api/handler/todo"
	"TodoListApi/api/handler/user"
	"TodoListApi/api/router"
	"TodoListApi/internal/config"
	"TodoListApi/internal/database"
	"TodoListApi/internal/service"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var server http.Server

func init() {
	err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	err = database.ConnectDatabase(*config.Cfg)
	if err != nil {
		panic(err)
	}

	validator2 := validator.New()

	userService := service.NewUserService(database.DB)
	userHandler := user.NewUserHandler(userService, validator2)

	todoService := service.NewTodoService(database.DB)
	todoHandler := todo.NewTodoHandler(todoService, validator2)

	route := router.CreateNewRouter(userHandler, todoHandler)
	server = api.CreateNewServer(route)

}

func main() {
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
