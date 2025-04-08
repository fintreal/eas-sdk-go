package provisioningprofile

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/testutils"
)

func TestCreate(t *testing.T) {
	accountId := "test-account-id"
	expectedData := &ProvisioningProfileData{
		Id:                    "test-id",
		AppBundleIdentifierId: "test-bundle-identifier-id",
		Base64:                "test-base64-string",
	}
	expectedVariables := map[string]any{
		"accountId":            accountId,
		"appleAppIdentifierId": expectedData.AppBundleIdentifierId,
		"base64":               expectedData.Base64,
	}

	input := &CreateProvisioningProfileData{
		AccountId:             accountId,
		AppBundleIdentifierId: expectedData.AppBundleIdentifierId,
		Base64:                expectedData.Base64,
	}

	mockResponse := createProvisioningProfileResponse{
		CreateProvisioningProfile: createProvisioningProfile{
			Data: provisioningProfileData{
				Id:                  expectedData.Id,
				Base64:              expectedData.Base64,
				AppBundleIdentifier: appleAppIdentifier{Id: expectedData.AppBundleIdentifierId},
			},
		},
	}

	config := testutils.TestConfig[CreateProvisioningProfileData, ProvisioningProfileData, createProvisioningProfileResponse, ProvisioningProfileService]{
		NewServiceFunction: NewProvisioningProfileService,
		FunctionUnderTest:  "Create",
		Input:              input,
		MockResponse:       mockResponse,
		ExpectedQuery:      createQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       expectedData,
	}
	testutils.Test(t, config)
}
