package main

import "net/http"

// use this function signature if need to define an http handler
// for go standard library
// empty struct is used to indicate that this
// handler does not require any input
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, struct{}{})
}
