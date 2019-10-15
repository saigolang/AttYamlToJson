package structs

type Employees struct {
	Employees    []Employee     `json:"employees,omitempty"`
	ErrorMessage ErrorContainer `json:"ErrorMessage,omitempty"`
}

type Employee struct {
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	ID        string `json:"id,omitempty"`
}

type ErrorContainer struct {
	StatusCode int    `json:"statusCode,omitempty"`
	RootCause  string `json:"rootCause,omitempty"`
	Trace      string `json:"trace,omitempty"`
}
