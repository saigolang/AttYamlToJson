package models

type Employees struct {
	Employees    []Employee `json:employee`
	ErrorMessage ErrorLog   `json:ErrorMessage`
}

type Employee struct {
	FirstName string `json:firstName`
	LastName  string `json:lastName`
	ID        string `json:id`
}

type ErrorLog struct {
	StatusCode string `json:statusCode`
	RootCause  string `json:rootCause`
}
