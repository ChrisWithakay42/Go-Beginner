package model

import "gorm.io/gorm"

type BaseModel struct {
	ID int `gorm:"type:int;primary_key"`
}

type AdminBaseModel struct {
	// Makefile your own implementations here!
	// with admin related fields
	gorm.Model
}
