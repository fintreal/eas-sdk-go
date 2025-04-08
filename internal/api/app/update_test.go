package app

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/testutils"
)

func TestUpdate(t *testing.T) {
	expectedData := &AppData{
		Id:   "test-id",
		Name: "test-name",
	}

	expectedVariables := map[string]any{
		"id":   expectedData.Id,
		"name": expectedData.Name,
	}

	input := &UpdateAppData{
		Id:   expectedData.Id,
		Name: expectedData.Name,
	}

	mockResponse := updateAppResponse{UpdateApp: updateApp{Data: expectedData}}
	config := testutils.TestConfig[UpdateAppData, AppData, updateAppResponse, AppService]{
		NewServiceFunction: NewAppService,
		FunctionUnderTest:  "Update",
		Input:              input,
		MockResponse:       mockResponse,
		ExpectedQuery:      updateAppMutation,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       expectedData,
	}
	testutils.Test(t, config)
}
