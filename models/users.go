package models

import (
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Password  string    ` json:"password"`
	IsAdmin   bool      `json:"isAdmin"`
	CreatedAt time.Time ` json:"createdAt"`
	UpdatedAt time.Time ` json:"updatedAt"`
}
