package main

import (
	"AttYamlToJson/pkg/structs"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	logger = &logrus.Logger{}
)

func Test_Employees(t *testing.T) {
	req, err := http.NewRequest("GET", "/employees", nil)
	if err != nil {
		t.Fatal(err)
	}

	// creating a response recorder to record the response
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(Employees(logger))

	handler.ServeHTTP(rr, req)

	// Checking the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func Test_writeResponseToServer(t *testing.T) {
	// creating a response recorder to record the response
	rr := httptest.NewRecorder()
	writeResponseToServer(rr, structs.Employees{}, logger)

	// Checking the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
