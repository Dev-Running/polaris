package course

import (
	"context"
	"github.com/laurentino14/user/graph/model"
	"github.com/laurentino14/user/prisma"
)

func GetAllCourses(ctx context.Context) []*model.Course {
	client := prisma.NewClient()

	if err := client.Prisma.Connect(); err != nil {
		return nil
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	exec, err := client.Course.FindMany().Take(10).Exec(ctx)

	if err != nil {
		return nil
	}

	allCourses := []*model.Course{}

	for _, list := range exec {

		user := &model.Course{
			ID:          list.ID,
			Title:       list.Title,
			Slug:        list.Slug,
			Description: &list.Description,
			CreatedAt:   list.CreatedAt.String(),
			UpdatedAt:   list.UpdatedAt.String(),
		}
		allCourses = append(allCourses, user)
	}

	return allCourses
}
