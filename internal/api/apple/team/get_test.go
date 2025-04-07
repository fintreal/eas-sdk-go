package team

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/graphql"
	"github.com/stretchr/testify/assert"
)

func TestGetByIdentifier(t *testing.T) {
	expectedData := &TeamData{
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

	mockData := &teamData{
		Id:         expectedData.Id,
		Name:       expectedData.Name,
		Identifier: expectedData.Identifier,
		Type:       expectedData.Type,
		Account:    account{Id: expectedData.AccountId},
	}
	mockResponse := getTeamResponse{GetTeam: getTeam{Data: mockData}}

	graphQLMock := graphql.NewGraphQLMock(mockResponse)

	service := NewTeamService(graphQLMock)

	actualData, actualErr := service.GetByIdentifier(expectedData.Identifier, expectedData.AccountId)

	actualQuery := graphQLMock.Calls[0].Arguments.Get(0).(string)
	actualVariables := graphQLMock.Calls[0].Arguments.Get(1).(map[string]any)

	assert.Equal(t, getQuery, actualQuery)
	assert.Equal(t, expectedVariables, actualVariables)
	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}
