package todo

import (
	"TodoListApi/entities"
	"TodoListApi/internal/service"
	"TodoListApi/utils"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type TodoHandler struct {
	todoService service.TodoService
	validator   *validator.Validate
}

func NewTodoHandler(todoService service.TodoService, validator *validator.Validate) TodoHandler {
	return TodoHandler{
		todoService: todoService,
		validator:   validator,
	}
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id").(string)
	var req TodoRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.JSONResponse(w, utils.Response{Msg: err.Error()}, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	err = h.validator.Struct(req)
	if err != nil {
		utils.JSONResponse(w, utils.Response{Msg: err.Error()}, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	err = h.todoService.AddTodo(req.Title, userId)
	if err != nil {
		utils.JSONResponse(w, utils.Response{Msg: err.Error()}, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	utils.JSONResponse(w, utils.Response{Msg: "Todo Succesfuly Added"}, http.StatusCreated)
}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id").(string)
	todoId := r.URL.Query().Get("todo_id")

	if todoId == "" {
		utils.JSONResponse(w, utils.Response{Msg: "missing todo id"}, http.StatusBadRequest)
		return
	}

	valid, err := h.todoService.CheckOwner(todoId, userId)
	if err != nil {
		utils.JSONResponse(w, utils.Response{Msg: err.Error()}, http.StatusBadRequest)
		return
	}
	
	if !valid {
		utils.JSONResponse(w, utils.Response{Msg: "todo not owned"}, http.StatusBadRequest)
		return
	}

	err = h.todoService.DeleteTodo(todoId)
	if err != nil {
		utils.JSONResponse(w, utils.Response{Msg: err.Error()}, http.StatusBadRequest)
		return
	}
	utils.JSONResponse(w, utils.Response{Msg: "Todo Deleted"}, http.StatusOK)
}

func (h *TodoHandler) GetListTodo(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id").(string)

	todos, err := h.todoService.GetTodo(userId)
	if err != nil {
		utils.JSONResponse(w, utils.Response{Msg: err.Error()}, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	utils.JSONResponse(w, utils.Response{Msg: "Success", Data: struct {
		Todos []entities.Todo `json:"todos"`
	}{
		Todos: todos,
	}}, http.StatusOK)
}
