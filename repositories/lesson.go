package repositories

import (
	"context"
	"github.com/laurentino14/user/graph/model"
	"github.com/laurentino14/user/prisma"
	"github.com/laurentino14/user/prisma/connect"
)

type ILessonRepository interface {
	Create(input *model.NewLesson, ctx context.Context) (*model.Lesson, error)
	GetAll() ([]*model.Lesson, error)
}

type LessonRepository struct {
	DB *connect.DB
}

func (c *LessonRepository) Create(input *model.NewLesson, ctx context.Context) (*model.Lesson, error) {
	exec, err := c.DB.Client.Lesson.CreateOne(
		prisma.Lesson.Title.Set(input.Title),
		prisma.Lesson.Slug.Set(input.Slug),
		prisma.Lesson.Link.Set(input.Link),
		prisma.Lesson.Step.Link(prisma.Step.ID.Equals(input.StepID)),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	createdAt := exec.CreatedAt.String()
	updatedAt := exec.UpdatedAt.String()

	lessonData := &model.Lesson{
		ID:        &exec.ID,
		Title:     exec.Title,
		Slug:      exec.Slug,
		Link:      exec.Link,
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
		StepID:    exec.StepID,
	}
	return lessonData, nil
}
