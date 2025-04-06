package team

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/graphql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func UpdateCreate(t *testing.T) {
	expectedData := &TeamData{
		Id:         "test-id",
		Name:       "test-name",
		Identifier: "test-identifier",
		Type:       "test-type",
		AccountId:  "test-account-id",
	}

	expectedVariables := map[string]any{
		"id":   expectedData.Id,
		"name": expectedData.Name,
		"type": expectedData.Type,
	}

	inputData := UpdateTeamData{
		Id:   expectedData.Id,
		Name: expectedData.Name,
		Type: expectedData.Type,
	}

	graphQLMock := newUpdateGraphQLMock(expectedData)

	service := NewTeamService(graphQLMock)

	actualData, actualErr := service.Update(inputData)

	actualQuery := graphQLMock.Calls[0].Arguments.Get(0).(string)
	actualVariables := graphQLMock.Calls[0].Arguments.Get(1).(map[string]any)

	assert.Equal(t, createQuery, actualQuery)
	assert.Equal(t, expectedVariables, actualVariables)
	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}

func newUpdateGraphQLMock(data *TeamData) *graphql.GraphQLMock {
	mockData := &teamData{
		Id:         data.Id,
		Name:       data.Name,
		Identifier: data.Identifier,
		Type:       data.Type,
		Account:    account{Id: data.AccountId},
	}
	graphQLMock := graphql.NewGraphQLMock()
	graphQLMock.On("Query", mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		*args.Get(2).(*updateTeamResponse) = updateTeamResponse{UpdateAppleTeam: updateTeam{Data: mockData}}
	}).Return(nil)
	return graphQLMock
}
