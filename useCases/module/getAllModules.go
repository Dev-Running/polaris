package module

import (
	"context"
	"github.com/laurentino14/user/graph/model"
	"github.com/laurentino14/user/prisma"
)

func GetAllModules(ctx context.Context) []*model.Module {
	client := prisma.NewClient()

	if err := client.Prisma.Connect(); err != nil {
		return nil
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	exec, err := client.Module.FindMany().Take(10).Exec(ctx)

	if err != nil {
		return nil
	}

	allModules := []*model.Module{}

	for _, list := range exec {

		user := &model.Module{
			ID:          list.ID,
			Title:       list.Title,
			Slug:        list.Slug,
			Description: list.Description,
			CreatedAt:   list.CreatedAt.String(),
			UpdatedAt:   list.UpdatedAt.String(),
			Lessons:     nil,
			Course:      nil,
		}
		allModules = append(allModules, user)
	}

	return allModules
}
