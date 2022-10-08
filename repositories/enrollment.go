package repositories

import (
	"context"
	"github.com/laurentino14/user/graph/model"
	"github.com/laurentino14/user/prisma"
	"github.com/laurentino14/user/prisma/connect"
	"time"
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

	r.DB.Client.User.UpsertOne(prisma.User.ID.Equals(exec.UserID)).Update(prisma.User.Enrollment.Link(prisma.Enrollment.ID.Equals(exec.ID)))

	up, _ := exec.UpdatedAt()
	upd := up.String()
	de, _ := exec.DeletedAt()
	del := de.String()
	enrollmentData := &model.Enrollment{
		ID:        exec.ID,
		CreatedAt: exec.CreatedAt.String(),
		UpdatedAt: upd,
		DeletedAt: del,
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

		up, _ := list.UpdatedAt()

		updatedAt := up.String()

		del, _ := list.DeletedAt()

		deletedAt := del.String()

		user := &model.Enrollment{
			ID:        list.ID,
			CreatedAt: list.CreatedAt.String(),
			UpdatedAt: updatedAt,
			DeletedAt: deletedAt,
			UserID:    list.UserID,
			CourseID:  list.CourseID,
		}
		allEnrollments = append(allEnrollments, user)
	}

	return allEnrollments, nil
}
