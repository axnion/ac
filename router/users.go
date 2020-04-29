package router

import (
	"net/http"

	"github.com/axnion/ac/lib/users"
	"github.com/google/jsonapi"
	"github.com/gorilla/mux"
)

func UsersRouter(r *mux.Router) {
	router := r.PathPrefix("/users").Subrouter()
	router.HandleFunc("", getUsers).Methods(http.MethodGet)
	router.HandleFunc("", postUser).Methods(http.MethodPost)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	// Fetch users
	users := users.GetUsers()

	// Send response
	w.WriteHeader(http.StatusAccepted)
	if err := jsonapi.MarshalPayload(w, users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func postUser(w http.ResponseWriter, r *http.Request) {
	// Create user object from body
	user := new(users.User)
	if err := jsonapi.UnmarshalPayload(r.Body, user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write to datastore
	users.Write(user)

	// Send response
	w.WriteHeader(http.StatusAccepted)
	if err := jsonapi.MarshalPayload(w, user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
