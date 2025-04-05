package applecertificate

import "github.com/fintreal/eas-sdk-go/internal/graphql"

type AppleCertificateData struct {
	Id           string `json:"id"`
	SerialNumber string `json:"serialNumber"`
	// P12Base64    string `json:"certificateP12"`
	// Password     string `json:"certificatePassword"`
	// PrivateKey   string `json:"certificatePrivateSigningKey"`
}

type AppleCertificateService interface {
	GetAppleCertificateBySerialNumber(serialNumber string, accountId string) (*AppleCertificateData, error)
}

type appleCertificateService struct {
	graphql graphql.GraphQL
}

var _ AppleCertificateService = (*appleCertificateService)(nil)

func NewAppleCertificateService(graphql graphql.GraphQL) AppleCertificateService {
	return &appleCertificateService{
		graphql: graphql,
	}
}
