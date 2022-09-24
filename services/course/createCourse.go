package course

import (
	"context"
	"github.com/laurentino14/user/graph/model"
	"github.com/laurentino14/user/prisma"
	"time"
)

func CreateCourse(input model.NewCourse, ctx context.Context) (*model.Course, error) {
	client := prisma.NewClient()

	if err := client.Prisma.Connect(); err != nil {
		return nil, err
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
	exec, err := client.Course.CreateOne(
		prisma.Course.Title.Set(input.Title),
		prisma.Course.Slug.Set(input.Slug),
		prisma.Course.Description.Set(*input.Description),
		prisma.Course.CreatedAt.Set(time.Now()),
		prisma.Course.UpdatedAt.Set(time.Now()),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	courseData := &model.Course{
		ID:          exec.ID,
		Title:       exec.Title,
		Slug:        exec.Slug,
		Description: &exec.Description,
		CreatedAt:   exec.CreatedAt.String(),
		UpdatedAt:   exec.UpdatedAt.String(),
		Lessons:     nil,
		Steps:       nil,
		Enrollments: nil,
	}
	return courseData, nil
}
