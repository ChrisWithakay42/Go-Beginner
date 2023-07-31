package request

import "time"

type UpdateUserRequest struct {
	Id          int       `validate:"required"`
	Name        string    `validate:"required,min=1,max=200" json:"name"`
	DateOfBirth time.Time `validate:"required" json:"dob"`
	Email       string    `validate:"required,email" json:"email"`
}
