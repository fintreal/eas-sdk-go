package appletest

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/eas"
	"github.com/fintreal/eas-sdk-go/test/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetAppleAppBundleIdentifierByIdentifier(t *testing.T) {
	input := eas.GetByIdentifierAppleBundleIdentifierData{
		Identifier: "my.test.app.identifier",
		AccountId:  utils.AccountId,
	}

	actualData, actualErr := utils.Client.Apple.BundleIdentifier.GetByIdentifier(input)

	expectedData := &eas.AppleAppBundleIdentifierData{
		Id:         "41ea486c-676a-4723-97b8-e0f80c53845f",
		Identifier: input.Identifier,
		TeamId:     utils.ImmutableAppleTeamId,
	}

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
