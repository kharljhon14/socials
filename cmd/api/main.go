package main

import "log"

func main() {
	cfg := config{
		addr: ":8080",
	}

	app := &application{
		config: cfg,
	}

	handler := app.mount()
	log.Fatal(app.serve(handler))

}
