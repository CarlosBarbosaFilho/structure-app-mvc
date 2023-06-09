package response

import "time"

type UserResponse struct {
	ID       uint      `json:"id"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
}
