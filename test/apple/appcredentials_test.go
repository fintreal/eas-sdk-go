package appletest

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/eas"
	"github.com/fintreal/eas-sdk-go/test/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetAppCredentials(t *testing.T) {
	input := eas.GetAppCredentialsData{
		Id:    utils.ImmutableAppCredentialsId,
		AppId: utils.ImmutableAppId,
	}
	actualData, actualErr := utils.Client.Apple.AppCredentials.Get(input)

	expectedData := &eas.AppCredentialsData{
		Id:              input.Id,
		AppId:           input.AppId,
		AppIdentifierId: "90289e95-99c2-4b21-88cc-8d84b6ab7477",
	}

	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}

func TestCreateAndDeleteAppCredentials(t *testing.T) {
	input := eas.CreateAppCredentialsData{
		AppId:           utils.ImmutableAppId,
		AppIdentifierId: utils.MutableAppIdentifierId,
	}

	actualData, actualErr := utils.Client.Apple.AppCredentials.Create(input)
	assert.Equal(t, input.AppId, actualData.AppId)
	assert.Equal(t, input.AppIdentifierId, actualData.AppIdentifierId)
	assert.Equal(t, nil, actualErr)

	_, actualErr = utils.Client.Apple.AppCredentials.Delete(actualData.Id)

	assert.Equal(t, nil, actualErr)
}
