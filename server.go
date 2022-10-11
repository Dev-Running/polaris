package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/laurentino14/user/graph"
	"github.com/laurentino14/user/graph/generated"
	"github.com/laurentino14/user/prisma/connect"
	"github.com/laurentino14/user/repositories"
	"github.com/laurentino14/user/services"
	"github.com/rs/cors"
)

const defaultPort = "3131"

func main() {
	connect := connect.NewPrismaConnect()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		LessonService:     services.NewLessonService(repositories.NewLessonRepository(connect)),
		CourseService:     services.NewCourseService(repositories.NewCourseRepository(connect)),
		StepService:       services.NewStepService(repositories.NewStepRepository(connect)),
		UserService:       services.NewUserService(repositories.NewUserRepository(connect)),
		EnrollmentService: services.NewEnrollmentService(repositories.NewEnrollmentRepository(connect)),
		AuthService:       services.NewAuthService(repositories.NewAuthRepository(connect)),
	}}))
	http.Handle("/static/", http.StripPrefix("/static/", cors.AllowAll().Handler(http.FileServer(http.Dir("./assets")))))
	http.Handle("/", cors.AllowAll().Handler(playground.Handler("GraphQL playground", "/graphql")))
	http.Handle("/graphql", cors.AllowAll().Handler(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	var listenAddress = flag.String("listen", ":3131", "Listen address.")

	flag.Parse()

	httpServer := http.Server{
		Addr: *listenAddress,
	}
	cors.AllowAll().Handler(srv)
	idleConnectionsClosed := make(chan struct{})

	if err := http.ListenAndServe(":3131", nil); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe Error: %v", err)
	}

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		if err := httpServer.Shutdown(context.Background()); err != nil {
			log.Fatalf("HTTP Server Shutdown Error: %v", err)
		}
		close(idleConnectionsClosed)
	}()

	<-idleConnectionsClosed

	connect.Disconnect()

}
