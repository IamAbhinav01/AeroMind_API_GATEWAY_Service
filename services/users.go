package services

import (
	DB "AeromindGO/DB/repositories"
	"fmt"
)

type UserService interface {
	Create()
}

type UserServiceImpl struct {
	UserRepository DB.UserRepository
}

func (user *UserServiceImpl) Create() {
	fmt.Println("Creating user from user service ")
	user.UserRepository.Create()
}

func NewUserService(_userRepository DB.UserRepository) UserService{
	return &UserServiceImpl{
		UserRepository: _userRepository,
	}
}