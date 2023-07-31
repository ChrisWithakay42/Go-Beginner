package response

import "time"

type UserResponse struct {
	Id    int       `json:"id"`
	Name  string    `json:"name"`
	Dob   time.Time `json:"dob"`
	Email string    `json:"email"`
}
