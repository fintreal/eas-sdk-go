package app

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/graphql"
	"github.com/stretchr/testify/assert"
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

	mockResponse := getAppResponse{Data: expectedData}
	graphQLMock := graphql.NewGraphQLMock(mockResponse)

	service := NewAppService(graphQLMock)

	acutalData, actualErr := service.Get(expectedData.Id)

	actualQuery := graphQLMock.Calls[0].Arguments.Get(0).(string)
	actualVariables := graphQLMock.Calls[0].Arguments.Get(1).(map[string]any)

	assert.Equal(t, nil, actualErr)
	assert.Equal(t, getAppQuery, actualQuery)
	assert.Equal(t, expectedVariables, actualVariables)
	assert.Equal(t, expectedData, acutalData)
}
