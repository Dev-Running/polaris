package step

import (
	"context"
	"time"

	"github.com/laurentino14/user/graph/model"
	"github.com/laurentino14/user/prisma"
)

func CreateStep(input model.NewStep, ctx context.Context) (*model.Step, error) {
	client := prisma.NewClient()

	if err := client.Prisma.Connect(); err != nil {
		return nil, err
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	exec, err := client.Step.CreateOne(
		prisma.Step.Title.Set(input.Title),
		prisma.Step.Description.Set(input.Description),
		prisma.Step.Slug.Set(input.Slug),
		prisma.Step.Course.Link(prisma.Course.ID.Equals(input.CourseID)),
		prisma.Step.CreatedAt.Set(time.Now()),
		prisma.Step.UpdatedAt.Set(time.Now()),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	moduleData := &model.Step{
		ID:          exec.ID,
		Title:       exec.Title,
		Description: exec.Description,
		Slug:        exec.Slug,
		Lessons:     nil,
		CourseID:    exec.CourseID,
	}
	return moduleData, nil
}
