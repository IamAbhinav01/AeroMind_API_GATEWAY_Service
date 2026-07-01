package services

import (
	DB "AeromindGO/DB/repositories"
	"fmt"
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
	user.UserRepository.GetUserByID(id)
}


func NewUserService(_userRepository DB.UserRepository) UserService{
	return &UserServiceImpl{
		UserRepository: _userRepository,
	}
}