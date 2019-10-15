package main

import (
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
)

func main() {

	rawData, err := ioutil.ReadFile("employee.yaml")
	if err != nil {
		fmt.Println("error in converting yaml file to bytes is ", err.Error())
	}

	var employee Employees

	err = yaml.Unmarshal(rawData, &employee)
	if err != nil {
		fmt.Println("error in converting yaml to json is ", err.Error())
	}

	fmt.Println(employee)
}

type Employees struct {
	Employees []Employee `json:employee`
}
type Employee struct {
	FirstName string `json:firstName`
	LastName  string `json:lastName`
	ID        string `json:id`
}
