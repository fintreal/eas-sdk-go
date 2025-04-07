package bundleidentifier

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/graphql"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
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
		"teamId":     expectedResponse.TeamId,
	}

	mockResponse := createAppBundleIdentifierResponse{
		CreateAppIdentifier: createAppBundleIdentifier{
			Data: appBundleIdentifierData{
				Id:         expectedResponse.Id,
				Identifier: expectedResponse.Identifier,
				Team: team{
					Id: expectedResponse.TeamId,
				},
			},
		},
	}

	graphQLMock := graphql.NewGraphQLMock(mockResponse)

	service := NewAppBundleIdentifierService(graphQLMock)

	result, err := service.Create(CreateAppBundleIdentifierData{
		AccountId:  accountId,
		Identifier: identifier,
		TeamId:     expectedResponse.TeamId,
	})

	actualQuery := graphQLMock.Calls[0].Arguments.Get(0).(string)
	actualVariables := graphQLMock.Calls[0].Arguments.Get(1).(map[string]any)

	assert.Equal(t, createQuery, actualQuery)
	assert.Equal(t, expectedVariables, actualVariables)
	assert.Equal(t, expectedResponse, result)
	assert.Equal(t, nil, err)

}
