package services

import (
	"context"

	"github.com/laurentino14/user/graph/model"
	"github.com/laurentino14/user/repositories"
)

type IUserService interface {
	Avatar(input model.NewUser, imageType string) string
	Create(input model.NewUser, ctx context.Context) (*model.User, error)
	CreateUserGITHUB(input model.NewUserGithub, ctx context.Context) (*model.User, error)
	CreateUserGOOGLE(input model.NewUserGoogle, ctx context.Context) (*model.User, error)
	GetAll(ctx context.Context) ([]*model.User, error)
	GetUserByID(id string, ctx context.Context) (*model.User, error)
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

func (r *UserService) Avatar(input model.NewUser, imageType string) string {
	return r.UserRepository.Avatar(input, imageType)
}

func (r *UserService) CreateUserGITHUB(input model.NewUserGithub, ctx context.Context) (*model.User, error) {
	return r.UserRepository.CreateUserGITHUB(input, ctx)
}

func (r *UserService) CreateUserGOOGLE(input model.NewUserGoogle, ctx context.Context) (*model.User, error) {
	return r.UserRepository.CreateUserGOOGLE(input, ctx)
}

func (r *UserService) GetUserByID(id string, ctx context.Context) (*model.User, error) {
	return r.UserRepository.GetUserByID(id, ctx)
}
