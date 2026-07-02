package controllers

import (
	"AeromindGO/dto"
	"AeromindGO/services"
	JSON "AeromindGO/utils/json"
	utils "AeromindGO/utils/responseFormatters"
	validators "AeromindGO/utils/validators"
	"net/http"
	"strings"
)

type UserController struct {
	UserService services.UserService
}

func (user *UserController) Create(w http.ResponseWriter , r *http.Request){

	var payload dto.Login_And_SignUp_UserDTO

	if err := JSON.FromJSON(r,&payload); err != nil {
		utils.ErrorResponse(w,http.StatusBadRequest,"Error occured while reading json.",err)
		return
	}

	if validationErr := validators.Validate.Struct(payload); validationErr != nil {
		utils.ErrorResponse(w,http.StatusBadRequest,"Invalid request payload",validationErr)
		return
	}

	response,err:=user.UserService.Create(payload.Email,payload.Password)
	if err != nil{
		status := http.StatusInternalServerError
		if strings.Contains(strings.ToLower(err.Error()), "duplicate") || strings.Contains(err.Error(), "1062") {
			status = http.StatusConflict
		}
		utils.ErrorResponse(w,status,"Error occured while creating user.",err)
		return
	} 
	utils.SuccessResponse(w,http.StatusCreated,"User created successfully.",response)
}

func (user *UserController) GetUserByID(w http.ResponseWriter,r *http.Request){
	response,err:=user.UserService.GetUserByID(2)
	if err != nil{
		utils.ErrorResponse(w,http.StatusInternalServerError,"Error occured while fetching user.",err)
		return
	}
	utils.SuccessResponse(w,http.StatusOK,"User fetched successfully.",response)
}

func (user *UserController) GetAllUsers(w http.ResponseWriter,r *http.Request){
	response,err:=user.UserService.GetAllUsers()
	if err != nil{
		utils.ErrorResponse(w,http.StatusInternalServerError,"Error occured while fetching users.",err)
		return
	}
	utils.SuccessResponse(w,http.StatusOK,"Users fetched successfully.",response)
}

func (user *UserController) DeleteUserByID(w http.ResponseWriter,r *http.Request){
	response,err:=user.UserService.DeleteUserByID(2)
	if err != nil{
		utils.ErrorResponse(w,http.StatusInternalServerError,"Error occured while deleting user.",err)
		return
	}
	utils.SuccessResponse(w,http.StatusOK,"User deleted successfully.",response)
}

func (user *UserController) Login(w http.ResponseWriter,r *http.Request){
	response,err:=user.UserService.Login("mahi@xx.com","thisismeAbhinav")
	if err != nil{
		status := http.StatusInternalServerError
		if strings.Contains(strings.ToLower(err.Error()), "invalid credentials") {
			status = http.StatusUnauthorized
		}
		utils.ErrorResponse(w,status,"Error occured while logging in.",err)
		return
	}
	utils.SuccessResponse(w,http.StatusOK,"Login successful.",response)
}

func NewUserController(_userService services.UserService) *UserController{
	return &UserController{
		UserService: _userService,
	}
} 