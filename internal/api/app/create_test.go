package app

import (
	"testing"

	"github.com/fintreal/expo-eas-sdk-go/internal/graphql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	accountId := "test-account-id"
	expectedData := &AppData{
		Id:   "test-id",
		Name: "test-name",
		Slug: "test-slug",
	}

	expectedVariables := map[string]any{
		"accountId": accountId,
		"name":      expectedData.Name,
		"slug":      expectedData.Slug,
	}

	inputData := CreateAppData{
		AccountId: accountId,
		Name:      expectedData.Name,
		Slug:      expectedData.Slug,
	}

	graphQLMock := newCreateGraphQLMock(expectedData)

	service := NewAppService(graphQLMock)

	data, err := service.Create(inputData)

	actualQuery := graphQLMock.Calls[0].Arguments.Get(0).(string)
	actualVariables := graphQLMock.Calls[0].Arguments.Get(1).(map[string]any)

	assert.Equal(t, nil, err)
	assert.Equal(t, createAppMutation, actualQuery)
	assert.Equal(t, expectedVariables, actualVariables)
	assert.Equal(t, expectedData, data)
}

func newCreateGraphQLMock(data *AppData) *graphql.GraphQLMock {
	graphQLMock := graphql.NewGraphQLMock()
	graphQLMock.On("Query", mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		*args.Get(2).(*createAppResponse) = createAppResponse{CreateApp: createApp{Data: data}}
	}).Return(nil)
	return graphQLMock
}
