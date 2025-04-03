package account

import (
	"testing"

	"github.com/fintreal/expo-eas-sdk-go/internal/graphql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetByName(t *testing.T) {
	expectedData := &AccountData{
		Id:   "test-account-id",
		Name: "test-account-name",
	}
	expectedVariables := map[string]any{"name": expectedData.Name}

	graphQLMock := newGetGraphQLMock(expectedData)

	service := NewAccountService(graphQLMock)

	acutalData, actualErr := service.GetByName(expectedData.Name)

	actualQuery := graphQLMock.Calls[0].Arguments.Get(0).(string)
	actualVariables := graphQLMock.Calls[0].Arguments.Get(1).(map[string]any)

	assert.Equal(t, nil, actualErr)
	assert.Equal(t, getAccountByNameQuery, actualQuery)
	assert.Equal(t, expectedVariables, actualVariables)
	assert.Equal(t, expectedData, acutalData)
}

func newGetGraphQLMock(data *AccountData) *graphql.GraphQLMock {
	graphQLMock := graphql.NewGraphQLMock()
	graphQLMock.On("Query", mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		*args.Get(2).(*getAccountResponse) = getAccountResponse{Account: getAccount{ByName: data}}
	}).Return(nil)
	return graphQLMock
}
