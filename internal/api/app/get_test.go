package app

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/utils"
)

func TestGet(t *testing.T) {
	expectedData := &Data{
		Id:   "test-id",
		Name: "test-name",
		Slug: "test-slug",
	}

	config := utils.TestConfig[string, Data, getResponse, Service]{
		NewServiceFunction: NewService,
		FunctionUnderTest:  "Get",
		Input:              &expectedData.Id,
		MockResponse:       getResponse{Data: expectedData},
		ExpectedQuery:      getQuery,
		ExpectedVariables:  map[string]any{"id": expectedData.Id},
		ExpectedData:       expectedData,
	}
	utils.Test(t, config)
}
