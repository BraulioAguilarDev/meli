package config

import (
	"os"

	"github.com/joho/godotenv"
)

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

	Container struct {
		App  *App
		DB   *DB
		HTTP *HTTP
		Meli *Meli
	}

	Meli struct {
		URL string
	}
)

func New() (*Container, error) {
	if os.Getenv("APP_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			return nil, err
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

	meli := &Meli{
		URL: "https://api.mercadolibre.com",
	}

	return &Container{
		app, db, http, meli,
	}, nil
}
