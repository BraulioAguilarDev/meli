package main

import (
	"context"
	_ "embed"
	"log"
	c "meli/internal/adapter/config"
	"meli/internal/adapter/handler/http"
	"meli/internal/adapter/storage/postgres/repository"
	"meli/internal/core/service"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

//go:embed schema.sql
var ddl string

func main() {
	ctx := context.Background()
	// postgres conexion
	db, err := pgx.Connect(ctx, c.Config.DB.DSN)
	if err != nil {
		log.Fatal(err)
	}

	// Run migrations
	if _, err := db.Exec(ctx, ddl); err != nil {
		log.Fatal(err)
	}

	// DI
	queries := repository.New(db)
	itemRepository := repository.NewItemRepository(queries)
	itemService := service.ProvideBaseService(itemRepository)
	itemHandler := http.ProvideItemHandler(itemService)

	router, err := http.NewRouter(c.Config.HTTP, *itemHandler)
	if err != nil {
		log.Fatal(err)
	}

	if err := router.Serve(c.Config.HTTP.Address); err != nil {
		log.Fatal(err)
	}
}
