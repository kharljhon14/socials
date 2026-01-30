package main

import (
	"log"

	"github.com/kharljhon14/socials/internal/env"
	"github.com/kharljhon14/socials/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}
	store := store.NewPostgresStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	handler := app.mount()
	log.Fatal(app.serve(handler))

}
