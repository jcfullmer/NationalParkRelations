package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/jcfullmer/NationalParkRelations/goBackend/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	db   *database.Queries
	Port string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading environment: %s", err)
		os.Exit(2)
	}
	dbURL := os.Getenv("DB_URL")
	db, err := ConnectToDBAndGetQuery(dbURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	conf := Config{
		db:   db,
		Port: ":8080",
	}
	mux := http.NewServeMux()

	mux.HandleFunc("GET  /park", conf.HandlerGetParks)
	mux.HandleFunc("GET /park/state/{state}", conf.HandlerStateSearch)
	log.Printf("Listening on port %s", conf.Port)
	err = http.ListenAndServe(conf.Port, mux)
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectToDBAndGetQuery(dbURL string) (*database.Queries, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return &database.Queries{}, err
	}
	dbQuery := database.New(db)
	return dbQuery, nil
}
