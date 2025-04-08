package me

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/testutils"
)

func TestGet(t *testing.T) {
	var expectedData = &MeData{
		Id:          "test-id",
		DisplayName: "Test Display Name",
	}
	mockResponse := getMeResponse{Data: expectedData}
	config := testutils.TestConfig[any, MeData, getMeResponse, MeService]{
		NewServiceFunction: NewMeService,
		FunctionUnderTest:  "Get",
		Input:              nil,
		MockResponse:       mockResponse,
		ExpectedQuery:      meQuery,
		ExpectedVariables:  map[string]any{},
		ExpectedData:       expectedData,
	}
	testutils.Test(t, config)
}
