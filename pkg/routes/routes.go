package routes

import (
	"AttYamlToJson/pkg/constants"
	"AttYamlToJson/pkg/converter"
	"AttYamlToJson/pkg/structs"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Employees(w http.ResponseWriter, r *http.Request) {
	rawData, err := ioutil.ReadFile("employee.yaml")
	if err != nil {
		writeResponseToServer(w, structs.Employees{
			Employees: nil,
			ErrorMessage: structs.ErrorContainer{
				RootCause:  err.Error(),
				Trace:      constants.SystemError,
				StatusCode: http.StatusInternalServerError,
			},
		})
		fmt.Println("error in converting yaml file to bytes: ", err.Error())
	}

	// convert yaml to json
	response := converter.YamlToJson(rawData)

	// writing final response to server
	writeResponseToServer(w, response)
}

func writeResponseToServer(w http.ResponseWriter, finalResponse interface{}) {
	respBodyBytes := new(bytes.Buffer)

	err := json.NewEncoder(respBodyBytes).Encode(&finalResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("error in marshalling the response: ", err.Error())
	}

	w.WriteHeader(http.StatusOK)
	w.Write(respBodyBytes.Bytes())
}
