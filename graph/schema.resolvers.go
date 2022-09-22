package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/laurentino14/user/useCases/user"

	"github.com/laurentino14/user/graph/generated"
	"github.com/laurentino14/user/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	userData := user.CreateUser(input, ctx)
	if userData == nil {
		return nil, fmt.Errorf("Erro de conexão com o banco de dados")
	}

	return userData, nil
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
	panic(fmt.Errorf("not implemented: Courses - courses"))
}

// Modules is the resolver for the modules field.
func (r *queryResolver) Modules(ctx context.Context) ([]*model.Module, error) {
	panic(fmt.Errorf("not implemented: Modules - modules"))
}

// Lessons is the resolver for the lessons field.
func (r *queryResolver) Lessons(ctx context.Context) ([]*model.Lesson, error) {
	panic(fmt.Errorf("not implemented: Lessons - lessons"))
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
type queryResolver struct{ *Resolver }
