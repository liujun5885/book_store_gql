package main

//go:generate go run github.com/99designs/gqlgen

import (
	"errors"
	"fmt"
	"github.com/liujun5885/book_store_gql/graph/dataloader"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-pg/pg/v9"
	"github.com/joho/godotenv"
	"github.com/rs/cors"

	"github.com/liujun5885/book_store_gql/constants"
	"github.com/liujun5885/book_store_gql/db"
	"github.com/liujun5885/book_store_gql/db/dborm"
	"github.com/liujun5885/book_store_gql/graph/generated"
	"github.com/liujun5885/book_store_gql/graph/resolver"
	MyMiddleware "github.com/liujun5885/book_store_gql/middleware"
)

const defaultPort = "8000"

func loadEnvs() error {
	err := godotenv.Load(".env")

	requiredKey := []string{
		constants.JWTSecret,
	}

	for _, key := range requiredKey {
		value := os.Getenv(key)
		if value == "" {
			return errors.New(fmt.Sprintf("required Key (%s) does not exist", key))
		}
	}

	return err
}

func main() {
	if err := loadEnvs(); err != nil {
		log.Fatal(err)
		return
	}

	assetConn := db.Conn(&pg.Options{
		User:     "book_store",
		Database: "book_store_assets",
		Password: "aaaaa",
		Addr:     "localhost:5432",
	})
	defer assetConn.Close()

	ugcConn := db.Conn(&pg.Options{
		User:     "book_service",
		Database: "book_store_ugc",
		Password: "aaaaa",
		Addr:     "localhost:5432",
	})
	defer ugcConn.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(MyMiddleware.AuthMiddleware(ugcConn))

	config := generated.Config{Resolvers: &resolver.Resolver{
		ORMBooks:     dborm.Book{DB: assetConn},
		ORMPublisher: dborm.Publisher{DB: assetConn},
		ORMAuthor:    dborm.Author{DB: assetConn},
		ORMUser:      dborm.User{DB: ugcConn},
	}}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", dataloader.DataLoader(assetConn, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
