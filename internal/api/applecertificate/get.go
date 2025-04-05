package applecertificate

type byId struct {
	Data []AppleCertificateData `json:"appleDistributionCertificates"`
}

type getAppleCertificates struct {
	ById byId `json:"byId"`
}

type getAppleCertificatesResponse struct {
	Account getAppleCertificates `json:"account"`
}

const getAppleCertificateQuery = `
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

func (service *appleCertificateService) GetAppleCertificateBySerialNumber(serialNumber string, accountId string) (*AppleCertificateData, error) {
	variables := map[string]any{
		"accountId": accountId,
	}

	var response getAppleCertificatesResponse

	err := service.graphql.Query(getAppleCertificateQuery, variables, &response)
	if err != nil {
		return nil, err
	}
	return findAppleCertificateBySerialNumber(response.Account.ById.Data, serialNumber)
}

func findAppleCertificateBySerialNumber(certificates []AppleCertificateData, serialNumber string) (*AppleCertificateData, error) {
	for _, certificate := range certificates {
		if certificate.SerialNumber == serialNumber {
			return &certificate, nil
		}
	}
	return nil, nil
}
