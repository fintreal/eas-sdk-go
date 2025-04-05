package provisioningprofile

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/graphql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	accountId := "test-account-id"
	expectedResponse := &ProvisioningProfileData{
		Id:                    "test-id",
		AppBundleIdentifierId: "test-bundle-identifier-id",
		Base64:                "test-base64-string",
	}
	expectedVariables := map[string]any{
		"accountId":            accountId,
		"appleAppIdentifierId": expectedResponse.AppBundleIdentifierId,
		"base64":               expectedResponse.Base64,
	}

	input := CreateProvisioningProfileData{
		AccountId:             accountId,
		AppBundleIdentifierId: expectedResponse.AppBundleIdentifierId,
		Base64:                expectedResponse.Base64,
	}

	graphQLMock := newCreateGraphQLMock(expectedResponse)
	service := NewProvisioningProfileService(graphQLMock)
	actualResponse, actualErr := service.Create(input)

	actualQuery := graphQLMock.Calls[0].Arguments.Get(0).(string)
	actualVariables := graphQLMock.Calls[0].Arguments.Get(1).(map[string]any)

	assert.Equal(t, createQuery, actualQuery)
	assert.Equal(t, expectedVariables, actualVariables)
	assert.Equal(t, expectedResponse, actualResponse)
	assert.Equal(t, nil, actualErr)
}

func newCreateGraphQLMock(data *ProvisioningProfileData) *graphql.GraphQLMock {
	mockResponse := createProvisioningProfileResponse{
		CreateProvisioningProfile: createProvisioningProfile{
			Data: provisioningProfileData{
				Id:                  data.Id,
				Base64:              data.Base64,
				AppBundleIdentifier: appleAppIdentifier{Id: data.AppBundleIdentifierId},
			},
		},
	}

	graphQLMock := graphql.NewGraphQLMock()
	graphQLMock.On("Query", mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		response := args.Get(2).(*createProvisioningProfileResponse)
		*response = mockResponse
	}).Return(nil)

	return graphQLMock
}
