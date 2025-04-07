package appvariable

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/graphql"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	expectedAppId := "test-app-id"
	expectedData := &AppVariableData{
		Id:           "test-id",
		Name:         "test-name",
		Value:        "test-value",
		Visibility:   "PUBLIC",
		Environments: []string{"PREVIEW", "PRODUCTION"},
	}

	expectedVariables := map[string]any{
		"appId":        expectedAppId,
		"name":         expectedData.Name,
		"value":        expectedData.Value,
		"environments": expectedData.Environments,
		"visibility":   expectedData.Visibility,
	}

	inputData := CreateAppVariableData{
		AppId:        expectedAppId,
		Name:         expectedData.Name,
		Value:        expectedData.Value,
		Visibility:   expectedData.Visibility,
		Environments: expectedData.Environments,
	}

	mockResponse := createAppVariableResponse{CreateAppVariable: createAppVariable{Data: expectedData}}
	graphQLMock := graphql.NewGraphQLMock(mockResponse)

	service := NewAppVariableService(graphQLMock)

	actualData, actualErr := service.Create(inputData)

	actualQuery := graphQLMock.Calls[0].Arguments.Get(0).(string)
	actualVariables := graphQLMock.Calls[0].Arguments.Get(1).(map[string]any)

	assert.Equal(t, createAppVariableMutation, actualQuery)
	assert.Equal(t, expectedVariables, actualVariables)
	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}
