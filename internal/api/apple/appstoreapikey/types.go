package appstoreapikey

import "github.com/fintreal/eas-sdk-go/internal/graphql"

type AppStoreApiKeyData struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	IssuerIdentifier string `json:"issuerIdentifier"`
	KeyIdentifier    string `json:"keyIdentifier"`
}

type AppStoreApiKeyService interface {
	GetByIdentifier(identifier string) (*AppStoreApiKeyData, error)
}

type appStoreApiKeyService struct {
	graphql graphql.GraphQL
}

var _ AppStoreApiKeyService = (*appStoreApiKeyService)(nil)

func NewAppStoreApiKeyService(graphql graphql.GraphQL) AppStoreApiKeyService {
	return &appStoreApiKeyService{
		graphql: graphql,
	}
}
