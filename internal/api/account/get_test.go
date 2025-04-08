package account

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/testutils"
)

func TestGetByName(t *testing.T) {
	expectedData := &AccountData{
		Id:   "test-account-id",
		Name: "test-account-name",
	}
	expectedVariables := map[string]any{"name": expectedData.Name}

	mockResponse := getAccountResponse{Account: getAccount{ByName: expectedData}}

	config := testutils.TestConfig[string, AccountData, getAccountResponse, AccountService]{
		NewServiceFunction: NewAccountService,
		FunctionUnderTest:  "GetByName",
		Input:              &expectedData.Name,
		MockResponse:       mockResponse,
		ExpectedQuery:      getAccountByNameQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       expectedData,
	}
	testutils.Test(t, config)
}
