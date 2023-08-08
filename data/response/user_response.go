package response

import (
	"strconv"
	"time"
)

type UserResponse struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"_"`
	Email       string    `json:"email"`
}

func (ur *UserResponse) MarshalJSON() ([]byte, error) {
	type Alias UserResponse
	return []byte(`{"id":` + strconv.Itoa(ur.Id) +
		`,"name":"` + ur.Name +
		`","date_of_birth":"` + ur.DateOfBirth.Format("2006-01-02") + // Format the date
		`","email":"` + ur.Email + `"}`), nil
}
