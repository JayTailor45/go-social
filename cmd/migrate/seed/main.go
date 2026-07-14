package main

import (
	"github.com/JayTailor45/go-social/internal/db"
	"github.com/JayTailor45/go-social/internal/env"
	"github.com/JayTailor45/go-social/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable")
	println("Connecting to database at : ", addr)
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	store := store.NewPostgresStorage(conn)

	db.Seed(store, conn)
}
