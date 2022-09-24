package repositories

import (
	"context"
	"github.com/laurentino14/user/graph/model"
	"github.com/laurentino14/user/prisma"
	"github.com/laurentino14/user/prisma/connect"
	"time"
)

type ICourseRepository interface {
	Create(input model.NewCourse, ctx context.Context) (*model.Course, error)
	GetAl(ctx context.Context) ([]*model.Course, error)
}

type CourseRepository struct {
	DB *connect.DB
}

func NewCourseRepository(db *connect.DB) *CourseRepository {
	return &CourseRepository{DB: db}
}

func (r *CourseRepository) Create(input model.NewCourse, ctx context.Context) (*model.Course, error) {
	exec, err := r.DB.Client.Course.CreateOne(
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

func (r *CourseRepository) GetAll(ctx context.Context) ([]*model.Course, error) {
	exec, err := r.DB.Client.Course.FindMany().Take(10).Exec(ctx)

	if err != nil {
		return nil, err
	}

	allCourses := []*model.Course{}

	for _, list := range exec {

		user := &model.Course{
			ID:          list.ID,
			Title:       list.Title,
			Slug:        list.Slug,
			Description: &list.Description,
			CreatedAt:   list.CreatedAt.String(),
			UpdatedAt:   list.UpdatedAt.String(),
			Lessons:     nil,
			Enrollments: nil,
		}
		allCourses = append(allCourses, user)
	}

	return allCourses, nil
}
