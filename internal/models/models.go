package models

import "context"

type User struct {
	ID       string `json:"id" db:"id""`
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
type Repository interface {
	CreateUser(context.Context, *User) (*User, error)
}
