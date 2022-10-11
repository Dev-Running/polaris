package services

import (
	"context"

	"github.com/laurentino14/user/graph/model"
	"github.com/laurentino14/user/repositories"
)

type IAuthService interface {
	Auth(input *model.AuthenticationInput, ctx context.Context) (*model.User, error)
	GetUserAuthenticated(input *model.GetUserAuthInput, ctx context.Context) (*model.UserAuthenticated, error)
}

type AuthService struct {
	AuthRepository *repositories.AuthRepository
}

func (r *AuthService) GetUserAuthenticated(input *model.GetUserAuthInput, ctx context.Context) (*model.UserAuthenticated, error) {
	return r.AuthRepository.GetUserAuthenticated(input, ctx)
}

func NewAuthService(authRepository *repositories.AuthRepository) *AuthService {
	return &AuthService{AuthRepository: authRepository}
}

func (r *AuthService) Auth(input *model.AuthenticationInput, ctx context.Context) (*model.User, error) {
	return r.AuthRepository.Auth(input, ctx)
}
