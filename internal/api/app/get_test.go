package app

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/testutils"
)

func TestGet(t *testing.T) {
	expectedData := &AppData{
		Id:   "test-id",
		Name: "test-name",
		Slug: "test-slug",
	}

	config := testutils.TestConfig[string, AppData, getAppResponse, AppService]{
		NewServiceFunction: NewAppService,
		FunctionUnderTest:  "Get",
		Input:              &expectedData.Id,
		MockResponse:       getAppResponse{Data: expectedData},
		ExpectedQuery:      getAppQuery,
		ExpectedVariables:  map[string]any{"id": expectedData.Id},
		ExpectedData:       expectedData,
	}
	testutils.Test(t, config)
}
