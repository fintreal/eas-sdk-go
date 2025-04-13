package appcredentials

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/utils"
)

func TestGetByIdentifier(t *testing.T) {
	input := &GetData{
		Id:    "test-id",
		AppId: "test-app-id",
	}

	expectedData := &Data{
		Id:                        input.Id,
		AppId:                     input.AppId,
		Identifier:                "test-identifier",
		GoogleServiceAccountKeyId: "test-google-service-account-key-id",
	}

	expectedVariables := map[string]any{
		"appId": input.AppId,
	}
	mockResponse := getResponse{
		AppByAppId: appByAppId{
			Data: []data{{
				Id:         expectedData.Id,
				Identifier: expectedData.Identifier,
				App: objWithId{
					Id: expectedData.AppId,
				},
				GoogleServiceAccountKey: objWithId{
					Id: expectedData.GoogleServiceAccountKeyId,
				},
			}},
		},
	}

	config := utils.TestConfig[GetData, Data, getResponse, Service]{
		NewServiceFunction: NewService,
		FunctionUnderTest:  "Get",
		Input:              input,
		MockResponse:       mockResponse,
		ExpectedQuery:      getQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       expectedData,
	}
	utils.Test(t, config)
}
