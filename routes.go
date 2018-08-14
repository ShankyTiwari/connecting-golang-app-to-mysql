package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func setStaticFolder(r *mux.Router) {
	fs := http.FileServer(http.Dir("./public/"))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))
}

func routes(r *mux.Router) {

	setStaticFolder(r)

	r.HandleFunc("/", renderHome)

	r.HandleFunc("/users/{name}", getUsers).Methods("GET")

	fmt.Println("Routes are Loded.")
}
