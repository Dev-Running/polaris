package lesson

import (
	"context"
	"github.com/laurentino14/user/graph/model"
	"github.com/laurentino14/user/prisma"
)

func GetAllLessons(ctx context.Context) []*model.Lesson {
	client := prisma.NewClient()

	if err := client.Prisma.Connect(); err != nil {
		return nil
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	exec, err := client.Lesson.FindMany().Take(10).Exec(ctx)

	if err != nil {
		return nil
	}

	allLessons := []*model.Lesson{}

	for _, list := range exec {

		createdAt := list.CreatedAt.String()
		updatedAt := list.CreatedAt.String()

		lesson := &model.Lesson{
			ID:        &list.ID,
			Title:     list.Title,
			Slug:      list.Slug,
			Link:      list.Link,
			CreatedAt: &createdAt,
			UpdatedAt: &updatedAt,
			StepID:    list.StepID,
		}
		allLessons = append(allLessons, lesson)
	}

	return allLessons
}
