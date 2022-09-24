package lesson

import (
	"context"
	"github.com/laurentino14/user/graph/model"
	"github.com/laurentino14/user/prisma"
	"github.com/laurentino14/user/prisma/connect"
)

func (c *connect.DB.DB) CreateL(input model.NewLesson, ctx context.Context) (*model.Lesson, error) {
	client := c..Client

}
