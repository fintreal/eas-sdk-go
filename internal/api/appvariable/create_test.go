package appvariable

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/testutils"
)

func TestCreate(t *testing.T) {
	expectedAppId := "test-app-id"
	expectedData := &AppVariableData{
		Id:           "test-id",
		Name:         "test-name",
		Value:        "test-value",
		Visibility:   "PUBLIC",
		Environments: []string{"PREVIEW", "PRODUCTION"},
	}

	expectedVariables := map[string]any{
		"appId":        expectedAppId,
		"name":         expectedData.Name,
		"value":        expectedData.Value,
		"environments": expectedData.Environments,
		"visibility":   expectedData.Visibility,
	}

	input := &CreateAppVariableData{
		AppId:        expectedAppId,
		Name:         expectedData.Name,
		Value:        expectedData.Value,
		Visibility:   expectedData.Visibility,
		Environments: expectedData.Environments,
	}

	mockResponse := createAppVariableResponse{CreateAppVariable: createAppVariable{Data: expectedData}}

	config := testutils.TestConfig[CreateAppVariableData, AppVariableData, createAppVariableResponse, AppVariableService]{
		NewServiceFunction: NewAppVariableService,
		FunctionUnderTest:  "Create",
		Input:              input,
		MockResponse:       mockResponse,
		ExpectedQuery:      createAppVariableMutation,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       expectedData,
	}
	testutils.Test(t, config)
}
