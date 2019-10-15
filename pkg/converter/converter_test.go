package converter

import (
	"AttYamlToJson/pkg/structs"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"testing"
)

func TestYamlToJson(t *testing.T) {

	t.Run("FailPath-YamlFileIsEmpty", func(t *testing.T) {
		path := filepath.Join("testdata", "empty.yaml")
		resp, err := ioutil.ReadFile(path)
		if err != nil {
			t.Error("error in converting yaml file to bytes")
		}

		result := YamlToJson(resp)
		assert.Len(t, result.Employees, 0)
		assert.Equal(t, "no data found", result.ErrorMessage.RootCause)
		assert.Equal(t, http.StatusNoContent, result.ErrorMessage.StatusCode)
	})

	t.Run("FailPath-YamlFileIsCorrupted", func(t *testing.T) {
		path := filepath.Join("testdata", "badfile.yaml")
		resp, err := ioutil.ReadFile(path)
		if err != nil {
			t.Error("error in converting yaml file to bytes")
		}

		result := YamlToJson(resp)
		assert.Len(t, result.Employees, 0)
		assert.Equal(t, "a system error has occurred", result.ErrorMessage.Trace)
		assert.Equal(t, http.StatusInternalServerError, result.ErrorMessage.StatusCode)
		assert.Equal(t, "error converting YAML to JSON: yaml: line 4: did not find expected '-' indicator", result.ErrorMessage.RootCause)
	})

	t.Run("HappyPath-ValidYamlFile", func(t *testing.T) {
		path := filepath.Join("testdata", "validdata.yaml")
		resp, err := ioutil.ReadFile(path)
		if err != nil {
			t.Error("error in converting yaml file to bytes")
		}

		result := YamlToJson(resp)
		assert.Len(t, result.Employees, 1)
		assert.Equal(t, structs.ErrorContainer{}, result.ErrorMessage)
	})
}
