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
		if msg.TypeMessage == "update" {
			fmt.Println("e pra dar update")
		}

	}
	c.Close()
}
