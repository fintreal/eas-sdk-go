package test

import (
	"os"
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/api/account"
	"github.com/stretchr/testify/assert"
)

func TestGetAccountByName(t *testing.T) {
	expectedData := &account.AccountData{
		Id:   os.Getenv("EXPO_ACCOUNT_ID"),
		Name: os.Getenv("EXPO_ACCOUNT_NAME"),
	}

	actualData, actualErr := client.Account.GetByName(expectedData.Name)

	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}
