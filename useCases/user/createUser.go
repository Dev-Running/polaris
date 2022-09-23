package user

import (
	"context"
	"github.com/laurentino14/user/graph/model"
	"github.com/laurentino14/user/prisma"
)

func CreateUser(input model.NewUser, ctx context.Context) (*model.User, error) {
	client := prisma.NewClient()

	if err := client.Prisma.Connect(); err != nil {
		return nil, err
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
	exec, err := client.User.CreateOne(
		prisma.User.Firstname.Set(input.Firstname),
		prisma.User.Lastname.Set(input.Lastname),
		prisma.User.Email.Set(input.Email),
		prisma.User.Password.Set(input.Password),
		prisma.User.Cellphone.Set(input.Cellphone),
		prisma.User.TokenUser.Set(*input.TokenUser),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	userData := &model.User{
		ID:        exec.ID,
		Firstname: input.Firstname,
		Lastname:  input.Lastname,
		Email:     input.Email,
		Password:  input.Password,
		Cellphone: input.Cellphone,
		TokenUser: input.TokenUser,
	}
	return userData, nil
}
