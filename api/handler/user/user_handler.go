package user

import (
	"TodoListApi/internal/service"
	"TodoListApi/utils"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

type UserHandler struct {
	userService service.UserService
	validator   *validator.Validate
}

func NewUserHandler(userService service.UserService, validator *validator.Validate) UserHandler {
	return UserHandler{
		userService: userService,
		validator:   validator,
	}
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req UserRequest
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

	user, err := h.userService.GetUser(req.Username)
	if err != nil {
		utils.JSONResponse(w, utils.Response{Msg: err.Error()}, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		utils.JSONResponse(w, utils.Response{Msg: "Password Salah"}, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	token, err := createToken(user.Uid)
	if err != nil {
		utils.JSONResponse(w, utils.Response{Msg: "Internal Server Error "}, http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	utils.JSONResponse(w, utils.Response{
		Msg: "Login Success",
		Data: struct {
			Token string `json:"token"`
		}{
			Token: token,
		},
	}, http.StatusOK)
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req UserRequest

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

	_, err = h.userService.AddUser(req.Username, req.Password) //left return user id for future uses
	if err != nil {
		utils.JSONResponse(w, utils.Response{Msg: err.Error()}, http.StatusBadRequest)
		return
	}

	utils.JSONResponse(w, utils.Response{Msg: "Success"}, http.StatusCreated)
}

func createToken(userId string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userId,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
		"iat": time.Now().Unix(),
	})

	token, err := claims.SignedString([]byte(os.Getenv("ACCESS_KEY")))
	if err != nil {
		return "", err
	}

	return token, nil
}
