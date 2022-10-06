package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/laurentino14/user/graph"
	"github.com/laurentino14/user/graph/generated"
	"github.com/laurentino14/user/prisma/connect"
	"github.com/laurentino14/user/repositories"
	"github.com/laurentino14/user/services"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"os/signal"
)

const defaultPort = "3131"

func main() {
	connect := connect.NewPrismaConnect()
	a := os.Getenv("SECRET")
	fmt.Println(a)
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		LessonService:     services.NewLessonService(repositories.NewLessonRepository(connect)),
		CourseService:     services.NewCourseService(repositories.NewCourseRepository(connect)),
		StepService:       services.NewStepService(repositories.NewStepRepository(connect)),
		UserService:       services.NewUserService(repositories.NewUserRepository(connect)),
		EnrollmentService: services.NewEnrollmentService(repositories.NewEnrollmentRepository(connect)),
		AuthService:       services.NewAuthService(repositories.NewAuthRepository(connect)),
	}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	router.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	var listenAddress = flag.String("listen", ":3131", "Listen address.")

	flag.Parse()

	httpServer := http.Server{
		Addr: *listenAddress,
	}
	idleConnectionsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		if err := httpServer.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP Server Shutdown Error: %v", err)
		}
		close(idleConnectionsClosed)
	}()

	if err := http.ListenAndServe(":3131", cors.Default().Handler(router)); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe Error: %v", err)
	}

	<-idleConnectionsClosed

	connect.Disconnect()

}
