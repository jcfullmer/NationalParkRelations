package main

import (
	"log"
	"net/http"
)

type Config struct {
	Port string
}

func main() {
	conf := Config{
		Port: ":8080",
	}
	mux := http.NewServeMux()

	mux.HandleFunc("PUT /", handlerRoot)

	log.Printf("Listening on port %s", conf.Port)
	err := http.ListenAndServe(conf.Port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
