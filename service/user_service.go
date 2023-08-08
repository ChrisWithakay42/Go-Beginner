package service

import (
	"go_rest/data/request"
	"go_rest/data/response"
)

type UserService interface {
	Create(user request.CreateUserRequest)
	Update(user request.UpdateUserRequest)
	Delete(userId int)
	FindById(userId int) response.UserResponse
	FindAll() []response.UserResponse
}
