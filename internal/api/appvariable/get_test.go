package appvariable

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/testutils"
)

var expectedAppId = "test-app-id"

var expectedData = &Data{
	Id:           "test-id",
	Name:         "test-name",
	Value:        "test-value",
	Visibility:   "PUBLIC",
	Environments: []string{"DEVELOPMENT"},
}

var mockResponse = getResponse{AppByAppId: appById{Data: []Data{*expectedData}}}

func TestGetByName(t *testing.T) {
	getData := &GetByNameData{
		AppId: expectedAppId,
		Name:  expectedData.Name,
	}
	expectedVariables := map[string]any{"appId": expectedAppId}

	config := testutils.TestConfig[GetByNameData, Data, getResponse, Service]{
		NewServiceFunction: NewService,
		FunctionUnderTest:  "GetByName",
		Input:              getData,
		MockResponse:       mockResponse,
		ExpectedQuery:      getQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       expectedData,
	}
	testutils.Test(t, config)
}

func TestGetById(t *testing.T) {
	getData := &GetData{
		AppId: expectedAppId,
		Id:    expectedData.Id,
	}
	expectedVariables := map[string]any{"appId": expectedAppId}

	config := testutils.TestConfig[GetData, Data, getResponse, Service]{
		NewServiceFunction: NewService,
		FunctionUnderTest:  "Get",
		Input:              getData,
		MockResponse:       mockResponse,
		ExpectedQuery:      getQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       expectedData,
	}
	testutils.Test(t, config)
}
