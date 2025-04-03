package appvariable

import (
	"testing"

	"github.com/fintreal/expo-eas-sdk-go/internal/graphql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var expectedAppId = "test-app-id"

var expectedData = &AppVariableData{
	Id:           "test-id",
	Name:         "test-name",
	Value:        "test-value",
	Visibility:   "PUBLIC",
	Environments: []string{"DEVELOPMENT"},
}

func TestGetByName(t *testing.T) {
	expectedVariables := map[string]any{"appId": expectedAppId}

	graphQLMock := newGetGraphQLMock(expectedData)

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
	graphQLMock := newGetGraphQLMock(expectedData)

	appVariableService := NewAppVariableService(graphQLMock)

	actualData, actualErr := appVariableService.Get(expectedData.Id, expectedAppId)

	actualQuery := graphQLMock.Calls[0].Arguments.Get(0).(string)
	actualVariables := graphQLMock.Calls[0].Arguments.Get(1).(map[string]any)

	assert.Equal(t, getAppVariableQuery, actualQuery)
	assert.Equal(t, expectedVariables, actualVariables)
	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)

}

func newGetGraphQLMock(data *AppVariableData) *graphql.GraphQLMock {
	graphQLMock := graphql.NewGraphQLMock()
	graphQLMock.On("Query", mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		*args.Get(2).(*getAppVariablesResponse) = getAppVariablesResponse{
			AppByAppId: getAppVariables{
				Data: []AppVariableData{*data},
			},
		}
	}).Return(nil)
	return graphQLMock
}
