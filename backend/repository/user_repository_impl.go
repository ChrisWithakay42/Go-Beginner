package repository

import (
	"errors"
	"go_rest/data/request"
	"go_rest/model"
	"go_rest/utils"
	"gorm.io/gorm"
)

type UserRepoImpl struct {
	Db *gorm.DB
}

func NewUserRepoImpl(Db *gorm.DB) UserRepository {
	return &UserRepoImpl{Db: Db}
}

func (u *UserRepoImpl) FindAll() []model.User {
	var users []model.User
	result := u.Db.Find(&users)
	utils.ErrorPanic(result.Error)
	return users
}

func (u *UserRepoImpl) FindById(userId int) (users model.User, err error) {
	var user model.User
	result := u.Db.Find(&user, userId)
	if result != nil {
		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

func (u *UserRepoImpl) Save(user model.User) {
	result := u.Db.Create(&user)
	utils.ErrorPanic(result.Error)
}

func (u *UserRepoImpl) Update(user model.User) {
	var updateUser = request.UpdateUserRequest{
		Id:          user.ID,
		Name:        user.Name,
		DateOfBirth: user.DateOfBirth,
		Email:       user.Email,
	}
	result := u.Db.Model(&user).Updates(updateUser)
	utils.ErrorPanic(result.Error)
}

func (u *UserRepoImpl) Delete(userId int) {
	var users model.User
	result := u.Db.Where("id = ?", userId).Delete(&users)
	utils.ErrorPanic(result.Error)
}
