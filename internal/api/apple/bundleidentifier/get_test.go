package bundleidentifier

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/testutils"
	"github.com/fintreal/eas-sdk-go/internal/utils"
)

func TestGetByIdentifier(t *testing.T) {
	identifier := "test-identifier"
	accountId := "test-account-id"
	input := &GetByIdentifierData{
		Identifier: identifier,
		AccountId:  accountId,
	}
	expectedData := &AppBundleIdentifierData{
		Id:         "test-id",
		Identifier: identifier,
		TeamId:     "test-team-id",
	}

	expectedVariables := map[string]any{
		"identifier": identifier,
		"accountId":  accountId,
	}
	mockResponse := getMockResponse(expectedData)

	config := testutils.TestConfig[GetByIdentifierData, AppBundleIdentifierData, utils.AccountResponse[appBundleIdentifierResponse], AppBundleIdentifierService]{
		NewServiceFunction: NewAppBundleIdentifierService,
		FunctionUnderTest:  "GetByIdentifier",
		Input:              input,
		MockResponse:       mockResponse,
		ExpectedQuery:      getQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       expectedData,
	}
	testutils.Test(t, config)
}

func getMockResponse(data *AppBundleIdentifierData) utils.AccountResponse[appBundleIdentifierResponse] {
	mockData := appBundleIdentifierData{
		Id:         data.Id,
		Identifier: data.Identifier,
		Team: team{
			Id: data.TeamId,
		},
	}
	return utils.AccountResponse[appBundleIdentifierResponse]{
		Account: utils.Account[appBundleIdentifierResponse]{
			ById: appBundleIdentifierResponse{
				Data: []appBundleIdentifierData{mockData},
			},
		},
	}
}
