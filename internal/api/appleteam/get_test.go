package appleteam

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/graphql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetByIdentifier(t *testing.T) {
	expectedData := &AppleTeamData{
		Id:         "test-id",
		Name:       "test-name",
		Identifier: "test-identifier",
		Type:       "test-type",
		AccountId:  "test-account-id",
	}

	expectedVariables := map[string]any{
		"identifier": expectedData.Identifier,
		"accountId":  expectedData.AccountId,
	}

	graphQLMock := newGetGraphQLMock(expectedData)

	service := NewAppleTeamService(graphQLMock)

	actualData, actualErr := service.GetByIdentifier(expectedData.Identifier, expectedData.AccountId)

	actualQuery := graphQLMock.Calls[0].Arguments.Get(0).(string)
	actualVariables := graphQLMock.Calls[0].Arguments.Get(1).(map[string]any)

	assert.Equal(t, getAppleTeamByIdentifierQuery, actualQuery)
	assert.Equal(t, expectedVariables, actualVariables)
	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}

func newGetGraphQLMock(data *AppleTeamData) *graphql.GraphQLMock {
	mockData := &appleTeamData{
		Id:         data.Id,
		Name:       data.Name,
		Identifier: data.Identifier,
		Type:       data.Type,
		Account:    account{Id: data.AccountId},
	}
	graphQLMock := graphql.NewGraphQLMock()
	graphQLMock.On("Query", mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		*args.Get(2).(*getAppleTeamResponse) = getAppleTeamResponse{GetAppleTeam: getAppleTeam{Data: mockData}}
	}).Return(nil)
	return graphQLMock
}
