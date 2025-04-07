package app

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/graphql"
	"github.com/stretchr/testify/assert"
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

	mockResponse := createAppResponse{CreateApp: createApp{Data: expectedData}}
	graphQLMock := graphql.NewGraphQLMock(mockResponse)

	service := NewAppService(graphQLMock)

	data, err := service.Create(inputData)

	actualQuery := graphQLMock.Calls[0].Arguments.Get(0).(string)
	actualVariables := graphQLMock.Calls[0].Arguments.Get(1).(map[string]any)

	assert.Equal(t, nil, err)
	assert.Equal(t, createAppMutation, actualQuery)
	assert.Equal(t, expectedVariables, actualVariables)
	assert.Equal(t, expectedData, data)
}
