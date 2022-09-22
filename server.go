package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/laurentino14/user/graph"
	"github.com/laurentino14/user/graph/generated"
)

const defaultPort = "3131"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	go http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	go http.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	go log.Fatal(http.ListenAndServe(":"+port, nil))
}
