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
	GetUser(context.Context, *User) (*User, error)
}

type CreateUserReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserRes struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type LoginUserRes struct {
	accessToken string
	ID          string `json:"id"`
	Name        string `json:"name"`
}

type Service interface {
	CreateUser(context.Context, *CreateUserReq) (*CreateUserRes, error)
	LoginUser(context.Context, *LoginUserReq) (*LoginUserRes, error)
}
