package app

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/graphql"
	"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {
	expectedData := &AppData{
		Id:   "test-id",
		Name: "test-name",
	}

	expectedVariables := map[string]any{
		"id":   expectedData.Id,
		"name": expectedData.Name,
	}

	inputData := UpdateAppData{
		Id:   expectedData.Id,
		Name: expectedData.Name,
	}

	mockResponse := updateAppResponse{UpdateApp: updateApp{Data: expectedData}}
	graphQLMock := graphql.NewGraphQLMock(mockResponse)

	service := NewAppService(graphQLMock)

	data, err := service.Update(inputData)

	actualQuery := graphQLMock.Calls[0].Arguments.Get(0).(string)
	actualVariables := graphQLMock.Calls[0].Arguments.Get(1).(map[string]any)

	assert.Equal(t, nil, err)
	assert.Equal(t, updateAppMutation, actualQuery)
	assert.Equal(t, expectedVariables, actualVariables)
	assert.Equal(t, expectedData, data)
}
