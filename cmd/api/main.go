package main

import (
	"go-dock/internal/db"
	"go-dock/internal/env"
	"go-dock/internal/store"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// loading dotenv file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://loginuser:adminpassword@localhost:5433/social?sslmode=disable"),
			maxOpenConns: 30,
			maxIdleConns: 30,
			maxIdleTime:  "15min",
		},
	}

	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	log.Println("DB connection pool established")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}
	mux := app.mount()
	log.Fatal(app.run(mux))

}
