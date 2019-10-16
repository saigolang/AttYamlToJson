package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEmployees(t *testing.T) {
	req, err := http.NewRequest("GET", "/employees", nil)
	if err != nil {
		t.Fatal(err)
	}

	// creating a response recorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Employees)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
