package kfk

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/laurentino14/user/prisma"
	"github.com/laurentino14/user/prisma/connect"
)

type Message struct {
	TypeMessage string ``
}

type MessageNewCourse struct {
	TypeMessage string    `json:"typeMessage"`
	Message     NewCourse `json:"message"`
}

type NewCourse struct {
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Image       string `json:"image"`
	ID          string `json:"id"`
}
type NewModule struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	CourseID    string `json:"courseID"`
}

type MessageNewModule struct {
	TypeMessage string    `json:"typeMessage"`
	Message     NewModule `json:"message"`
}

type NewLesson struct {
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Link        string `json:"link"`
	ModuleID    string `json:"moduleID"`
	CourseID    string `json:"courseID"`
	ID          string `json:"id"`
}
type MessageNewLesson struct {
	TypeMessage string    `json:"typeMessage"`
	Message     NewLesson `json:"message"`
}

func KafkaRun(c *kafka.Consumer, run bool, connect *connect.DB, ctx context.Context) {

	for run {

		ev, err := c.ReadMessage(100 * time.Millisecond)
		if err != nil {
			// Errors are informational and automatically handled by the consumer
			continue
		}
		var msg Message
		erro := json.Unmarshal(ev.Value, &msg)
		if erro != nil {
			log.Println(erro)
		}

		if msg.TypeMessage == "teste" {
			fmt.Println("teste pegou")
		}

		if msg.TypeMessage == "newCourse" {
			var newCourse MessageNewCourse
			json.Unmarshal(ev.Value, &newCourse)

			a, err := connect.Client.Course.CreateOne(
				prisma.Course.Title.Set(newCourse.Message.Title),
				prisma.Course.Slug.Set(newCourse.Message.Slug),
				prisma.Course.Description.Set(newCourse.Message.Description),
				prisma.Course.Image.Set(newCourse.Message.Image),
				prisma.Course.CreatedAt.Set(time.Now()),
				prisma.Course.UpdatedAt.Set(time.Now()),
				prisma.Course.ID.Set(newCourse.Message.ID),
			).Exec(ctx)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(a.ID)
			// var a NewCourse
			// errojson.Unmarshal(dados.message, &a)
			// fmt.Println(a)
		}
		if msg.TypeMessage == "newModule" {
			var newModule MessageNewModule
			json.Unmarshal(ev.Value, &newModule)

			a, err := connect.Client.Step.CreateOne(prisma.Step.Title.Set(newModule.Message.Title),
				prisma.Step.Description.Set(newModule.Message.Description),
				prisma.Step.Slug.Set(newModule.Message.Slug),
				prisma.Step.Course.Link(prisma.Course.ID.Equals(newModule.Message.CourseID)),
				prisma.Step.ID.Set(newModule.Message.ID),
				prisma.Step.CreatedAt.Set(time.Now())).Exec(ctx)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(a.ID)

		}

		if msg.TypeMessage == "newLesson" {
			var newLesson MessageNewLesson
			json.Unmarshal(ev.Value, &newLesson)

			a, err := connect.Client.Lesson.CreateOne(prisma.Lesson.Title.Set(newLesson.Message.Title),
				prisma.Lesson.Slug.Set(newLesson.Message.Slug),
				prisma.Lesson.Description.Set(newLesson.Message.Description),
				prisma.Lesson.Link.Set(newLesson.Message.Link),
				prisma.Lesson.Step.Link(prisma.Step.ID.Equals(newLesson.Message.ModuleID)),
				prisma.Lesson.Course.Link(prisma.Course.ID.Equals(newLesson.Message.CourseID)),
				prisma.Lesson.CourseID.Set(newLesson.Message.CourseID),
				prisma.Lesson.ID.Set(newLesson.Message.ID),
			).Exec(ctx)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(a.ID)

		}

	}
	c.Close()
}
