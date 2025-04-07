package provisioningprofile

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/graphql"
	"github.com/fintreal/eas-sdk-go/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	accountId := "test-account-id"
	identifier := "test-identifier"

	expectedResponse := &ProvisioningProfileData{
		Id:                    "test-id",
		AppBundleIdentifierId: identifier,
		Base64:                "test-base64-string",
	}

	expectedVariables := map[string]any{"accountId": accountId}

	graphQLMock := graphql.NewGraphQLMock(mockResponse(expectedResponse))

	service := NewProvisioningProfileService(graphQLMock)
	actualResponse, actualErr := service.Get(expectedResponse.Id, accountId)

	actualQuery := graphQLMock.Calls[0].Arguments.Get(0).(string)
	actualVariables := graphQLMock.Calls[0].Arguments.Get(1).(map[string]any)

	assert.Equal(t, getQuery, actualQuery)
	assert.Equal(t, expectedVariables, actualVariables)
	assert.Equal(t, expectedResponse, actualResponse)
	assert.Equal(t, nil, actualErr)
}

func mockResponse(data *ProvisioningProfileData) utils.AccountResponse[getProvisioningProfilesResponse] {
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
