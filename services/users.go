package services

import (
	DB "AeromindGO/DB/repositories"
	"fmt"
	"log"
)

type UserService interface {
	Create()
	GetUserByID(id int)
}

type UserServiceImpl struct {
	UserRepository DB.UserRepository
}

func (user *UserServiceImpl) Create() {
	fmt.Println("Creating user from user service ")
	user.UserRepository.Create()
}

func (user *UserServiceImpl) GetUserByID(id int) {
	fmt.Println("Fetching user details based on id")
	response,err:=user.UserRepository.GetUserByID(id)
	if err != nil{
		log.Fatal("Error while fetching the user by id : ",err)
	}
	fmt.Println("User details fetched successfully : ",response)
}


func NewUserService(_userRepository DB.UserRepository) UserService{
	return &UserServiceImpl{
		UserRepository: _userRepository,
	}
}