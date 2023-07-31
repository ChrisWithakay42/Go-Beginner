package model

import "gorm.io/gorm"

type BaseModel struct {
	gorm.Model
	ID uint `gorm:"type:int;primary_key"`
}

type AdminBaseModel struct {
	// Make your own implementations here!
	gorm.Model
}
