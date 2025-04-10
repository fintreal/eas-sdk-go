package appletest

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/eas"
	"github.com/fintreal/eas-sdk-go/test/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetAppleAppBundleIdentifierByIdentifier(t *testing.T) {
	input := eas.GetByIdentifierAppleBundleIdentifierData{
		Identifier: utils.ImmutableAppIdentifierName,
		AccountId:  utils.AccountId,
	}

	actualData, actualErr := utils.Client.Apple.BundleIdentifier.GetByIdentifier(input)

	expectedData := &eas.AppleAppBundleIdentifierData{
		Id:         utils.ImmutableAppIdentifierId,
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
