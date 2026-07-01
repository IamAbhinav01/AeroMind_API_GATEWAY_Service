package controllers

import (
	"AeromindGO/services"
	"net/http"
)

type UserController struct {
	UserService services.UserService
}

func (user *UserController) Create(w http.ResponseWriter , r *http.Request){
	user.UserService.Create()
	w.Write([] byte("user fetching ongoing ..."))
}

func 

func NewUserController(_userService services.UserService) *UserController{
	return &UserController{
		UserService: _userService,
	}
} 