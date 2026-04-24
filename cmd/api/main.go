package main

import (
	"log"

	"github.com/JayTailor45/go-social/internal/env"
	"github.com/JayTailor45/go-social/internal/store"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conf := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	store := store.NewPostgresStorage(nil)

	app := &application{
		config: conf,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))

}
