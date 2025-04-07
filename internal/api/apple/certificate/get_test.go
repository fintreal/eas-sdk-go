package certificate

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/graphql"
	"github.com/fintreal/eas-sdk-go/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetBySerialNumber(t *testing.T) {
	accountId := "test-account-id"
	expectedResponse := &CertificateData{
		Id:           "test-id",
		SerialNumber: "test-serial-number",
		// P12Base64:    "test-p12-base64",
		// Password:     "test-password",
		// PrivateKey:   "test-private-key",
	}
	expectedVariables := map[string]any{"accountId": accountId}

	mockResponse := utils.AccountResponse[getCertificatesResponse]{
		Account: utils.Account[getCertificatesResponse]{
			ById: getCertificatesResponse{
				Data: []CertificateData{*expectedResponse},
			},
		},
	}
	graphQLMock := graphql.NewGraphQLMock(mockResponse)

	service := NewCertificateService(graphQLMock)
	actualResponse, actualErr := service.GetBySerialNumber(expectedResponse.SerialNumber, accountId)

	actualQuery := graphQLMock.Calls[0].Arguments.Get(0).(string)
	actualVariables := graphQLMock.Calls[0].Arguments.Get(1).(map[string]any)

	assert.Equal(t, getQuery, actualQuery)
	assert.Equal(t, expectedVariables, actualVariables)
	assert.Equal(t, expectedResponse, actualResponse)
	assert.Equal(t, nil, actualErr)
}
