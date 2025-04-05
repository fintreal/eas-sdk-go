package provisioningprofile

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/graphql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

	graphQLMock := newGetGraphQLMock(expectedResponse)
	service := NewProvisioningProfileService(graphQLMock)
	actualResponse, actualErr := service.Get(expectedResponse.Id, accountId)

	actualQuery := graphQLMock.Calls[0].Arguments.Get(0).(string)
	actualVariables := graphQLMock.Calls[0].Arguments.Get(1).(map[string]any)

	assert.Equal(t, getQuery, actualQuery)
	assert.Equal(t, expectedVariables, actualVariables)
	assert.Equal(t, expectedResponse, actualResponse)
	assert.Equal(t, nil, actualErr)
}

func newGetGraphQLMock(data *ProvisioningProfileData) *graphql.GraphQLMock {
	mockResponse := getProvisioningProfilesResponse{
		Account: getProvisioningProfiles{
			ById: byId{
				Data: []provisioningProfileData{
					{
						Id:                  data.Id,
						Base64:              data.Base64,
						AppBundleIdentifier: appleAppIdentifier{Id: data.AppBundleIdentifierId},
					},
				},
			},
		},
	}

	graphQLMock := graphql.NewGraphQLMock()
	graphQLMock.On("Query", mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		response := args.Get(2).(*getProvisioningProfilesResponse)
		*response = mockResponse
	}).Return(nil)

	return graphQLMock
}
