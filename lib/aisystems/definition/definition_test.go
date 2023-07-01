package definition_test

import (
	"encoding/json"
	"feather-ai/service-core/lib/aisystems/aisystemscore"
	"feather-ai/service-core/lib/aisystems/definition"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadDefinition(t *testing.T) {

	data := `{
	"steps": [
		{
			"inputs": [
				{
					"_name": "picker",
					"component_type": "FilePicker",
					"extensions": [
						"*.*"
					],
					"id": "40026d6e-cd0f-4baf-9067-f2307064f6f3",
					"max_files": 2,
					"schema": "{\"files\": [{\"data\": \"binary\", \"name\": \"string\"}, {\"data\": \"binary\", \"name\": \"string\"}]}",
					"type": "COMPONENT",
					"version": "1.0.0"
				},
				{
					"_name": "textbox",
					"component_type": "TextBox",
					"id": "bdc52b85-687f-47e4-9d8e-c7d0dd71d417",
					"num_lines": 3,
					"schema": "{\"text\": \"string\"}",
					"type": "COMPONENT",
					"version": "1.0.0"
				}
			],
			"name": "Step1",
			"outputs": []
		}
	],
	"version": "1.0.0"
}`

	rawDef, err := definition.Parse(data)
	assert.NoError(t, err)

	name, err := definition.GetStepNameByIndex(rawDef, 0)
	assert.NoError(t, err)
	assert.Equal(t, "Step1", name)

	r := aisystemscore.ExecuteSystemRequest{
		Id:        "paul",
		StepToRun: "hello",
		InputData: json.RawMessage("{ \"file1\":\"9uej9wejf09jfe09f9fef\"}"),
	}
	req, err := json.Marshal(r)
	reqstr := string(req)
	fmt.Printf(reqstr)
}
