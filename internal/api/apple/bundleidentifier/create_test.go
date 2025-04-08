package bundleidentifier

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/testutils"
)

func TestCreate(t *testing.T) {
	identifier := "test-identifier"
	accountId := "test-account-id"
	expectedData := &AppBundleIdentifierData{
		Id:         "test-id",
		Identifier: identifier,
		TeamId:     "test-team-id",
	}

	expectedVariables := map[string]any{
		"identifier": identifier,
		"accountId":  accountId,
		"teamId":     expectedData.TeamId,
	}

	mockResponse := createAppBundleIdentifierResponse{
		CreateAppIdentifier: createAppBundleIdentifier{
			Data: appBundleIdentifierData{
				Id:         expectedData.Id,
				Identifier: expectedData.Identifier,
				Team: team{
					Id: expectedData.TeamId,
				},
			},
		},
	}

	input := &CreateAppBundleIdentifierData{
		AccountId:  accountId,
		Identifier: identifier,
		TeamId:     expectedData.TeamId,
	}

	config := testutils.TestConfig[CreateAppBundleIdentifierData, AppBundleIdentifierData, createAppBundleIdentifierResponse, AppBundleIdentifierService]{
		NewServiceFunction: NewAppBundleIdentifierService,
		FunctionUnderTest:  "Create",
		Input:              input,
		MockResponse:       mockResponse,
		ExpectedQuery:      createQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       expectedData,
	}
	testutils.Test(t, config)
}
