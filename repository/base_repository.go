package repository

//type UserRepository interface {
//	Save(users model.User)
//	Update(users model.User)
//	Delete(usersId int)
//	FindById(usersId int) (users model.User, err error)
//	FindAll() []model.User
//}

type Repository interface {
	Save(data interface{})
	Update(data interface{})
	Delete(id int)
	FindById(id int) (interface{}, error)
	FindAll() []interface{}
}
