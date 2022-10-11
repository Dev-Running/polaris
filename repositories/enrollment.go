package repositories

import (
	"context"
	"time"

	"github.com/laurentino14/user/graph/model"
	"github.com/laurentino14/user/prisma"
	"github.com/laurentino14/user/prisma/connect"
	"github.com/laurentino14/user/repositories/utils"
)

type IEnrollmentRepository interface {
	Create(input model.NewEnrollment, ctx context.Context) (*model.Enrollment, error)
	GetAll(ctx context.Context) ([]*model.Enrollment, error)
}

type EnrollmentRepository struct {
	DB *connect.DB
}

func NewEnrollmentRepository(db *connect.DB) *EnrollmentRepository {
	return &EnrollmentRepository{DB: db}
}

func (r *EnrollmentRepository) Create(input model.NewEnrollment, ctx context.Context) (*model.Enrollment, error) {
	exec, err := r.DB.Client.Enrollment.CreateOne(
		prisma.Enrollment.CreatedAt.Set(time.Now()),
		prisma.Enrollment.User.Link(prisma.User.ID.Equals(input.UserID)),
		prisma.Enrollment.Course.Link(prisma.Course.ID.Equals(input.CourseID)),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	enrollmentData := &model.Enrollment{
		ID:        exec.ID,
		CreatedAt: exec.CreatedAt.String(),
		UpdatedAt: utils.ExtractData(exec.UpdatedAt),
		DeletedAt: utils.ExtractData(exec.DeletedAt),
		UserID:    exec.UserID,
		CourseID:  exec.CourseID,
	}
	return enrollmentData, nil
}

func (r *EnrollmentRepository) GetAll(ctx context.Context) ([]*model.Enrollment, error) {
	exec, err := r.DB.Client.Enrollment.FindMany().Take(10).Exec(ctx)

	if err != nil {
		return nil, err
	}

	allEnrollments := []*model.Enrollment{}

	for _, list := range exec {

		user := &model.Enrollment{
			ID:        list.ID,
			CreatedAt: list.CreatedAt.String(),
			UpdatedAt: utils.ExtractData(list.UpdatedAt),
			DeletedAt: utils.ExtractData(list.DeletedAt),
			UserID:    list.UserID,
			CourseID:  list.CourseID,
		}
		allEnrollments = append(allEnrollments, user)
	}

	return allEnrollments, nil
}
