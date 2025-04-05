package test

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/api/account"
	"github.com/stretchr/testify/assert"
)

func TestGetAccountByName(t *testing.T) {
	expectedData := &account.AccountData{
		Id:   accountId,
		Name: "expo-eas-sdk-go",
	}

	actualData, actualErr := client.Account.GetByName(expectedData.Name)

	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}
