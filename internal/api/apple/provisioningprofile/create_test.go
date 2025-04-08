package provisioningprofile

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/testutils"
)

func TestCreate(t *testing.T) {
	accountId := "test-account-id"
	expectedData := &Data{
		Id:                    "test-id",
		AppBundleIdentifierId: "test-bundle-identifier-id",
		Base64:                "test-base64-string",
	}
	expectedVariables := map[string]any{
		"accountId":            accountId,
		"appleAppIdentifierId": expectedData.AppBundleIdentifierId,
		"base64":               expectedData.Base64,
	}

	input := &CreateData{
		AccountId:             accountId,
		AppBundleIdentifierId: expectedData.AppBundleIdentifierId,
		Base64:                expectedData.Base64,
	}

	mockResponse := createResponse{
		AppleProvisioningProfile: appleProvisioningProfile{
			Data: data{
				Id:                 expectedData.Id,
				Base64:             expectedData.Base64,
				AppleAppIdentifier: appleAppIdentifier{Id: expectedData.AppBundleIdentifierId},
			},
		},
	}

	config := testutils.TestConfig[CreateData, Data, createResponse, Service]{
		NewServiceFunction: NewService,
		FunctionUnderTest:  "Create",
		Input:              input,
		MockResponse:       mockResponse,
		ExpectedQuery:      createQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       expectedData,
	}
	testutils.Test(t, config)
}
