package controllers

import (
	"AeromindGO/dto"
	"AeromindGO/services"
	utils "AeromindGO/utils/responseFormatters"
	validators "AeromindGO/utils/validators"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

type UserController struct {
	UserService services.UserService
}

func (user *UserController) Create(w http.ResponseWriter , r *http.Request){


	payload := r.Context().Value("payload").(dto.CreateUserDTO)

	if validationErr := validators.Validate.Struct(payload); validationErr != nil {
		utils.ErrorResponse(w,http.StatusBadRequest,"Invalid request payload",validationErr)
		return
	}

	response,err:=user.UserService.Create(payload)
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

	idStr:= chi.URLParam(r,"id")
	id,err:=strconv.Atoi(idStr)

	if err != nil{
		log.Fatal("Error happenend in controller layer: ",err)
	}
	response,err:=user.UserService.GetUserByID(id)
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
	payload := r.Context().Value("payload").(dto.LoginUserDTO)

	if validationErr := validators.Validate.Struct(payload);validationErr != nil{
		utils.ErrorResponse(w,http.StatusBadRequest,"Invalid request payload",validationErr)
		return
	}

	response,err:=user.UserService.Login(payload)
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