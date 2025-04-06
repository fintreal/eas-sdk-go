package bundleidentifier

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/graphql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

	graphQLMock := newCreateGraphQLMock(expectedResponse)

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

func newCreateGraphQLMock(data *AppBundleIdentifierData) *graphql.GraphQLMock {
	mockResponse := createAppBundleIdentifierResponse{
		CreateAppIdentifier: createAppBundleIdentifier{
			Data: appBundleIdentifierData{
				Id:         data.Id,
				Identifier: data.Identifier,
				Team: team{
					Id: data.TeamId,
				},
			},
		},
	}

	graphQLMock := graphql.NewGraphQLMock()
	graphQLMock.On("Query", mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		*args.Get(2).(*createAppBundleIdentifierResponse) = mockResponse
	}).Return(nil)
	return graphQLMock
}
