package model

import "time"

type User struct {
	Name        string    `gorm:"type:varchar(255)"`
	Email       string    `gorm:"type:varchar(255)" json:"email"`
	DateOfBirth time.Time `gorm:"type:DATE"`
}
