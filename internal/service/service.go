package service

import (
	"chart/internal/models"
	"chart/util"
	"context"
	"errors"
)

type service struct {
	models.Repository
}

func New(r models.Repository) models.Service {
	return &service{r}
}

func (s *service) CreateUser(ctx context.Context, userReq *models.CreateUserReq) (*models.CreateUserRes, error) {
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
	u := models.User{
		Email:    loginUser.Email,
		Password: loginUser.Password,
	}

	user, err := s.Repository.GetUser(ctx, &u)
	if err != nil {
		return nil, err
	}

	if util.CheckPsw(user.Password, loginUser.Password) != nil {
		return nil, errors.New("Wrong data")
	}

	return &models.LoginUserRes{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}
