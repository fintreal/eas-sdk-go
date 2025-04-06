package certificate

import "github.com/fintreal/eas-sdk-go/internal/graphql"

type CertificateData struct {
	Id           string `json:"id"`
	SerialNumber string `json:"serialNumber"`
	// P12Base64    string `json:"certificateP12"`
	// Password     string `json:"certificatePassword"`
	// PrivateKey   string `json:"certificatePrivateSigningKey"`
}

type CertificateService interface {
	GetBySerialNumber(serialNumber string, accountId string) (*CertificateData, error)
}

type certificateService struct {
	graphql graphql.GraphQL
}

var _ CertificateService = (*certificateService)(nil)

func NewCertificateService(graphql graphql.GraphQL) CertificateService {
	return &certificateService{
		graphql: graphql,
	}
}
