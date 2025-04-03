package appvariable

import (
	"testing"

	"github.com/fintreal/expo-eas-sdk-go/internal/graphql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

	graphQLMock := newCreateGraphQLMock(expectedData)

	service := NewAppVariableService(graphQLMock)

	actualData, actualErr := service.Create(inputData)

	actualQuery := graphQLMock.Calls[0].Arguments.Get(0).(string)
	actualVariables := graphQLMock.Calls[0].Arguments.Get(1).(map[string]any)

	assert.Equal(t, createAppVariableMutation, actualQuery)
	assert.Equal(t, expectedVariables, actualVariables)
	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}

func newCreateGraphQLMock(data *AppVariableData) *graphql.GraphQLMock {
	graphQLMock := graphql.NewGraphQLMock()
	graphQLMock.On("Query", mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		*args.Get(2).(*createAppVariableResponse) = createAppVariableResponse{CreateAppVariable: createAppVariable{Data: data}}
	}).Return(nil)
	return graphQLMock
}
