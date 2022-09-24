package graph

import (
	"github.com/laurentino14/user/prisma/connect"
	"github.com/laurentino14/user/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Connect *connect.DB

	LessonService services.ILessonService
}
