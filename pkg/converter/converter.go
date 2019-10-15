package converter

import (
	"AttYamlToJson/pkg/constants"
	"AttYamlToJson/pkg/structs"
	"fmt"
	"github.com/ghodss/yaml"
	"net/http"
)

func YamlToJson(rawData []byte) structs.Employees {
	var employees structs.Employees

	err := yaml.Unmarshal(rawData, &employees)
	if err != nil {
		return structs.Employees{
			Employees: nil,
			ErrorMessage: structs.ErrorContainer{
				RootCause:  err.Error(),
				StatusCode: http.StatusInternalServerError,
				Trace:      constants.SystemError,
			},
		}
		fmt.Println("error in unmarshalling raw yaml data to json: ", err.Error())
	}

	// checking if there is no data
	if len(employees.Employees) == 0 {
		return structs.Employees{
			Employees: nil,
			ErrorMessage: structs.ErrorContainer{
				RootCause:  "no data found",
				StatusCode: http.StatusNoContent,
			},
		}
	}

	// We came here with no errors
	return employees
}
