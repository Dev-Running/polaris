package graph

import (
	"github.com/laurentino14/user/prisma/connect"
	services2 "github.com/laurentino14/user/repositories"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Connect *connect.DB

	Lesson *services2.ILessonRepository
}
