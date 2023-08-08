package repository

import "gorm.io/gorm"

type UserRepository struct {
	baseRepo *BaseRepository
}

func NewUserRepository(Db *gorm.DB) UserRepository {
	return UserRepository{baseRepo: &BaseRepository{Db: Db}}
}
