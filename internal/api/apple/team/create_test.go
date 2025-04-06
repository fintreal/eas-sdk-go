package team

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/graphql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	expectedData := &TeamData{
		Id:         "test-id",
		Name:       "test-name",
		Identifier: "test-identifier",
		Type:       "test-type",
		AccountId:  "test-account-id",
	}

	expectedVariables := map[string]any{
		"identifier": expectedData.Identifier,
		"name":       expectedData.Name,
		"accountId":  expectedData.AccountId,
		"type":       expectedData.Type,
	}

	inputData := CreateTeamData{
		Name:       expectedData.Name,
		Identifier: expectedData.Identifier,
		Type:       expectedData.Type,
		AccountId:  expectedData.AccountId,
	}

	graphQLMock := newCreateGraphQLMock(expectedData)

	service := NewTeamService(graphQLMock)

	actualData, actualErr := service.Create(inputData)

	actualQuery := graphQLMock.Calls[0].Arguments.Get(0).(string)
	actualVariables := graphQLMock.Calls[0].Arguments.Get(1).(map[string]any)

	assert.Equal(t, createQuery, actualQuery)
	assert.Equal(t, expectedVariables, actualVariables)
	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}

func newCreateGraphQLMock(data *TeamData) *graphql.GraphQLMock {
	mockData := &teamData{
		Id:         data.Id,
		Name:       data.Name,
		Identifier: data.Identifier,
		Type:       data.Type,
		Account:    account{Id: data.AccountId},
	}
	graphQLMock := graphql.NewGraphQLMock()
	graphQLMock.On("Query", mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		*args.Get(2).(*createTeamResponse) = createTeamResponse{CreateTeam: createTeam{Data: mockData}}
	}).Return(nil)
	return graphQLMock
}
