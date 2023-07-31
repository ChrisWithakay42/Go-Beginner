package repository

import "go_rest/model"

type UserRepository interface {
	Save(users model.User)
	Update(users model.User)
	Delete(usersId int)
	FindById(usersId int) (users model.User, err error)
	FindAll() []model.User
}
