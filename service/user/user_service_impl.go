package service

import (
	"fmt"
	"jin-example/data/request"
	"jin-example/data/response"
	"jin-example/helper"
	"jin-example/model"
)

func (u *UserServiceImpl) CreateUser(user request.CreateUserRequest) {
	err := u.Validate.Struct(user)
	helper.ErrorPanic(err)

	userModel := model.User{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
	}
	u.UserRepository.Save(userModel)
}

func (u *UserServiceImpl) DeleteUser(user request.DeleteUserRequest) {
	u.UserRepository.Delete(user.Id)
}

func (u *UserServiceImpl) UpdateUser(user request.UpdateUserRequest) {
	userData, err := u.UserRepository.FindById(user.Id)
	helper.ErrorPanic(err)
	userData.Password = user.Password
	userData.Email = user.Email
	u.UserRepository.Update(userData)
}

func (u *UserServiceImpl) FindUserById(user request.FindUserRequest) response.UserResponse {
	userData, err := u.UserRepository.FindById(user.Id)
	helper.ErrorPanic(err)
	fmt.Println(userData)
	userResponse := response.UserResponse{
		Id:       userData.Id,
		Username: userData.Username,
		Password: userData.Password,
		Email:    userData.Email,
	}
	return userResponse
}

func (u *UserServiceImpl) FindAllUser() []response.UserResponse {
	result := u.UserRepository.FindAll()

	var users []response.UserResponse

	for _, userData := range result {
		tag := response.UserResponse{
			Id:       userData.Id,
			Username: userData.Username,
			Password: userData.Password,
			Email:    userData.Email,
		}
		users = append(users, tag)
	}

	return users
}
