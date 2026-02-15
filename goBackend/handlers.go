package main

import (
	"context"
	"net/http"
)

func (conf Config) HandlerRoot(w http.ResponseWriter, r *http.Request) {
	parks, err := conf.db.GetParks(context.Background())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error getting parks", err)
	}
	respondWithJSON(w, http.StatusAccepted, parks)
	// w.Write([]byte("This is the root."))
}

func (conf Config) HandlerStateSearch(w http.ResponseWriter, r *http.Request) {
}
