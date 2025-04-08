package app

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/testutils"
)

func TestCreate(t *testing.T) {
	accountId := "test-account-id"
	expectedData := &AppData{
		Id:   "test-id",
		Name: "test-name",
		Slug: "test-slug",
	}
	expectedVariables := map[string]any{
		"accountId": accountId,
		"name":      expectedData.Name,
		"slug":      expectedData.Slug,
	}
	input := &CreateAppData{
		AccountId: accountId,
		Name:      expectedData.Name,
		Slug:      expectedData.Slug,
	}

	mockResponse := createAppResponse{CreateApp: createApp{Data: expectedData}}

	config := testutils.TestConfig[CreateAppData, AppData, createAppResponse, AppService]{
		NewServiceFunction: NewAppService,
		FunctionUnderTest:  "Create",
		Input:              input,
		MockResponse:       mockResponse,
		ExpectedQuery:      createAppMutation,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       expectedData,
	}
	testutils.Test(t, config)
}
