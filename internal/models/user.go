package models

import "time"

type User struct {
	ID         uint      `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Role       string    `json:"role"`
	ProfilePic string    `json:"profilepic"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
