package step

import (
	"context"
	"github.com/laurentino14/user/graph/model"
	"github.com/laurentino14/user/prisma"
)

func GetAllSteps(ctx context.Context) []*model.Step {
	client := prisma.NewClient()

	if err := client.Prisma.Connect(); err != nil {
		return nil
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	exec, err := client.Step.FindMany().Take(10).Exec(ctx)

	if err != nil {
		return nil
	}

	allSteps := []*model.Step{}

	for _, list := range exec {

		user := &model.Step{
			ID:          list.ID,
			Title:       list.Title,
			Slug:        list.Slug,
			Description: list.Description,
			CreatedAt:   list.CreatedAt.String(),
			UpdatedAt:   list.UpdatedAt.String(),
			Lessons:     nil,
			CourseID:    list.CourseID,
		}
		allSteps = append(allSteps, user)
	}

	return allSteps
}
