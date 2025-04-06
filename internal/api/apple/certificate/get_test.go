package certificate

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/graphql"
	"github.com/fintreal/eas-sdk-go/internal/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
	graphQLMock := newGetGraphQLMock(expectedResponse)
	service := NewCertificateService(graphQLMock)
	result, err := service.GetBySerialNumber(expectedResponse.SerialNumber, accountId)

	actualQuery := graphQLMock.Calls[0].Arguments.Get(0).(string)
	actualVariables := graphQLMock.Calls[0].Arguments.Get(1).(map[string]any)

	assert.Equal(t, getQuery, actualQuery)
	assert.Equal(t, expectedVariables, actualVariables)
	assert.Equal(t, expectedResponse, result)
	assert.Equal(t, nil, err)
}

func newGetGraphQLMock(data *CertificateData) *graphql.GraphQLMock {
	mockResponse := utils.AccountResponse[getCertificatesResponse]{
		Account: utils.Account[getCertificatesResponse]{
			ById: getCertificatesResponse{
				Data: []CertificateData{*data},
			},
		},
	}

	graphQLMock := graphql.NewGraphQLMock()
	graphQLMock.On("Query", mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		*args.Get(2).(*utils.AccountResponse[getCertificatesResponse]) = mockResponse
	}).Return(nil)
	return graphQLMock
}
