package main

import "net/http"

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is the root."))
}
