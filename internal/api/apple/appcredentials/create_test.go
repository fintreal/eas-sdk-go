package appcredentials

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/api/apple/appbuildcredentials"
	"github.com/fintreal/eas-sdk-go/internal/utils"
)

func TestCreate(t *testing.T) {
	expectedData := &Data{
		Id:               "test-id",
		AppIdentifierId:  "test-app-identifier-id",
		AppId:            "test-app-id",
		BuildCredentials: []appbuildcredentials.Data{},
	}

	expectedVariables := map[string]any{
		"appId":           expectedData.AppId,
		"appIdentifierId": expectedData.AppIdentifierId,
	}

	mockResponse := createResponse{
		IosAppCredentials: createIosAppCredentials{
			Data: data{
				Id: expectedData.Id,
				App: objWithId{
					Id: expectedData.AppId,
				},
				AppIdentifier: objWithId{
					Id: expectedData.AppIdentifierId,
				},
				BuildCredentials: []buildCredentials{},
			},
		},
	}

	input := &CreateData{
		AppId:           expectedData.AppId,
		AppIdentifierId: expectedData.AppIdentifierId,
	}

	config := utils.TestConfig[CreateData, Data, createResponse, Service]{
		NewServiceFunction: NewService,
		FunctionUnderTest:  "Create",
		Input:              input,
		MockResponse:       mockResponse,
		ExpectedQuery:      createQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       expectedData,
	}
	utils.Test(t, config)
}
