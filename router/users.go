package router

import (
	"AeromindGO/controllers"

	"github.com/go-chi/chi/v5"
)

type UserRouter struct {
	UserController controllers.UserController
}

func (user *UserRouter) Register(r chi.Router){
	r.Post("/create",user.UserController.Create)
}

func NewRouter(_userController *controllers.UserController) Router{
	return  &UserRouter{
		UserController: *_userController,
	}
}