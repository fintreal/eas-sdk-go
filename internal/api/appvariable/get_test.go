package appvariable

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/testutils"
)

var expectedAppId = "test-app-id"

var expectedData = &AppVariableData{
	Id:           "test-id",
	Name:         "test-name",
	Value:        "test-value",
	Visibility:   "PUBLIC",
	Environments: []string{"DEVELOPMENT"},
}

var mockResponse = getAppVariablesResponse{AppByAppId: getAppVariables{Data: []AppVariableData{*expectedData}}}

func TestGetByName(t *testing.T) {
	getData := &GetByNameAppVariableData{
		AppId: expectedAppId,
		Name:  expectedData.Name,
	}
	expectedVariables := map[string]any{"appId": expectedAppId}

	config := testutils.TestConfig[GetByNameAppVariableData, AppVariableData, getAppVariablesResponse, AppVariableService]{
		NewServiceFunction: NewAppVariableService,
		FunctionUnderTest:  "GetByName",
		Input:              getData,
		MockResponse:       mockResponse,
		ExpectedQuery:      getAppVariablesQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       expectedData,
	}
	testutils.Test(t, config)
}

func TestGetById(t *testing.T) {
	getData := &GetAppVariableData{
		AppId: expectedAppId,
		Id:    expectedData.Id,
	}
	expectedVariables := map[string]any{"appId": expectedAppId}

	config := testutils.TestConfig[GetAppVariableData, AppVariableData, getAppVariablesResponse, AppVariableService]{
		NewServiceFunction: NewAppVariableService,
		FunctionUnderTest:  "Get",
		Input:              getData,
		MockResponse:       mockResponse,
		ExpectedQuery:      getAppVariablesQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       expectedData,
	}
	testutils.Test(t, config)
}
