package router

import (
	"TodoListApi/api/handler/todo"
	"TodoListApi/api/handler/user"
	"TodoListApi/api/middleware"
	"net/http"
)

func CreateNewRouter(uh user.UserHandler, th todo.TodoHandler) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("POST /register", uh.Register)
	router.HandleFunc("POST /login", uh.Login)

	protectedRoute := http.NewServeMux()
	protectedRoute.HandleFunc("POST /todo", th.CreateTodo)
	protectedRoute.HandleFunc("GET /todo", th.GetListTodo)
	protectedRoute.HandleFunc("DELETE /todo", th.DeleteTodo)

	router.Handle("/", middleware.AuthenticateToken(protectedRoute))

	return router
}
