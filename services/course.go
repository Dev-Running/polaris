package services

import (
	"context"

	"github.com/laurentino14/user/graph/model"
	"github.com/laurentino14/user/repositories"
)

type ICourseService interface {
	Create(input model.NewCourse, ctx context.Context) (*model.Course, error)
	GetAll(ctx context.Context) ([]*model.Course, error)
}

type CourseService struct {
	CourseRepository *repositories.CourseRepository
}

func NewCourseService(courseRepository *repositories.CourseRepository) *CourseService {
	return &CourseService{
		CourseRepository: courseRepository,
	}
}

func (c *CourseService) Create(input model.NewCourse, ctx context.Context) (*model.Course, error) {
	return c.CourseRepository.Create(input, ctx)
}

func (c *CourseService) GetAll(ctx context.Context) ([]*model.Course, error) {
	return c.CourseRepository.GetAll(ctx)
}
