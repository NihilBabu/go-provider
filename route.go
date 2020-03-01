package main

import (
	"github.com/NihilBabu/go-provider/handlers"
	"github.com/gorilla/mux"
)

//ProviderRoutes is ..
func ProviderRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.IndexHandler).Methods("GET")
	return r
}
