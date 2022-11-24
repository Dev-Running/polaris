package services

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/laurentino14/user/graph/model"
	"github.com/laurentino14/user/repositories"
)

type ILessonService interface {
	Create(input model.NewLesson, ctx context.Context) (*model.Lesson, error)
	GetAll(ctx context.Context) ([]*model.Lesson, error)
}

type LessonService struct {
	LessonRepository *repositories.LessonRepository
	KAFKA            *kafka.Consumer
}

func NewLessonService(lessonRepository *repositories.LessonRepository, k *kafka.Consumer) *LessonService {
	return &LessonService{
		LessonRepository: lessonRepository,
		KAFKA:            k,
	}
}

func (l *LessonService) Create(input model.NewLesson, ctx context.Context) (*model.Lesson, error) {

	return l.LessonRepository.Create(input, ctx)
}

func (l *LessonService) GetAll(ctx context.Context) ([]*model.Lesson, error) {
	return l.LessonRepository.GetAll(ctx)
}
