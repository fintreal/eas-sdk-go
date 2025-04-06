package appletest

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/eas"
	"github.com/fintreal/eas-sdk-go/test/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetAppleAppBundleIdentifierByIdentifier(t *testing.T) {
	expectedData := &eas.AppleAppBundleIdentifierData{
		Id:         "41ea486c-676a-4723-97b8-e0f80c53845f",
		Identifier: "my.test.app.identifier",
		TeamId:     utils.ImmutableAppleTeamId,
	}

	actualData, actualErr := utils.Client.Apple.BundleIdentifier.GetByIdentifier(expectedData.Identifier, utils.AccountId)

	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}

func TestCreateAppleAppBundleIdentifier(t *testing.T) {
	expectedIdentifier := utils.GenerateRandomString(10)
	expectedTeamId := utils.ImmutableAppleTeamId

	input := eas.CreateAppleAppBundleIdentifierData{
		AccountId:  utils.AccountId,
		Identifier: expectedIdentifier,
		TeamId:     expectedTeamId,
	}

	actualData, actualErr := utils.Client.Apple.BundleIdentifier.Create(input)

	assert.Equal(t, expectedIdentifier, actualData.Identifier)
	assert.Equal(t, expectedTeamId, actualData.TeamId)
	assert.Equal(t, nil, actualErr)
}
