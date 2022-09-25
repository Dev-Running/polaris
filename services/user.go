package services

import (
	"context"
	"github.com/laurentino14/user/graph/model"
	"github.com/laurentino14/user/repositories"
)

type IUserService interface {
	Create(input model.NewUser, ctx context.Context) (*model.User, error)
	GetAll(ctx context.Context) ([]*model.User, error)
}

type UserService struct {
	UserRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (r *UserService) Create(input model.NewUser, ctx context.Context) (*model.User, error) {
	return r.UserRepository.Create(input, ctx)
}

func (r *UserService) GetAll(ctx context.Context) ([]*model.User, error) {
	return r.UserRepository.GetAll(ctx)
}