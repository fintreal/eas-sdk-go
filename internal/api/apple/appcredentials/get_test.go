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
		Id:              input.Id,
		AppIdentifierId: "test-app-identifier-id",
		AppId:           input.AppId,
	}

	expectedVariables := map[string]any{
		"id":    input.Id,
		"appId": input.AppId,
	}
	mockResponse := getResponse{
		IosAppCredentials: GetIosAppCredentials{
			Data: data{
				Id: expectedData.Id,
				App: objWithId{
					Id: expectedData.AppId,
				},
				AppIdentifier: objWithId{
					Id: expectedData.AppIdentifierId,
				},
			},
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
