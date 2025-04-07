package appstoreapikey

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/graphql"
	"github.com/fintreal/eas-sdk-go/internal/utils"
	"github.com/stretchr/testify/assert"
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

	mockResponse := utils.AccountResponse[appStoreApiKeysResponse]{
		Account: utils.Account[appStoreApiKeysResponse]{
			ById: appStoreApiKeysResponse{
				Data: []AppStoreApiKeyData{*expectedResponse},
			},
		},
	}

	graphQLMock := graphql.NewGraphQLMock(mockResponse)
	service := NewAppStoreApiKeyService(graphQLMock)

	actualResponse, actualErr := service.GetByIdentifier(identifier, accountId)

	actualQuery := graphQLMock.Calls[0].Arguments.Get(0).(string)
	actualVariables := graphQLMock.Calls[0].Arguments.Get(1).(map[string]any)

	assert.Equal(t, getQuery, actualQuery)
	assert.Equal(t, expectedVariables, actualVariables)
	assert.Equal(t, expectedResponse, actualResponse)
	assert.Equal(t, nil, actualErr)
}
