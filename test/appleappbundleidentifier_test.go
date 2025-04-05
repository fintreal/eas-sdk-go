package test

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/eas"
	"github.com/stretchr/testify/assert"
)

func TestGetAppleAppBundleIdentifierByIdentifier(t *testing.T) {
	expectedData := &eas.AppleAppBundleIdentifierData{
		Id:         "41ea486c-676a-4723-97b8-e0f80c53845f",
		Identifier: "my.test.app.identifier",
		TeamId:     immutableAppleTeamId,
	}

	actualData, actualErr := client.AppleAppBundleIdentifier.GetByIdentifier(expectedData.Identifier, accountId)

	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}

func TestCreateAppleAppBundleIdentifier(t *testing.T) {
	expectedIdentifier := generateRandomString(10)
	expectedTeamId := immutableAppleTeamId

	input := eas.CreateAppleAppBundleIdentifierData{
		AccountId:  accountId,
		Identifier: expectedIdentifier,
		TeamId:     expectedTeamId,
	}

	actualData, actualErr := client.AppleAppBundleIdentifier.Create(input)

	assert.Equal(t, expectedIdentifier, actualData.Identifier)
	assert.Equal(t, expectedTeamId, actualData.TeamId)
	assert.Equal(t, nil, actualErr)
}
