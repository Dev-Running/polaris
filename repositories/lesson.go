package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/laurentino14/user/graph/model"
	"github.com/laurentino14/user/prisma"
	"github.com/laurentino14/user/prisma/connect"
	"github.com/laurentino14/user/repositories/utils"
)

type ILessonRepository interface {
	Create(input model.NewLesson, ctx context.Context) (*model.Lesson, error)
	GetAll(ctx context.Context) ([]*model.Lesson, error)
}

type LessonRepository struct {
	DB *connect.DB
}

type Message struct {
	Type string
	Ok   string
}

func kafkaRun(c *kafka.Consumer, run bool) {
	for run {

		ev, err := c.ReadMessage(100 * time.Millisecond)
		if err != nil {
			// Errors are informational and automatically handled by the consumer
			continue
		}
		var dados Message
		erro := json.Unmarshal(ev.Value, &dados)
		if erro != nil {
			log.Println(erro)
		}

		if dados.Type == "create" {
			fmt.Println("e pra dar create")
		}
		if dados.Type == "update" {
			fmt.Println("e pra dar update")
		}

	}
	c.Close()
}

// NewLessonRepository implements LessonRepository
func NewLessonRepository(db *connect.DB) *LessonRepository {

	k, err := kafka.NewConsumer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9091,localhost:9092,localhost:9093",
		"group.id":          "polaris",
		"auto.offset.reset": "earliest"})
	if err != nil {
		log.Fatal(err)
	}

	err = k.SubscribeTopics([]string{"polaris"}, nil)

	go kafkaRun(k, true)
	return &LessonRepository{
		DB: db,
	}
}

// Create implements LessonRepository
func (r *LessonRepository) Create(input model.NewLesson, ctx context.Context) (*model.Lesson, error) {
	exec, err := r.DB.Client.Lesson.CreateOne(
		prisma.Lesson.Title.Set(*input.Title),
		prisma.Lesson.Slug.Set(*input.Slug),
		prisma.Lesson.Description.Set(*input.Description),
		prisma.Lesson.Link.Set(*input.Link),
		prisma.Lesson.Step.Link(prisma.Step.ID.Equals(*input.StepID)),
		prisma.Lesson.Course.Link(prisma.Course.ID.Equals(*input.CourseID)),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	lessonData := &model.Lesson{
		ID:          exec.ID,
		Title:       exec.Title,
		Slug:        exec.Slug,
		Link:        exec.Link,
		CreatedAt:   exec.CreatedAt.String(),
		UpdatedAt:   utils.ExtractData(exec.UpdatedAt),
		StepID:      exec.StepID,
		CourseID:    exec.CourseID,
		Description: exec.Description,
	}
	return lessonData, nil
}

// GetAll implements LessonRepository
func (r *LessonRepository) GetAll(ctx context.Context) ([]*model.Lesson, error) {
	exec, err := r.DB.Client.Lesson.FindMany().Take(10).Exec(ctx)

	if err != nil {
		return nil, err
	}

	allLessons := []*model.Lesson{}

	for _, list := range exec {

		lesson := &model.Lesson{
			ID:          list.ID,
			Title:       list.Title,
			Slug:        list.Slug,
			Link:        list.Link,
			CreatedAt:   list.CreatedAt.String(),
			UpdatedAt:   utils.ExtractData(list.UpdatedAt),
			StepID:      list.StepID,
			Description: list.Description,
		}
		allLessons = append(allLessons, lesson)
	}

	return allLessons, nil
}
