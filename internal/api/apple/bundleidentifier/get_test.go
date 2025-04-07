package bundleidentifier

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/graphql"
	"github.com/fintreal/eas-sdk-go/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetByIdentifier(t *testing.T) {
	identifier := "test-identifier"
	accountId := "test-account-id"
	expectedResponse := &AppBundleIdentifierData{
		Id:         "test-id",
		Identifier: identifier,
		TeamId:     "test-team-id",
	}

	expectedVariables := map[string]any{
		"identifier": identifier,
		"accountId":  accountId,
	}
	graphQLMock := graphql.NewGraphQLMock(mockResponse(expectedResponse))

	service := NewAppBundleIdentifierService(graphQLMock)

	result, err := service.GetByIdentifier(identifier, accountId)

	actualQuery := graphQLMock.Calls[0].Arguments.Get(0).(string)
	actualVariables := graphQLMock.Calls[0].Arguments.Get(1).(map[string]any)

	assert.Equal(t, getQuery, actualQuery)
	assert.Equal(t, expectedVariables, actualVariables)
	assert.Equal(t, expectedResponse, result)
	assert.Equal(t, nil, err)
}

func mockResponse(data *AppBundleIdentifierData) utils.AccountResponse[appBundleIdentifierResponse] {
	mockData := appBundleIdentifierData{
		Id:         data.Id,
		Identifier: data.Identifier,
		Team: team{
			Id: data.TeamId,
		},
	}
	return utils.AccountResponse[appBundleIdentifierResponse]{
		Account: utils.Account[appBundleIdentifierResponse]{
			ById: appBundleIdentifierResponse{
				Data: []appBundleIdentifierData{mockData},
			},
		},
	}
}
