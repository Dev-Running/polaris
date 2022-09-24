package services

import (
	"context"
	"github.com/laurentino14/user/graph/model"
	"github.com/laurentino14/user/repositories"
)

type IStepService interface {
	Create(input model.NewStep, ctx context.Context) (*model.Step, error)
	GetAll(ctx context.Context) ([]*model.Step, error)
}
type StepService struct {
	StepRepository *repositories.StepRepository
}

func NewStepService(stepRepository *repositories.StepRepository) *StepService {
	return &StepService{
		StepRepository: stepRepository,
	}
}

func (r *StepService) Create(input model.NewStep, ctx context.Context) (*model.Step, error) {
	return r.StepRepository.Create(input, ctx)
}

func (r *StepService) GetAll(ctx context.Context) ([]*model.Step, error) {
	return r.StepRepository.GetAll(ctx)
}
