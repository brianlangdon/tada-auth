package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/brianlangdon/tada-auth/graph"
	"github.com/brianlangdon/tada-auth/internal/auth"
	database "github.com/brianlangdon/tada-auth/internal/pkg/db/mysql"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

const defaultPort = "4000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	// turn off cors for development
	router.Use(cors.AllowAll().Handler)
	router.Use(auth.Middleware())

	db, err := database.InitDB("root:dbpass@(localhost:3306)/tada")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.CloseDB(db)
	database.Migrate(db)

	server := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.Resolver{DB: db},
			},
		),
	)
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe("localhost:"+port, router))
}
