package main

import (
	"github.com/laurentino14/user/database"
)

func main() {

	database.DB().AutoMigrate(&database.User{})
	database.DB().AutoMigrate(&database.Course{})
	database.DB().AutoMigrate(&database.Module{})
	database.DB().AutoMigrate(&database.Lesson{})
	database.DB().AutoMigrate(&database.Enrollment{})

}
