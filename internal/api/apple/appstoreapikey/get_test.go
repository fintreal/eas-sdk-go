package appstoreapikey

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/testutils"
	"github.com/fintreal/eas-sdk-go/internal/utils"
)

func TestGetByIdentifier(t *testing.T) {
	getData := &GeyByIdentifierAppStoreApiKeyData{
		Identifier: "test-identifier",
		AccountId:  "test-account-id",
	}
	expectedData := &AppStoreApiKeyData{
		Id:               "test-id",
		Name:             "test-name",
		IssuerIdentifier: "test-issuer-identifier",
		Identifier:       getData.Identifier,
	}
	expectedVariables := map[string]any{"accountId": getData.AccountId}

	mockResponse := utils.AccountResponse[appStoreApiKeysResponse]{
		Account: utils.Account[appStoreApiKeysResponse]{
			ById: appStoreApiKeysResponse{
				Data: []AppStoreApiKeyData{*expectedData},
			},
		},
	}

	config := testutils.TestConfig[GeyByIdentifierAppStoreApiKeyData, AppStoreApiKeyData, utils.AccountResponse[appStoreApiKeysResponse], AppStoreApiKeyService]{
		NewServiceFunction: NewAppStoreApiKeyService,
		FunctionUnderTest:  "GetByIdentifier",
		Input:              getData,
		MockResponse:       mockResponse,
		ExpectedQuery:      getQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       expectedData,
	}
	testutils.Test(t, config)
}
