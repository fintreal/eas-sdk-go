package provisioningprofile

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/testutils"
	"github.com/fintreal/eas-sdk-go/internal/utils"
)

func TestGet(t *testing.T) {
	input := &GetProvisioningProfileData{
		Id:        "test-id",
		AccountId: "test-account-id",
	}

	identifier := "test-identifier"

	expectedData := &ProvisioningProfileData{
		Id:                    input.Id,
		AppBundleIdentifierId: identifier,
		Base64:                "test-base64-string",
	}

	expectedVariables := map[string]any{"accountId": input.AccountId}

	mockResponse := getMockResponse(expectedData)

	config := testutils.TestConfig[GetProvisioningProfileData, ProvisioningProfileData, utils.AccountResponse[getProvisioningProfilesResponse], ProvisioningProfileService]{
		NewServiceFunction: NewProvisioningProfileService,
		FunctionUnderTest:  "Get",
		Input:              input,
		MockResponse:       mockResponse,
		ExpectedQuery:      getQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       expectedData,
	}
	testutils.Test(t, config)
}

func getMockResponse(data *ProvisioningProfileData) utils.AccountResponse[getProvisioningProfilesResponse] {
	mockData := provisioningProfileData{
		Id:                  data.Id,
		Base64:              data.Base64,
		AppBundleIdentifier: appleAppIdentifier{Id: data.AppBundleIdentifierId},
	}

	return utils.AccountResponse[getProvisioningProfilesResponse]{
		Account: utils.Account[getProvisioningProfilesResponse]{
			ById: getProvisioningProfilesResponse{
				Data: []provisioningProfileData{mockData},
			},
		},
	}
}
