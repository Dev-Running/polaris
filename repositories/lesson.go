package repositories

import (
	"context"
	"github.com/laurentino14/user/graph/model"
	"github.com/laurentino14/user/prisma"
	"github.com/laurentino14/user/prisma/connect"
)

type ILessonRepository interface {
	Create(input model.NewLesson, ctx context.Context) (*model.Lesson, error)
	GetAll(ctx context.Context) ([]*model.Lesson, error)
}

type LessonRepository struct {
	DB *connect.DB
}

// NewLessonRepository implements LessonRepository
func NewLessonRepository(db *connect.DB) *LessonRepository {
	return &LessonRepository{
		DB: db,
	}
}

// Create implements LessonRepository
func (c *LessonRepository) Create(input model.NewLesson, ctx context.Context) (*model.Lesson, error) {
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

// GetAll implements LessonRepository
func (c *LessonRepository) GetAll(ctx context.Context) ([]*model.Lesson, error) {
	exec, err := c.DB.Client.Lesson.FindMany().Take(10).Exec(ctx)

	if err != nil {
		return nil, err
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

	return allLessons, nil
}
