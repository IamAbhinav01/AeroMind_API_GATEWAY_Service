package controllers

import (
	"AeromindGO/services"
	utils "AeromindGO/utils/responseFormatters"
	"net/http"
)

type UserController struct {
	UserService services.UserService
}

func (user *UserController) Create(w http.ResponseWriter , r *http.Request){
	response,err:=user.UserService.Create()
	if err != nil{
		utils.ErrorResponse(w,http.StatusAccepted,"Error occured while creating user.",err)
	} else {
		utils.SuccessResponse(w,http.StatusCreated,"User created successfully.",response)
	}
}

func (user *UserController) GetUserByID(w http.ResponseWriter,r *http.Request){
	response,err:=user.UserService.GetUserByID(2)
	if err != nil{
		utils.ErrorResponse(w,http.StatusAccepted,"Error occured while fetching user.",err)
	} else {
		utils.SuccessResponse(w,http.StatusCreated,"User fetched successfully.",response)
	}
}

func (user *UserController) GetAllUsers(w http.ResponseWriter,r *http.Request){
	response,err:=user.UserService.GetAllUsers()
	if err != nil{
		utils.ErrorResponse(w,http.StatusAccepted,"Error occured while fetching users.",err)
	} else {
		utils.SuccessResponse(w,http.StatusCreated,"Users fetched successfully.",response)
	}
}

func (user *UserController) DeleteUserByID(w http.ResponseWriter,r *http.Request){
	response,err:=user.UserService.DeleteUserByID(2)
	if err != nil{
		utils.ErrorResponse(w,http.StatusAccepted,"Error occured while deleting user.",err)
	} else {
		utils.SuccessResponse(w,http.StatusCreated,"User deleted successfully.",response)
	}
}

func NewUserController(_userService services.UserService) *UserController{
	return &UserController{
		UserService: _userService,
	}
} 