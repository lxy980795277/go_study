package repository

import (
	"jin-example/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user model.User)
	Update(user model.User)
	Delete(userId int)
	FindById(Userid int) (tags model.User, err error)
	FindAll() []model.User
}

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}
