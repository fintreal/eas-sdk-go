package app

import (
	"testing"

	"github.com/fintreal/expo-eas-sdk-go/internal/graphql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGet(t *testing.T) {
	expectedData := &AppData{
		Id:   "test-id",
		Name: "test-name",
		Slug: "test-slug",
	}

	var expectedVariables = map[string]any{
		"id": expectedData.Id,
	}

	graphQLMock := newGetGraphQLMock(expectedData)

	service := NewAppService(graphQLMock)

	acutalData, actualErr := service.Get(expectedData.Id)

	actualQuery := graphQLMock.Calls[0].Arguments.Get(0).(string)
	actualVariables := graphQLMock.Calls[0].Arguments.Get(1).(map[string]any)

	assert.Equal(t, nil, actualErr)
	assert.Equal(t, getAppQuery, actualQuery)
	assert.Equal(t, expectedVariables, actualVariables)
	assert.Equal(t, expectedData, acutalData)
}

func newGetGraphQLMock(data *AppData) *graphql.GraphQLMock {
	graphQLMock := graphql.NewGraphQLMock()
	graphQLMock.On("Query", mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		*args.Get(2).(*getAppResponse) = getAppResponse{Data: data}
	}).Return(nil)
	return graphQLMock
}
