package service

import (
	"jin-example/data/request"
	"jin-example/data/response"
	repository "jin-example/repository/user"

	"github.com/go-playground/validator/v10"
)

type UserService interface {
	CreateUser(user request.CreateUserRequest)
	UpdateUser(user request.UpdateUserRequest)
	DeleteUser(user request.DeleteUserRequest)
	FindUserById(user request.FindUserRequest) response.UserResponse
	FindAllUser() []response.UserResponse
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewUserService(repository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: repository,
		Validate:       validate,
	}
}
