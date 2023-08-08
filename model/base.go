package model

import "gorm.io/gorm"

type BaseModel struct {
	ID int `gorm:"type:int;primary_key"`
}

type AdminBaseModel struct {
	// Make your own implementations here!
	gorm.Model
}
