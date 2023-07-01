package definition

import (
	"encoding/json"
	"errors"
)

var gErrWrongVersion = errors.New("Don't know how to parse system definition version")
var gErrBadStep = errors.New("Bad step value")

type RawDefinition map[string]interface{}

func Parse(definitionJson string) (RawDefinition, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(definitionJson), &data)
	if err != nil {
		return nil, err
	}

	if data["version"] != "1.0.0" {
		return nil, gErrWrongVersion
	}
	return RawDefinition(data), nil
}

func (d RawDefinition) NumSteps() int {
	steps, ok := d["steps"]
	if ok == false {
		return 0
	}
	return len(steps.([]interface{}))
}

func GetStepNameByIndex(definition RawDefinition, stepIndex int) (string, error) {
	steps := definition["steps"].([]interface{})
	if stepIndex >= len(steps) {
		return "", gErrBadStep
	}

	step := steps[stepIndex].(map[string]interface{})
	return step["name"].(string), nil
}
