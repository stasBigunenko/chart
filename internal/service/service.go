package service

import (
	"chart/internal/models"
	"chart/internal/repository"
	"chart/util"
	"context"
	"errors"
)

type Service interface {
	CreateUser(context.Context, *models.CreateUserReq) (*models.CreateUserRes, error)
	LoginUser(context.Context, *models.LoginUserReq) (*models.LoginUserRes, error)
	userExist(context.Context, string) bool
}

type service struct {
	repository.Repository
}

func New(r repository.Repository) Service {
	return &service{r}
}

func (s *service) CreateUser(ctx context.Context, userReq *models.CreateUserReq) (*models.CreateUserRes, error) {
	if s.userExist(ctx, userReq.Email) {
		return nil, errors.New("email already exists")
	}

	hashPassword, err := util.HashPsw(userReq.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:     userReq.Name,
		Email:    userReq.Email,
		Password: hashPassword,
	}

	u, err := s.Repository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &models.CreateUserRes{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}, nil
}

func (s *service) LoginUser(ctx context.Context, loginUser *models.LoginUserReq) (*models.LoginUserRes, error) {
	user, err := s.Repository.GetUser(ctx, loginUser.Email)
	if err != nil {
		return nil, errors.New("Wrong data")
	}

	if util.CheckPsw(user.Password, loginUser.Password) != nil {
		return nil, errors.New("Wrong data")
	}

	return &models.LoginUserRes{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}

func (s *service) userExist(ctx context.Context, email string) bool {
	u, _ := s.Repository.GetUser(ctx, email)
	if u != nil {
		return true
	}

	return false

}
