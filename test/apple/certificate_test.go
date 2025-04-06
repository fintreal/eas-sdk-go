package appletest

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/eas"
	"github.com/fintreal/eas-sdk-go/test/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetAppleCertificate(t *testing.T) {
	expectedData := &eas.AppleCertificateData{
		Id:           "702635c5-3aa1-477c-83b6-bb66a1644aad",
		SerialNumber: "3D986E25FF1B48C2417853A07AA15C55",
	}
	actualData, actualErr := utils.Client.Apple.Certificate.GetBySerialNumber(expectedData.SerialNumber, utils.AccountId)
	assert.Equal(t, expectedData, actualData)
	assert.NoError(t, actualErr)
}
