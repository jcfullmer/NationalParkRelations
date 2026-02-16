package main

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/jcfullmer/NationalParkRelations/goBackend/internal/database"
)

const LimitDefault = 50

type URLParams struct {
	Filter string
	Sort   string
	Limit  int
}

func GetURLParameters(r *http.Request) URLParams {
	vars := r.URL.Query()
	Params := URLParams{
		Filter: vars.Get("filter"),
		Sort:   vars.Get("sort"),
		Limit:  LimitDefault,
	}
	if vars.Get("limit") != "" {
		urlLimit, err := strconv.Atoi(vars.Get("limit"))
		if err != nil {
			log.Printf("Error parsing limit parameter from url: %v", err)
			return URLParams{}
		}
		Params.Limit = urlLimit
	}
	return Params
}

func (conf Config) HandlerGetParks(w http.ResponseWriter, r *http.Request) {
	params := GetURLParameters(r)
	GetParkParams := database.GetParksParams{
		Column1: params.Sort,
		Limit:   int32(params.Limit),
	}
	parks, err := conf.db.GetParks(context.Background(), GetParkParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error getting parks", err)
	}
	respondWithJSON(w, http.StatusAccepted, parks)
	// w.Write([]byte("This is the root."))
}

func (conf Config) HandlerStateSearch(w http.ResponseWriter, r *http.Request) {
	stateCode := r.PathValue("state")
	urlparams := GetURLParameters(r)
	args := database.GetParkByStateParams{
		StateCode: strings.ToUpper(stateCode),
		Limit:     int32(urlparams.Limit),
	}
	parks, err := conf.db.GetParkByState(context.Background(), args)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error getting Request", err)
	}
	respondWithJSON(w, http.StatusAccepted, parks)
}
