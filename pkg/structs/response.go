package structs

type Employees struct {
	Employees    []Employee `json:"employees,omitempty"`
	ErrorMessage ErrorLog   `json:"ErrorMessage,omitempty"`
}

type Employee struct {
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	ID        string `json:"id,omitempty"`
}

type ErrorLog struct {
	StatusCode int    `json:"statusCode,omitempty"`
	RootCause  string `json:"rootCause,omitempty"`
	Trace      string `json:"trace,omitempty"`
}
