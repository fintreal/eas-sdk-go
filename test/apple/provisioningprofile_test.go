package appletest

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/eas"
	"github.com/fintreal/eas-sdk-go/test/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetProvisioningProfile(t *testing.T) {
	expectedData := &eas.ProvisioningProfileData{
		Id:                    "8690db1b-c475-43d0-aa3f-67e103c96426",
		Base64:                utils.ImmutableProvisioningProfileBase64,
		AppBundleIdentifierId: "554fe9ea-dadd-49ec-8300-1940661f352b",
	}

	actualData, actualErr := utils.Client.Apple.ProvisioningProfile.Get(expectedData.Id, utils.AccountId)
	assert.Equal(t, expectedData, actualData)
	assert.NoError(t, actualErr)
}

func TestCreateAndDeleteProvisioningProfile(t *testing.T) {
	expectedBase64 := utils.ImmutableProvisioningProfileBase64
	expectedAppBundleIdentifierId := "554fe9ea-dadd-49ec-8300-1940661f352b"

	input := eas.CreateProvisioningProfileData{
		Base64:                utils.ImmutableProvisioningProfileBase64,
		AppBundleIdentifierId: "554fe9ea-dadd-49ec-8300-1940661f352b",
		AccountId:             utils.AccountId,
	}

	actualData, actualErr := utils.Client.Apple.ProvisioningProfile.Create(input)
	assert.Equal(t, expectedBase64, actualData.Base64)
	assert.Equal(t, expectedAppBundleIdentifierId, actualData.AppBundleIdentifierId)
	assert.NoError(t, actualErr)

	actualErr = utils.Client.Apple.ProvisioningProfile.Delete(actualData.Id)
	assert.NoError(t, actualErr)
}
