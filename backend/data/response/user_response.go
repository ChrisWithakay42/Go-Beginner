package response

import "time"

type UserResponse struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Email       string    `json:"email"`
}
