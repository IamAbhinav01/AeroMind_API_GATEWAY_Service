package services

import (
	DB "AeromindGO/DB/repositories"
	"AeromindGO/models"
	"database/sql"
	"fmt"
	"log"
)

type UserService interface {
	Create() (int, error)
	GetUserByID(id int)(models.User, error)
	GetAllUsers() ([]models.User, error)
	DeleteUserByID(id int)(sql.Result, error)
}

type UserServiceImpl struct {
	UserRepository DB.UserRepository
}

func (user *UserServiceImpl) Create() (int, error) {
	fmt.Println("Creating user from user service ")
	response , err := user.UserRepository.Create()
	if err != nil{
		log.Fatal("Error while creating the user : ",err)
	}
	return response,nil;
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

func NewUserService(_userRepository DB.UserRepository) UserService{
	return &UserServiceImpl{
		UserRepository: _userRepository,
	}
}