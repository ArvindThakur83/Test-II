package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRouting() {
	r := mux.NewRouter()

	r.HandleFunc("/weapons", CreatWeapon).Methods("POST")
	r.HandleFunc("/weapons", GetWeapons).Methods("GET")
	r.HandleFunc("/weapons/{id}", GetWeaponsById).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
