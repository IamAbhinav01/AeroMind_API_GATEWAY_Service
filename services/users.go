package services

import (
	DB "AeromindGO/DB/repositories"
	"AeromindGO/auth"
	"AeromindGO/models"
	"database/sql"
	"fmt"
	"log"
)

type UserService interface {
	Create(email string, password string) (int, error)
	GetUserByID(id int)(models.User, error)
	GetAllUsers() ([]models.User, error)
	DeleteUserByID(id int)(sql.Result, error)
	Login(email string, password string) (string, error)
}

type UserServiceImpl struct {
	UserRepository DB.UserRepository
}

func (user *UserServiceImpl) Create(email string, password string) (int, error) {
	fmt.Println("Creating user from user service ")
	HashedPassword,err := auth.HashPassword(password)
	if err != nil{
		log.Fatal("Error while hashing the password : ",err)
	}
	response , err := user.UserRepository.Create(email,HashedPassword)
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

func (user *UserServiceImpl) Login(email string, password string) (string, error) {

	userModel, err := user.UserRepository.GetUserByEmail(email)
	if err != nil {
		return "", fmt.Errorf("login: %w", err)
	}

	if (auth.CheckPasswordHash(password, userModel.Password) == true){
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