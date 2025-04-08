package app

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/testutils"
)

func TestUpdate(t *testing.T) {
	expectedData := &Data{
		Id:   "test-id",
		Name: "test-name",
	}

	expectedVariables := map[string]any{
		"id":   expectedData.Id,
		"name": expectedData.Name,
	}

	input := &UpdateData{
		Id:   expectedData.Id,
		Name: expectedData.Name,
	}

	mockResponse := updateResponse{UpdateApp: updateApp{Data: expectedData}}
	config := testutils.TestConfig[UpdateData, Data, updateResponse, Service]{
		NewServiceFunction: NewService,
		FunctionUnderTest:  "Update",
		Input:              input,
		MockResponse:       mockResponse,
		ExpectedQuery:      updateQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       expectedData,
	}
	testutils.Test(t, config)
}
