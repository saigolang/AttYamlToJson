package models

type Employees struct {
	Employees []Employee `json:employee`
}
type Employee struct {
	FirstName string `json:firstName`
	LastName  string `json:lastName`
	ID        string `json:id`
}
