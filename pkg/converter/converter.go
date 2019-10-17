package converter

import (
	"github.com/ghodss/yaml"
	"github.com/saigolang/AttYamlToJson/pkg/constants"
	"github.com/saigolang/AttYamlToJson/pkg/structs"
	"github.com/sirupsen/logrus"
	"net/http"
)

func YamlToJson(rawData []byte, logger *logrus.Logger) structs.Employees {
	var employees structs.Employees

	err := yaml.Unmarshal(rawData, &employees)
	if err != nil {
		// logging the error
		logger.WithFields(logrus.Fields{
			constants.RootCause:  err.Error(),
			constants.Trace:      constants.SystemError,
			constants.StatusCode: http.StatusInternalServerError,
		}).Error()

		return structs.Employees{
			Employees: nil,
			ErrorMessage: structs.ErrorContainer{
				RootCause:  err.Error(),
				StatusCode: http.StatusInternalServerError,
				Trace:      constants.SystemError,
			},
		}

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
