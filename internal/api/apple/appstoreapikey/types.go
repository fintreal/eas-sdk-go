package appstoreapikey

import "github.com/fintreal/eas-sdk-go/internal/graphql"

type AppStoreApiKeyData struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	IssuerIdentifier string `json:"issuerIdentifier"`
	Identifier       string `json:"keyIdentifier"`
}

type GeyByIdentifierAppStoreApiKeyData struct {
	Identifier string
	AccountId  string
}

type AppStoreApiKeyService interface {
	GetByIdentifier(GeyByIdentifierAppStoreApiKeyData) (*AppStoreApiKeyData, error)
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
