package account

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/graphql"
	"github.com/stretchr/testify/assert"
)

func TestGetByName(t *testing.T) {
	expectedData := &AccountData{
		Id:   "test-account-id",
		Name: "test-account-name",
	}
	expectedVariables := map[string]any{"name": expectedData.Name}

	mockResponse := getAccountResponse{Account: getAccount{ByName: expectedData}}
	graphQLMock := graphql.NewGraphQLMock(mockResponse)

	service := NewAccountService(graphQLMock)

	acutalData, actualErr := service.GetByName(expectedData.Name)

	actualQuery := graphQLMock.Calls[0].Arguments.Get(0).(string)
	actualVariables := graphQLMock.Calls[0].Arguments.Get(1).(map[string]any)

	assert.Equal(t, nil, actualErr)
	assert.Equal(t, getAccountByNameQuery, actualQuery)
	assert.Equal(t, expectedVariables, actualVariables)
	assert.Equal(t, expectedData, acutalData)
}
