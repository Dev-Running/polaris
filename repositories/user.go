package repositories

import (
	"context"
	"github.com/laurentino14/user/graph/model"
	"github.com/laurentino14/user/prisma"
	"github.com/laurentino14/user/prisma/connect"
)

type IUserRepository interface {
	Create(input model.NewUser, ctx context.Context) (*model.User, error)
	GetAll(ctx context.Context) ([]*model.User, error)
}

type UserRepository struct {
	DB *connect.DB
}

func NewUserRepository(db *connect.DB) *UserRepository {
	return &UserRepository{DB: db}
}
func (r *UserRepository) Create(input model.NewUser, ctx context.Context) (*model.User, error) {
	exec, err := r.DB.Client.User.CreateOne(
		prisma.User.Firstname.Set(input.Firstname),
		prisma.User.Lastname.Set(input.Lastname),
		prisma.User.Email.Set(input.Email),
		prisma.User.Password.Set(input.Password),
		prisma.User.Cellphone.Set(input.Cellphone),
		prisma.User.TokenUser.Set(""),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}
	a := ""
	userData := &model.User{
		ID:         exec.ID,
		Firstname:  input.Firstname,
		Lastname:   input.Lastname,
		Email:      input.Email,
		Password:   input.Password,
		Cellphone:  input.Cellphone,
		TokenUser:  &a,
		Enrollment: nil,
	}
	return userData, nil
}

func (r *UserRepository) GetAll(ctx context.Context) ([]*model.User, error) {
	exec, err := r.DB.Client.User.FindMany().Take(10).Exec(ctx)

	if err != nil {
		return nil, err
	}

	allUsers := []*model.User{}
	for _, list := range exec {
		user := &model.User{
			ID:         list.ID,
			Firstname:  list.Firstname,
			Lastname:   list.Lastname,
			Email:      list.Email,
			Password:   list.Password,
			Cellphone:  list.Cellphone,
			TokenUser:  &list.TokenUser,
			Enrollment: nil,
		}
		allUsers = append(allUsers, user)
	}

	return allUsers, nil
}
