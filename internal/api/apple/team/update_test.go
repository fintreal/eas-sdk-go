package team

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/graphql"
	"github.com/stretchr/testify/assert"
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

	mockResponse := &teamData{
		Id:         expectedData.Id,
		Name:       expectedData.Name,
		Identifier: expectedData.Identifier,
		Type:       expectedData.Type,
		Account:    account{Id: expectedData.AccountId},
	}
	graphQLMock := graphql.NewGraphQLMock(mockResponse)

	service := NewTeamService(graphQLMock)

	actualData, actualErr := service.Update(inputData)

	actualQuery := graphQLMock.Calls[0].Arguments.Get(0).(string)
	actualVariables := graphQLMock.Calls[0].Arguments.Get(1).(map[string]any)

	assert.Equal(t, createQuery, actualQuery)
	assert.Equal(t, expectedVariables, actualVariables)
	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}
