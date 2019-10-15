package main

import (
	"AttYamlToJson/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/employees", routes.Employees).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8000", router))
}
