package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/laurentino14/user/graph/generated"
	"github.com/laurentino14/user/graph/model"
)

// Authentication is the resolver for the authentication field.
func (r *mutationResolver) Authentication(ctx context.Context, input *model.AuthenticationInput) (*model.User, error) {
	authData, err := r.AuthService.Auth(input, ctx)
	if err != nil {
		return nil, err
	}
	return authData, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	userData, err := r.UserService.Create(input, ctx)
	if err != nil {
		return nil, fmt.Errorf("Já existe um usuário utilizando esse e-mail ou telefone")
	}

	return userData, nil
}

// CreateCourse is the resolver for the createCourse field.
func (r *mutationResolver) CreateCourse(ctx context.Context, input model.NewCourse) (*model.Course, error) {
	courseData, err := r.CourseService.Create(input, ctx)
	if err != nil {
		return nil, fmt.Errorf("Este curso já existe!")
	}

	return courseData, err
}

// CreateStep is the resolver for the createStep field.
func (r *mutationResolver) CreateStep(ctx context.Context, input model.NewStep) (*model.Step, error) {
	stepData, err := r.StepService.Create(input, ctx)
	if err != nil {
		return nil, err
	}

	return stepData, err
}

// CreateLesson is the resolver for the createLesson field.
func (r *mutationResolver) CreateLesson(ctx context.Context, input model.NewLesson) (*model.Lesson, error) {
	lessonData, err := r.LessonService.Create(input, ctx)
	if lessonData == nil {
		return nil, err
	}

	return lessonData, err
}

// CreateEnrollment is the resolver for the createEnrollment field.
func (r *mutationResolver) CreateEnrollment(ctx context.Context, input model.NewEnrollment) (*model.Enrollment, error) {
	enrollmentsData, err := r.EnrollmentService.Create(input, ctx)
	if err != nil {
		return nil, fmt.Errorf("Erro de conexão com o banco de dados")
	}

	return enrollmentsData, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	userData, err := r.UserService.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("Erro de conexão com o banco de dados")
	}

	return userData, nil
}

// Courses is the resolver for the courses field.
func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	coursesData, err := r.CourseService.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("Erro de conexão com o banco de dados")
	}

	return coursesData, nil
}

// Steps is the resolver for the steps field.
func (r *queryResolver) Steps(ctx context.Context) ([]*model.Step, error) {
	stepsData, err := r.StepService.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("Erro de conexão com o banco de dados")
	}

	return stepsData, nil
}

// Lessons is the resolver for the lessons field.
func (r *queryResolver) Lessons(ctx context.Context) ([]*model.Lesson, error) {
	lessonsData, err := r.LessonService.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("Erro de conexão com o banco de dados")
	}

	return lessonsData, nil
}

// Enrollments is the resolver for the enrollments field.
func (r *queryResolver) Enrollments(ctx context.Context) ([]*model.Enrollment, error) {
	enrollmentsData, err := r.EnrollmentService.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("Erro de conexão com o banco de dados")
	}

	return enrollmentsData, nil
}

// UserAuthenticated is the resolver for the userAuthenticated field.
func (r *queryResolver) UserAuthenticated(ctx context.Context, input *model.GetUserAuthInput) (*model.UserAuthenticated, error) {
	userAuthenticated, err := r.AuthService.GetUserAuthenticated(input, ctx)
	if err != nil {
		return nil, err
	}
	return userAuthenticated, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
