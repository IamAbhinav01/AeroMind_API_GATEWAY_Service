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

func (user *UserController) GetUserByID(w http.ResponseWriter,r *http.Request){
	user.UserService.GetUserByID(1)
	w.Write([] byte("user fetched for the id 2 .."))
}

func (user *UserController) GetAllUsers(w http.ResponseWriter,r *http.Request){
	user.UserService.GetAllUsers()
	w.Write([] byte("all users fetched .."))
}

func (user *UserController) DeleteUserByID(w http.ResponseWriter,r *http.Request){
	user.UserService.DeleteUserByID(1)
	w.Write([] byte("user deleted for the id 1 .."))
}

func NewUserController(_userService services.UserService) *UserController{
	return &UserController{
		UserService: _userService,
	}
} 