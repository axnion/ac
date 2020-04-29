package main

import (
	"log"
	"net/http"

	"github.com/axnion/ac/router"
	"github.com/google/jsonapi"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	router.UsersRouter(api)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func globalHeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", jsonapi.MediaType)
		next.ServeHTTP(w, r)
	})
}
