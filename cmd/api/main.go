package main

import (
	"log"

	"github.com/kharljhon14/socials/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	app := &application{
		config: cfg,
	}

	handler := app.mount()
	log.Fatal(app.serve(handler))

}
