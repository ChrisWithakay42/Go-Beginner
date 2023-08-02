package service

import (
	"github.com/go-playground/validator/v10"
	"go_rest/data/request"
	"go_rest/data/response"
	"go_rest/model"
	"go_rest/repository"
	"go_rest/utils"
	"time"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewUserServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

func (u *UserServiceImpl) Create(user request.CreateUserRequest) {
	err := u.Validate.Struct(user)
	utils.ErrorPanic(err)

	dob, err := time.Parse("2006-01-02", user.DateOfBirth)

	userModel := model.User{
		Name:        user.Name,
		DateOfBirth: dob,
		Email:       user.Email,
	}
	u.UserRepository.Save(userModel)
}

func (u *UserServiceImpl) Delete(userId int) {
	u.UserRepository.Delete(userId)
}

func (u *UserServiceImpl) FindAll() []response.UserResponse {
	result := u.UserRepository.FindAll()

	var users []response.UserResponse
	for _, value := range result {
		user := response.UserResponse{
			Id:          value.ID,
			Name:        value.Name,
			DateOfBirth: value.DateOfBirth,
			Email:       value.Email,
		}
		users = append(users, user)
	}
	return users
}

func (u *UserServiceImpl) FindById(userId int) response.UserResponse {
	userData, err := u.UserRepository.FindById(userId)
	utils.ErrorPanic(err)

	userResponse := response.UserResponse{
		Id:          userData.ID,
		Name:        userData.Name,
		DateOfBirth: userData.DateOfBirth,
		Email:       userData.Email,
	}
	return userResponse
}

func (u *UserServiceImpl) Update(user request.UpdateUserRequest) {
	userData, err := u.UserRepository.FindById(user.Id)
	utils.ErrorPanic(err)
	userData.Name = user.Name
	userData.DateOfBirth = user.DateOfBirth
	userData.Email = user.Email
	u.UserRepository.Update(userData)
}
