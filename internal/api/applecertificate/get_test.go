package applecertificate

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/graphql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGet(t *testing.T) {
	accountId := "test-account-id"
	expectedResponse := &AppleCertificateData{
		Id:           "test-id",
		SerialNumber: "test-serial-number",
		// P12Base64:    "test-p12-base64",
		// Password:     "test-password",
		// PrivateKey:   "test-private-key",
	}
	expectedVariables := map[string]any{"accountId": accountId}
	graphQLMock := newGetGraphQLMock(expectedResponse)
	service := NewAppleCertificateService(graphQLMock)
	result, err := service.GetAppleCertificateBySerialNumber(expectedResponse.SerialNumber, accountId)

	actualQuery := graphQLMock.Calls[0].Arguments.Get(0).(string)
	actualVariables := graphQLMock.Calls[0].Arguments.Get(1).(map[string]any)

	assert.Equal(t, getAppleCertificateQuery, actualQuery)
	assert.Equal(t, expectedVariables, actualVariables)
	assert.Equal(t, expectedResponse, result)
	assert.Equal(t, nil, err)
}

func newGetGraphQLMock(data *AppleCertificateData) *graphql.GraphQLMock {
	mockResponse := getAppleCertifiactesResponse{
		Account: getAppleCertifiactes{
			ById: byId{
				Data: []AppleCertificateData{*data},
			},
		},
	}

	graphQLMock := graphql.NewGraphQLMock()
	graphQLMock.On("Query", mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		*args.Get(2).(*getAppleCertifiactesResponse) = mockResponse
	}).Return(nil)
	return graphQLMock
}
