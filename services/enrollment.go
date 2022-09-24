package services

import (
	"context"
	"github.com/laurentino14/user/graph/model"
	"github.com/laurentino14/user/repositories"
)

type IEnrollmentService interface {
	Create(input model.NewEnrollment, ctx context.Context) (*model.Enrollment, error)
	GetAll(ctx context.Context) ([]*model.Enrollment, error)
}

type EnrollmentService struct {
	EnrollmentRepository *repositories.EnrollmentRepository
}

func NewEnrollmentService(enrollmentRepository *repositories.EnrollmentRepository) *EnrollmentService {
	return &EnrollmentService{
		EnrollmentRepository: enrollmentRepository,
	}
}

func (r *EnrollmentService) Create(input model.NewEnrollment, ctx context.Context) (*model.Enrollment, error) {
	return r.EnrollmentRepository.Create(input, ctx)
}

func (r *EnrollmentService) GetAll(ctx context.Context) ([]*model.Enrollment, error) {
	return r.EnrollmentRepository.GetAll(ctx)
}
