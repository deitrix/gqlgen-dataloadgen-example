package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/deitrix/gqlgen-dataloadgen-example/dataloader"
	"github.com/deitrix/gqlgen-dataloadgen-example/graph"
	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"

	egmysql "github.com/deitrix/gqlgen-dataloadgen-example/store/mysql"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	config := mysql.NewConfig()
	config.Addr = "localhost:3306"
	config.DBName = "network"
	config.User = "root"
	config.Passwd = "1234"

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	store := egmysql.NewStore(db)
	resolver := graph.NewResolver(store)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	r := chi.NewRouter()
	r.Use(dataloader.Middleware(store))

	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
