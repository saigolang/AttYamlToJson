package main

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/saigolang/AttYamlToJson/pkg/constants"
	"github.com/saigolang/AttYamlToJson/pkg/converter"
	"github.com/saigolang/AttYamlToJson/pkg/logging"
	"github.com/saigolang/AttYamlToJson/pkg/structs"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// configure logging
	logger := logging.ConfigureLogging()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/employees", Employees(logger)).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func Employees(logger *logrus.Logger) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rawData, err := ioutil.ReadFile(constants.FileName)
		if err != nil {
			writeResponseToServer(w, structs.Employees{
				Employees: nil,
				ErrorMessage: structs.ErrorContainer{
					RootCause:  err.Error(),
					Trace:      constants.SystemError,
					StatusCode: http.StatusInternalServerError,
				},
			}, logger)
			// logging the error
			logger.WithFields(logrus.Fields{
				constants.RootCause:  err.Error(),
				constants.Trace:      constants.SystemError,
				constants.StatusCode: http.StatusInternalServerError,
			}).Error()
			return
		}

		// convert yaml to json
		response := converter.YamlToJson(rawData, logger)

		// writing final response to server
		writeResponseToServer(w, response, logger)
	}
}

func writeResponseToServer(w http.ResponseWriter, finalResponse interface{}, logger *logrus.Logger) {
	respBodyBytes := new(bytes.Buffer)

	err := json.NewEncoder(respBodyBytes).Encode(&finalResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// logging the error
		logger.WithFields(logrus.Fields{
			constants.RootCause:  err.Error(),
			constants.Trace:      constants.SystemError,
			constants.StatusCode: http.StatusInternalServerError,
		}).Error()
	}

	w.WriteHeader(http.StatusOK)
	w.Write(respBodyBytes.Bytes())
}
