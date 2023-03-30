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

	router.Use(cors.AllowAll().Handler)
	router.Use(auth.Middleware())

	// CORS setting
	//cors := cors.AllowAll()

	//New(cors.Options{
	// Use this to allow specific origin hosts
	//	AllowedOrigins: []string{"https://*", "http://*"},
	// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
	//	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	//	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	//	ExposedHeaders:   []string{"Link"},
	//	AllowCredentials: true,
	//	MaxAge:           300, // Maximum value not ignored by any of major browsers
	//}//)

	database.InitDB()
	defer database.CloseDB()
	database.Migrate()

	server := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe("localhost:"+port, router))
}
