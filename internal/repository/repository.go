package repository

import (
	"chart/internal/models"
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"sync"
)

type Repository interface {
	CreateUser(context.Context, *models.User) (*models.User, error)
	GetUser(context.Context, string) (*models.User, error)
}

type repository struct {
	db *sql.DB
	mu sync.Mutex
}

func New(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := uuid.New().String()

	query := "INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4)"
	_, err := r.db.Exec(query, id, user.Name, user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	user.ID = id

	return user, nil
}

func (r *repository) GetUser(ctx context.Context, email string) (*models.User, error) {
	var u models.User

	query := "SELECT id, name, email, password FROM users WHERE email=$1"
	if err := r.db.QueryRow(query, email).Scan(&u.ID, &u.Name, &u.Email, &u.Password); err != nil {
		return nil, errors.New("email not found")
	}

	return &u, nil
}
