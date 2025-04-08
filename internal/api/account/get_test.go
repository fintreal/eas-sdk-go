package account

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/testutils"
)

func TestGetByName(t *testing.T) {
	expectedData := &Data{
		Id:   "test-account-id",
		Name: "test-account-name",
	}
	expectedVariables := map[string]any{"name": expectedData.Name}

	mockResponse := getResponse{Account: account{ByName: expectedData}}

	config := testutils.TestConfig[string, Data, getResponse, AccountService]{
		NewServiceFunction: NewAccountService,
		FunctionUnderTest:  "GetByName",
		Input:              &expectedData.Name,
		MockResponse:       mockResponse,
		ExpectedQuery:      getByNameQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       expectedData,
	}
	testutils.Test(t, config)
}
