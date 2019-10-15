package models

type Employees struct {
	Employees    []Employee `json:"employee,omitempty"`
	ErrorMessage ErrorLog   `json:"ErrorMessage,omitempty"`
}

type Employee struct {
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	ID        string `json:"id,omitempty"`
}

type ErrorLog struct {
	StatusCode string `json:"statusCode,omitempty"`
	RootCause  string `json:"rootCause,omitempty"`
}
