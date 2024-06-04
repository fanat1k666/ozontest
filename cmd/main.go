package main

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
	"log"
	"net/http"
	"ozon/cfg"
	"ozon/internal/adapters"
	"ozon/internal/controller/server"
	"ozon/internal/gql"
	"ozon/internal/in_memory"
	"ozon/internal/postgres"
)

func main() {

	router, db := initializeAPI()
	defer db.Close()

	log.Fatal(http.ListenAndServe(":4000", router))
}

func initializeAPI() (*chi.Mux, adapters.Storage) {
	env := cfg.NewEnv()
	router := chi.NewRouter()
	var dbInstance adapters.Storage
	if !env.Pg {
		im, err := in_memory.New()
		if err != nil {
			log.Fatal(err)
		}
		dbInstance = im
	} else {
		conn, err := sql.Open("postgres", postgres.ConnString(env.Host, env.Port, env.User, env.Password, env.DbName))
		if err != nil {
			panic(err)
		}
		err = conn.Ping()
		if err != nil {
			panic(err)
		}
		db := postgres.New(conn)
		if err != nil {
			log.Fatal(err)
		}
		dbInstance = db
	}

	rootQuery := gql.NewRoot(dbInstance)
	sc, err := graphql.NewSchema(
		graphql.SchemaConfig{Query: rootQuery.Query},
	)
	if err != nil {
		fmt.Println("Error creating schema: ", err)
	}

	s := server.Server{
		GqlSchema: &sc,
	}

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.StripSlashes,
		middleware.Recoverer,
	)

	router.Post("/graphql", s.GraphQL())

	return router, dbInstance
}
