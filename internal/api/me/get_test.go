package me

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/graphql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGet(t *testing.T) {
	var expectedData = &MeData{
		Id:          "test-id",
		DisplayName: "Test Display Name",
	}

	graphQLMock := newGetGraphQLMock(expectedData)

	service := NewMeService(graphQLMock)

	actualData, actualErr := service.Get()

	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}

func newGetGraphQLMock(data *MeData) *graphql.GraphQLMock {
	graphQLMock := graphql.NewGraphQLMock()
	graphQLMock.On("Query", mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		*args.Get(2).(*getMeResponse) = getMeResponse{Data: data}
	}).Return(nil)
	return graphQLMock
}
