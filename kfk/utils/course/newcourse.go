package course

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/laurentino14/user/kfk"
	"github.com/laurentino14/user/prisma"
	"github.com/laurentino14/user/prisma/connect"
)

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

func NewCourseUtil(connect *connect.DB, data kfk.MessageNewCourse, ctx context.Context) {

	a, err := connect.Client.Course.CreateOne(
		prisma.Course.Title.Set(data.Message.Title),
		prisma.Course.Slug.Set(data.Message.Slug),
		prisma.Course.Description.Set(data.Message.Description),
		prisma.Course.Image.Set(data.Message.Image),
		prisma.Course.CreatedAt.Set(time.Now()),
		prisma.Course.UpdatedAt.Set(time.Now()),
		prisma.Course.ID.Set(data.Message.ID),
	).Exec(ctx)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(a.ID)
}
