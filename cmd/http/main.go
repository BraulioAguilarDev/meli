package main

import (
	"context"
	_ "embed"
	"log"
	"meli/internal/adapter/config"
	"meli/internal/adapter/handler/http"
	"meli/internal/adapter/storage/postgres/repository"
	"meli/internal/core/service"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

//go:embed schema.sql
var ddl string

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	db, err := pgx.Connect(ctx, config.DB.DSN)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(ctx, ddl); err != nil {
		log.Fatal(err)
	}

	queries := repository.New(db)
	itemRepository := repository.NewItemRepository(queries)
	itemService := service.ProvideBaseService(itemRepository, config.Meli.URL)
	itemHandler := http.ProvideItemHandler(itemService)

	router, err := http.NewRouter(config.HTTP, *itemHandler)
	if err != nil {
		log.Fatal(err)
	}

	if err := router.Serve(config.HTTP.Address); err != nil {
		log.Fatal(err)
	}
}
