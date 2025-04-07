package me

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/graphql"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	var expectedData = &MeData{
		Id:          "test-id",
		DisplayName: "Test Display Name",
	}
	mockResponse := getMeResponse{Data: expectedData}
	graphQLMock := graphql.NewGraphQLMock(mockResponse)

	service := NewMeService(graphQLMock)

	actualData, actualErr := service.Get()

	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}
