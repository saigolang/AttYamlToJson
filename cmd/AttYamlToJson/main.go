package main

import (
	"AttYamlToJson/pkg/converter"
	"AttYamlToJson/pkg/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	// todo configure logging here
	router.HandleFunc("/employees", employees).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func employees(w http.ResponseWriter, r *http.Request) {
	rawData, err := ioutil.ReadFile("employee.yaml")
	if err != nil {
		fmt.Println("error in converting yaml file to bytes is ", err.Error())
	}
	// convert yaml to json
	response := converter.YamlToJson(rawData)
	// writing final response to server
	writeResponseToServer(w, response)
}
func writeResponseToServer(w http.ResponseWriter, finalResponse models.Employees) {
	respBodyBytes := new(bytes.Buffer)
	err := json.NewEncoder(respBodyBytes).Encode(&finalResponse)
	if err != nil {
		log.Println("error in marshalling the response")
	}
	w.WriteHeader(http.StatusOK)
	w.Write(respBodyBytes.Bytes())
}
