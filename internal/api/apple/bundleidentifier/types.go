package bundleidentifier

import "github.com/fintreal/eas-sdk-go/internal/graphql"

type team struct {
	Id string `json:"id"`
}

type appBundleIdentifierData struct {
	Id         string `json:"id"`
	Identifier string `json:"bundleIdentifier"`
	Team       team   `json:"appleTeam"`
}

type AppBundleIdentifierData struct {
	Id         string
	Identifier string
	TeamId     string
}

type CreateAppBundleIdentifierData struct {
	AccountId  string
	Identifier string
	TeamId     string
}

type GetByIdentifierData struct {
	Identifier string
	AccountId  string
}

type AppBundleIdentifierService interface {
	Create(data CreateAppBundleIdentifierData) (*AppBundleIdentifierData, error)
	GetByIdentifier(GetByIdentifierData) (*AppBundleIdentifierData, error)
}

type appBundleIdentifierService struct {
	graphql graphql.GraphQL
}

var _ AppBundleIdentifierService = (*appBundleIdentifierService)(nil)

func NewAppBundleIdentifierService(graphql graphql.GraphQL) AppBundleIdentifierService {
	return &appBundleIdentifierService{
		graphql: graphql,
	}
}
