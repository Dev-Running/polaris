package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/laurentino14/user/services/lesson"

	"github.com/laurentino14/user/graph/generated"
	"github.com/laurentino14/user/graph/model"
	"github.com/laurentino14/user/services/course"
	"github.com/laurentino14/user/services/step"
	"github.com/laurentino14/user/services/user"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	userData, err := user.CreateUser(input, ctx)
	if err != nil {
		return nil, fmt.Errorf("Já existe um usuário utilizando esse e-mail ou telefone")
	}

	return userData, nil
}

// CreateCourse is the resolver for the createCourse field.
func (r *mutationResolver) CreateCourse(ctx context.Context, input model.NewCourse) (*model.Course, error) {
	courseData, err := course.CreateCourse(input, ctx)
	if err != nil {
		return nil, fmt.Errorf("Este curso já existe!")
	}

	return courseData, err
}

// CreateStep is the resolver for the createStep field.
func (r *mutationResolver) CreateStep(ctx context.Context, input model.NewStep) (*model.Step, error) {
	stepData, err := step.CreateStep(input, ctx)
	if stepData == nil {
		return nil, err
	}

	return stepData, err
}

// CreateLesson is the resolver for the createLesson field.
func (r *mutationResolver) CreateLesson(ctx context.Context, input model.NewLesson) (*model.Lesson, error) {
	lessonData, err := lesson.CreateLesson(input, ctx, r.Connect)
	if lessonData == nil {
		return nil, err
	}

	return lessonData, err
}

// CreateEnrollment is the resolver for the createEnrollment field.
func (r *mutationResolver) CreateEnrollment(ctx context.Context, input model.NewEnrollment) (*model.Enrollment, error) {
	panic(fmt.Errorf("not implemented: CreateEnrollment - createEnrollment"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	userData := user.GetAllUsers(ctx)
	if userData == nil {
		return nil, fmt.Errorf("Erro de conexão com o banco de dados")
	}

	return userData, nil
}

// Courses is the resolver for the courses field.
func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	coursesData := course.GetAllCourses(ctx)
	if coursesData == nil {
		return nil, fmt.Errorf("Erro de conexão com o banco de dados")
	}

	return coursesData, nil
}

// Steps is the resolver for the steps field.
func (r *queryResolver) Steps(ctx context.Context) ([]*model.Step, error) {
	stepsData := step.GetAllSteps(ctx)
	if stepsData == nil {
		return nil, fmt.Errorf("Erro de conexão com o banco de dados")
	}

	return stepsData, nil
}

// Lessons is the resolver for the lessons field.
func (r *queryResolver) Lessons(ctx context.Context) ([]*model.Lesson, error) {
	lessonsData := lesson.GetAllLessons(ctx)
	if lessonsData == nil {
		return nil, fmt.Errorf("Erro de conexão com o banco de dados")
	}

	return lessonsData, nil
}

// Enrollments is the resolver for the enrollments field.
func (r *queryResolver) Enrollments(ctx context.Context) ([]*model.Enrollment, error) {
	panic(fmt.Errorf("not implemented: Enrollments - enrollments"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct {
	*Resolver
}
