package appvariable

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/graphql"
	"github.com/stretchr/testify/assert"
)

var expectedAppId = "test-app-id"

var expectedData = &AppVariableData{
	Id:           "test-id",
	Name:         "test-name",
	Value:        "test-value",
	Visibility:   "PUBLIC",
	Environments: []string{"DEVELOPMENT"},
}

var mockResponse = getAppVariablesResponse{AppByAppId: getAppVariables{Data: []AppVariableData{*expectedData}}}

func TestGetByName(t *testing.T) {
	expectedVariables := map[string]any{"appId": expectedAppId}

	graphQLMock := graphql.NewGraphQLMock(mockResponse)

	service := NewAppVariableService(graphQLMock)

	actualData, actualErr := service.GetByName(expectedData.Name, expectedAppId)

	actualQuery := graphQLMock.Calls[0].Arguments.Get(0).(string)
	actualVariables := graphQLMock.Calls[0].Arguments.Get(1).(map[string]any)

	assert.Equal(t, getAppVariableQuery, actualQuery)
	assert.Equal(t, expectedVariables, actualVariables)
	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}

func TestGetById(t *testing.T) {
	expectedVariables := map[string]any{"appId": expectedAppId}
	graphQLMock := graphql.NewGraphQLMock(mockResponse)

	appVariableService := NewAppVariableService(graphQLMock)

	actualData, actualErr := appVariableService.Get(expectedData.Id, expectedAppId)

	actualQuery := graphQLMock.Calls[0].Arguments.Get(0).(string)
	actualVariables := graphQLMock.Calls[0].Arguments.Get(1).(map[string]any)

	assert.Equal(t, getAppVariableQuery, actualQuery)
	assert.Equal(t, expectedVariables, actualVariables)
	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}
