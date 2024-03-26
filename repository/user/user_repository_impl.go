package repository

import (
	"errors"
	"jin-example/data/request"
	"jin-example/helper"
	"jin-example/model"
)

// Save implements TagsRepository
func (u *UserRepositoryImpl) Save(user model.User) {
	result := u.Db.Create(&user)
	helper.ErrorPanic(result.Error)
}

// Delete implements TagsRepository
func (u *UserRepositoryImpl) Delete(userId int) {
	var user model.User
	result := u.Db.Where("id = ?", userId).Delete(&user)
	helper.ErrorPanic(result.Error)
}

// Update implements TagsRepository
func (u *UserRepositoryImpl) Update(user model.User) {
	var updateUser = request.UpdateUserRequest{
		Id:       user.Id,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
	}
	result := u.Db.Model(&user).Updates(updateUser)
	helper.ErrorPanic(result.Error)
}

// FindAll implements TagsRepository
func (u *UserRepositoryImpl) FindAll() []model.User {
	var users []model.User
	result := u.Db.Find(&users)
	helper.ErrorPanic(result.Error)
	return users
}

// FindById implements TagsRepository
func (u *UserRepositoryImpl) FindById(userId int) (users model.User, err error) {
	var user model.User
	result := u.Db.Find(&user, userId)
	if result != nil {
		return user, nil
	} else {
		return user, errors.New("tag is not found")
	}
}
