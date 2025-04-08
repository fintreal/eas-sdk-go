package certificate

import "github.com/fintreal/eas-sdk-go/internal/utils"

type getCertificatesResponse struct {
	Data []CertificateData `json:"appleDistributionCertificates"`
}

const getQuery = `
	query ($accountId: String!) {
			account {
					byId(accountId: $accountId) {
							appleDistributionCertificates {
									id
									serialNumber
							}
					}
			}
	}`

func (service *certificateService) GetBySerialNumber(getData GetBySerialNumberCertificateData) (*CertificateData, error) {
	variables := map[string]any{
		"accountId":    getData.AccountId,
		"serialNumber": getData.SerialNumber,
	}

	var response utils.AccountResponse[getCertificatesResponse]

	err := service.graphql.Query(getQuery, variables, &response)
	if err != nil {
		return nil, err
	}
	return findBySerialNumber(response.Account.ById.Data, getData.SerialNumber)
}

func findBySerialNumber(certificates []CertificateData, serialNumber string) (*CertificateData, error) {
	for _, certificate := range certificates {
		if certificate.SerialNumber == serialNumber {
			return &certificate, nil
		}
	}
	return nil, nil
}
