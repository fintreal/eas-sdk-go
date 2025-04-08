package me

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/testutils"
)

func TestGet(t *testing.T) {
	var expectedData = &Data{
		Id:          "test-id",
		DisplayName: "Test Display Name",
	}
	mockResponse := getResponse{Data: expectedData}
	config := testutils.TestConfig[any, Data, getResponse, Service]{
		NewServiceFunction: NewService,
		FunctionUnderTest:  "Get",
		Input:              nil,
		MockResponse:       mockResponse,
		ExpectedQuery:      getQuery,
		ExpectedVariables:  map[string]any{},
		ExpectedData:       expectedData,
	}
	testutils.Test(t, config)
}
