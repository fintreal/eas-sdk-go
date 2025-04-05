package appleappbundleidentifier

import "github.com/fintreal/eas-sdk-go/internal/graphql"

type appleTeam struct {
	Id string `json:"id"`
}

type appleAppBundleIdentifierData struct {
	Id         string    `json:"id"`
	Identifier string    `json:"bundleIdentifier"`
	AppleTeam  appleTeam `json:"appleTeam"`
}

type AppleAppBundleIdentifierData struct {
	Id         string
	Identifier string
	TeamId     string
}

type CreateAppleAppBundleIdentifierData struct {
	AccountId  string
	Identifier string
	TeamId     string
}

type AppleAppBundleIdentifierService interface {
	Create(data CreateAppleAppBundleIdentifierData) (*AppleAppBundleIdentifierData, error)
	GetByIdentifier(identifier string, appleTeamId string) (*AppleAppBundleIdentifierData, error)
}

type appleAppBundleIdentifierService struct {
	graphql graphql.GraphQL
}

var _ AppleAppBundleIdentifierService = (*appleAppBundleIdentifierService)(nil)

func NewAppleAppBundleIdentifierService(graphql graphql.GraphQL) AppleAppBundleIdentifierService {
	return &appleAppBundleIdentifierService{
		graphql: graphql,
	}
}
