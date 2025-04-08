package certificate

import "github.com/fintreal/eas-sdk-go/internal/graphql"

type CertificateData struct {
	Id           string `json:"id"`
	SerialNumber string `json:"serialNumber"`
	// P12Base64    string `json:"certificateP12"`
	// Password     string `json:"certificatePassword"`
	// PrivateKey   string `json:"certificatePrivateSigningKey"`
}

type GetBySerialNumberCertificateData struct {
	SerialNumber string
	AccountId    string
}

type CertificateService interface {
	GetBySerialNumber(GetBySerialNumberCertificateData) (*CertificateData, error)
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
