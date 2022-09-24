package repositories

import (
	"context"
	"github.com/laurentino14/user/graph/model"
	"github.com/laurentino14/user/prisma"
	"github.com/laurentino14/user/prisma/connect"
	"time"
)

type IStepRepository interface {
	Create(input model.NewStep, ctx context.Context) (*model.Step, error)
	GetAll(ctx context.Context) ([]*model.Step, error)
}

type StepRepository struct {
	DB *connect.DB
}

func NewStepRepository(db *connect.DB) *StepRepository {
	return &StepRepository{
		DB: db,
	}
}

func (r *StepRepository) Create(input model.NewStep, ctx context.Context) (*model.Step, error) {
	exec, err := r.DB.Client.Step.CreateOne(
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

	stepData := &model.Step{
		ID:          exec.ID,
		Title:       exec.Title,
		Description: exec.Description,
		Slug:        exec.Slug,
		Lessons:     nil,
		CourseID:    exec.CourseID,
	}
	return stepData, nil
}

func (r *StepRepository) GetAll(ctx context.Context) ([]*model.Step, error) {
	exec, err := r.DB.Client.Step.FindMany().Take(10).Exec(ctx)

	if err != nil {
		return nil, err
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

	return allSteps, nil
}
