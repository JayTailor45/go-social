package main

import (
	"log"

	"github.com/JayTailor45/go-social/internal/db"
	"github.com/JayTailor45/go-social/internal/env"
	"github.com/JayTailor45/go-social/internal/store"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

const version = "0.0.1"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conf := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://user:adminpassword@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15min"),
		},
		env: env.GetString("ENV", "local"),
	}

	// Logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	// database
	db, err := db.New(
		conf.db.addr,
		conf.db.maxOpenConns,
		conf.db.maxIdleConns,
		conf.db.maxIdleTime,
	)

	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()
	logger.Info("database connection pool established")

	store := store.NewPostgresStorage(db)

	app := &application{
		config: conf,
		store:  store,
		logger: logger,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))

}
