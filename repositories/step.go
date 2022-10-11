package repositories

import (
	"context"
	"time"

	"github.com/laurentino14/user/graph/model"
	"github.com/laurentino14/user/prisma"
	"github.com/laurentino14/user/prisma/connect"
	"github.com/laurentino14/user/repositories/utils"
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
	).With(prisma.Step.Lessons.Fetch()).Exec(ctx)
	if err != nil {
		return nil, err
	}

	lessons := []*model.Lesson{}

	for _, l := range exec.RelationsStep.Lessons {
		lessons = append(lessons, &model.Lesson{
			ID:          l.ID,
			Title:       l.Title,
			Slug:        l.Slug,
			Description: l.Description,
			StepID:      l.StepID,
			CreatedAt:   l.CreatedAt.String(),
			UpdatedAt:   utils.ExtractData(l.UpdatedAt),
		})
	}

	stepData := &model.Step{
		ID:          exec.ID,
		Title:       exec.Title,
		Description: exec.Description,
		Slug:        exec.Slug,
		Lessons:     lessons,
		CreatedAt:   exec.CreatedAt.String(),
		UpdatedAt:   utils.ExtractData(exec.UpdatedAt),
		CourseID:    exec.CourseID,
	}
	return stepData, nil
}

func (r *StepRepository) GetAll(ctx context.Context) ([]*model.Step, error) {
	exec, err := r.DB.Client.Step.FindMany().With(prisma.Step.Lessons.Fetch()).Take(10).Exec(ctx)

	if err != nil {
		return nil, err
	}

	allSteps := []*model.Step{}
	lessons := []*model.Lesson{}
	for _, list := range exec {

		for _, l := range list.RelationsStep.Lessons {
			lessons = append(lessons, &model.Lesson{
				ID:          l.ID,
				Title:       l.Title,
				Slug:        l.Slug,
				Description: l.Description,
				StepID:      l.StepID,
				CreatedAt:   l.CreatedAt.String(),
				UpdatedAt:   utils.ExtractData(l.UpdatedAt),
			})
		}

		user := &model.Step{
			ID:          list.ID,
			Title:       list.Title,
			Slug:        list.Slug,
			Description: list.Description,
			CreatedAt:   list.CreatedAt.String(),
			UpdatedAt:   utils.ExtractData(list.UpdatedAt),
			Lessons:     lessons,
			CourseID:    list.CourseID,
		}
		allSteps = append(allSteps, user)
	}

	return allSteps, nil
}
