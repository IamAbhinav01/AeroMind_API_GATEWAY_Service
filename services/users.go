package services

import (
	DB "AeromindGO/DB/repositories"
	"AeromindGO/auth"
	"AeromindGO/dto"
	"AeromindGO/models"
	"database/sql"
	"fmt"
	"log"
)

type UserService interface {
	Create(dto.CreateUserDTO) (int, error)
	GetUserByID(id int)(models.User, error)
	GetAllUsers() ([]models.User, error)
	DeleteUserByID(id int)(sql.Result, error)
	Login(dto.LoginUserDTO)(string, error)
}

type UserServiceImpl struct {
	UserRepository DB.UserRepository
}

func (user *UserServiceImpl) Create(payload dto.CreateUserDTO) (int, error) {
	fmt.Println("Creating user from user service ")
	HashedPassword,err := auth.HashPassword(payload.Password)
	if err != nil{
		log.Fatal("Error while hashing the password : ",err)
	}
	response , err := user.UserRepository.Create(payload.Email,HashedPassword)
	if err != nil{
		log.Printf("Error while creating the user: %v", err)
		return 0, fmt.Errorf("create user service: %w", err)
	}
	return response,nil
}


func (user *UserServiceImpl) GetUserByID(id int) (models.User, error) {
	fmt.Println("Fetching user details based on id")
	response,err:=user.UserRepository.GetUserByID(id)
	if err != nil{
		log.Fatal("Error while fetching the user by id : ",err)
	}
	fmt.Println("User details fetched successfully ")
	return response,nil
}

func (user *UserServiceImpl) GetAllUsers()([]models.User, error){
	fmt.Println("Fetching all users")
	response,err:=user.UserRepository.GetAllUsers()
	if err != nil{
		log.Fatal("Error while fetching all users : ",err)
	}
	fmt.Println("All users fetched successfully  ")
	return response,nil
}

func (user *UserServiceImpl) DeleteUserByID(id int)(sql.Result, error){

	fmt.Println("Deleting user based on id")
	response,err := user.UserRepository.DeleteUserByID(id)
	if err != nil{
		log.Fatal("Error while deleting the user : ",err)
	}
	fmt.Println("User deleted successfully ")
	return response,nil
}

func (user *UserServiceImpl) Login(payload dto.LoginUserDTO) (string, error) {

	userModel, err := user.UserRepository.GetUserByEmail(payload.Email)
	if err != nil {
		return "", fmt.Errorf("login: %w", err)
	}

	if (auth.CheckPasswordHash(payload.Password, userModel.Password) == true){
		return "Login successful", nil
	}else{
		return "", fmt.Errorf("login: invalid credentials")
	}
}

func NewUserService(_userRepository DB.UserRepository) UserService{
	return &UserServiceImpl{
		UserRepository: _userRepository,
	}
}