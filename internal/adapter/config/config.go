package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Config *Container

type (
	App struct {
		Name string
		Env  string
	}

	DB struct {
		Engine string
		DSN    string
	}

	HTTP struct {
		Env            string
		Address        string
		AllowedOrigins string
	}

	API struct {
		URL   string
		TOKEN string
	}

	Container struct {
		App  *App
		DB   *DB
		HTTP *HTTP
		API  *API
	}
)

func init() {
	if os.Getenv("APP_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Fatal(err)
		}
	}

	app := &App{
		Name: os.Getenv("APP_NAME"),
		Env:  os.Getenv("APP_ENV"),
	}

	db := &DB{
		Engine: os.Getenv("DB_ENGINE"),
		DSN:    os.Getenv("DB_DSN"),
	}

	http := &HTTP{
		Env:            os.Getenv("APP_ENV"),
		Address:        os.Getenv("HTTP_ADDRESS"),
		AllowedOrigins: os.Getenv("HTTP_ALLOWED_ORIGINS"),
	}

	api := &API{
		URL:   os.Getenv("MELI_API"),
		TOKEN: os.Getenv("MELI_TOKEN"),
	}

	Config = &Container{
		app, db, http, api,
	}
}
