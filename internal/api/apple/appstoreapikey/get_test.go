package appstoreapikey

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/graphql"
	"github.com/fintreal/eas-sdk-go/internal/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetByIdentifier(t *testing.T) {
	identifier := "test-identifier"
	accountId := "test-account-id"
	expectedResponse := &AppStoreApiKeyData{
		Id:               "test-id",
		Name:             "test-name",
		IssuerIdentifier: "test-issuer-identifier",
		Identifier:       identifier,
	}
	expectedVariables := map[string]any{"accountId": accountId}

	graphQLMock := newGetGraphQLMock(expectedResponse)
	service := NewAppStoreApiKeyService(graphQLMock)

	actualResponse, actualErr := service.GetByIdentifier(identifier, accountId)

	actualQuery := graphQLMock.Calls[0].Arguments.Get(0).(string)
	actualVariables := graphQLMock.Calls[0].Arguments.Get(1).(map[string]any)

	assert.Equal(t, getQuery, actualQuery)
	assert.Equal(t, expectedVariables, actualVariables)
	assert.Equal(t, expectedResponse, actualResponse)
	assert.Equal(t, nil, actualErr)
}

func newGetGraphQLMock(data *AppStoreApiKeyData) *graphql.GraphQLMock {
	mockResponse := utils.AccountResponse[appStoreApiKeysResponse]{
		Account: utils.Account[appStoreApiKeysResponse]{
			ById: appStoreApiKeysResponse{
				Data: []AppStoreApiKeyData{*data},
			},
		},
	}

	graphQLMock := graphql.NewGraphQLMock()
	graphQLMock.On("Query", mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		*args.Get(2).(*utils.AccountResponse[appStoreApiKeysResponse]) = mockResponse
	}).Return(nil)
	return graphQLMock
}
