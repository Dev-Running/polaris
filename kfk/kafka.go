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
	Type string
	Ok   string
}

func KafkaRun(c *kafka.Consumer, run bool, connect *connect.DB, ctx context.Context) {

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
			a, err := connect.Client.Course.CreateOne(
				prisma.Course.Title.Set("123"),
				prisma.Course.Slug.Set(""),
				prisma.Course.Description.Set(""),
				prisma.Course.Image.Set(""),
				prisma.Course.CreatedAt.Set(time.Now()),
				prisma.Course.UpdatedAt.Set(time.Now()),
			).Exec(ctx)
			if err != nil {
				log.Println(err)
			}

			fmt.Println(a)
		}
		if dados.Type == "update" {
			fmt.Println("e pra dar update")
		}

	}
	c.Close()
}
