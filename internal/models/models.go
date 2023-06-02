package models

import "github.com/golang-jwt/jwt/v4"

type User struct {
	ID       string `json:"id" db:"id""`
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
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

type Claims struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}
