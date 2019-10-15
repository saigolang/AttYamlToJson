package main

import (
	"AttYamlToJson/pkg/converter"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	rawData, err := ioutil.ReadFile("employee.yaml")
	if err != nil {
		fmt.Println("error in converting yaml file to bytes is ", err.Error())
	}

	// convert yaml to json
	response := converter.YamlToJson(rawData)

	log.Println("response is ", response)
}

func UnMarhsalYaml() {

}
