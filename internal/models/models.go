package models

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
