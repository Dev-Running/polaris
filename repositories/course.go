package repositories

import (
	"context"
	"time"

	"github.com/laurentino14/user/graph/model"
	"github.com/laurentino14/user/prisma"
	"github.com/laurentino14/user/prisma/connect"
	"github.com/laurentino14/user/repositories/utils"
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
		prisma.Course.Title.Set(*input.Title),
		prisma.Course.Slug.Set(*input.Slug),
		prisma.Course.Description.Set(*input.Description),
		prisma.Course.Image.Set(*input.Image),
		prisma.Course.CreatedAt.Set(time.Now()),
		prisma.Course.UpdatedAt.Set(time.Now()),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	enrollments := []*model.Enrollment{}

	for _, l := range exec.RelationsCourse.Enrollment {
		enrollments = append(enrollments, &model.Enrollment{
			ID:        l.ID,
			CreatedAt: l.CreatedAt.String(),
			UpdatedAt: utils.ExtractData(l.UpdatedAt),
			UserID:    l.UserID,
			CourseID:  l.CourseID,
			DeletedAt: utils.ExtractData(l.DeletedAt),
		})
	}

	steps := []*model.Step{}

	for _, s := range exec.RelationsCourse.Step {

		lessons := []*model.Lesson{}

		for _, l := range s.RelationsStep.Lessons {

			lessons = append(lessons, &model.Lesson{
				ID:          l.ID,
				Title:       l.Title,
				Slug:        l.Slug,
				Description: l.Description,
				Link:        l.Link,
				CreatedAt:   l.CreatedAt.String(),
				UpdatedAt:   utils.ExtractData(l.UpdatedAt),
				StepID:      l.StepID,
			})
		}

		steps = append(steps, &model.Step{
			ID:          s.ID,
			Title:       s.Title,
			Slug:        s.Slug,
			Description: s.Description,
			CreatedAt:   s.CreatedAt.String(),
			UpdatedAt:   utils.ExtractData(s.UpdatedAt),
			Lessons:     lessons,
			CourseID:    s.CourseID,
		})
	}

	clLessons := []*model.Lesson{}

	for _, cl := range exec.RelationsCourse.Lesson {
		clLessons = append(clLessons, &model.Lesson{
			ID:          cl.ID,
			Title:       cl.Title,
			Slug:        cl.Slug,
			Description: cl.Description,
			Link:        cl.Link,
			CreatedAt:   cl.CreatedAt.String(),
			UpdatedAt:   utils.ExtractData(cl.UpdatedAt),
			StepID:      cl.StepID,
		})

	}

	courseData := &model.Course{
		ID:          exec.ID,
		Title:       exec.Title,
		Slug:        exec.Slug,
		Image:       exec.Image,
		Description: exec.Description,
		CreatedAt:   exec.CreatedAt.String(),
		UpdatedAt:   utils.ExtractData(exec.UpdatedAt),
		Lessons:     clLessons,
		Steps:       steps,
		Enrollments: enrollments,
	}
	return courseData, nil
}

func (r *CourseRepository) GetAll(ctx context.Context) ([]*model.Course, error) {
	exec, err := r.DB.Client.Course.FindMany().With(prisma.Course.Enrollment.Fetch(), prisma.Course.Step.Fetch().With(prisma.Step.Lessons.Fetch()), prisma.Course.Lesson.Fetch()).Exec(ctx)

	if err != nil {
		return nil, err
	}

	allCourses := []*model.Course{}
	allLessons := []*model.Lesson{}
	for _, list := range exec {

		for _, l := range list.RelationsCourse.Lesson {

			allLessons = append(allLessons, &model.Lesson{
				ID:          l.ID,
				Title:       l.Title,
				Slug:        l.Slug,
				Description: l.Description,
				Link:        l.Link,
				CreatedAt:   l.CreatedAt.String(),
				UpdatedAt:   utils.ExtractData(l.UpdatedAt),
				StepID:      l.StepID,
			})
		}

		enrollments := []*model.Enrollment{}

		for _, l := range list.RelationsCourse.Enrollment {
			enrollments = append(enrollments, &model.Enrollment{
				ID:        l.ID,
				CreatedAt: l.CreatedAt.String(),
				UpdatedAt: utils.ExtractData(l.UpdatedAt),
				UserID:    l.UserID,
				CourseID:  l.CourseID,
				DeletedAt: utils.ExtractData(l.DeletedAt),
			})
		}

		allSteps := []*model.Step{}
		for _, s := range list.RelationsCourse.Step {

			lessons := []*model.Lesson{}

			for _, l := range s.RelationsStep.Lessons {
				lessons = append(lessons, &model.Lesson{
					ID:          l.ID,
					Title:       l.Title,
					Slug:        l.Slug,
					Description: l.Description,
					Link:        l.Link,
					CreatedAt:   l.CreatedAt.String(),
					UpdatedAt:   utils.ExtractData(l.UpdatedAt),
					StepID:      l.StepID,
				})
			}

			allSteps = append(allSteps, &model.Step{
				ID:          s.ID,
				Title:       s.Title,
				Slug:        s.Slug,
				Description: s.Description,
				CreatedAt:   s.CreatedAt.String(),
				UpdatedAt:   utils.ExtractData(s.UpdatedAt),
				Lessons:     lessons,
				CourseID:    s.CourseID,
			})
		}

		COURSE := &model.Course{
			ID:          list.ID,
			Title:       list.Title,
			Slug:        list.Slug,
			Description: list.Description,
			Image:       list.Image,
			CreatedAt:   list.CreatedAt.String(),
			UpdatedAt:   utils.ExtractData(list.UpdatedAt),
			Lessons:     allLessons,
			Steps:       allSteps,
			Enrollments: enrollments,
		}
		allCourses = append(allCourses, COURSE)
	}

	return allCourses, nil
}
