package appcredentials

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/api/android/appbuildcredentials"
	"github.com/fintreal/eas-sdk-go/internal/utils"
)

func TestGetByIdentifier(t *testing.T) {
	input := &GetData{
		Id:    "test-id",
		AppId: "test-app-id",
	}

	buildCredential := appbuildcredentials.Data{
		Id:               "test-build-credentials-id",
		Name:             "test-build-credentials-name",
		AppCredentialsId: input.Id,
		KeystoreId:       "test-keystore-id",
	}

	expectedData := &Data{
		Id:                        input.Id,
		AppId:                     input.AppId,
		Identifier:                "test-identifier",
		GoogleServiceAccountKeyId: "test-google-service-account-key-id",
		BuildCredentials:          []appbuildcredentials.Data{buildCredential},
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
				BuildCredentials: []buildCredentialsData{{
					Id:   buildCredential.Id,
					Name: buildCredential.Name,
					Keystore: objWithId{
						Id: buildCredential.KeystoreId,
					},
				}},
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
