package user

import (
	"context"
	"github.com/laurentino14/user/graph/model"
	"github.com/laurentino14/user/prisma"
)

func GetAllUsers(ctx context.Context) []*model.User {
	client := prisma.NewClient()

	if err := client.Prisma.Connect(); err != nil {
		return nil
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	exec, err := client.User.FindMany().Take(10).Exec(ctx)

	if err != nil {
		return nil
	}

	allUsers := []*model.User{}

	for _, list := range exec {

		token := "*******"
		user := &model.User{
			ID:         list.ID,
			Firstname:  list.Firstname,
			Lastname:   list.Lastname,
			Email:      list.Email,
			Password:   list.Password,
			Cellphone:  list.Cellphone,
			TokenUser:  &token,
			Enrollment: nil,
		}
		allUsers = append(allUsers, user)
	}

	return allUsers
}
