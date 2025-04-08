package bundleidentifier

import "github.com/fintreal/eas-sdk-go/internal/graphql"

type team struct {
	Id string `json:"id"`
}

type data struct {
	Id         string `json:"id"`
	Identifier string `json:"bundleIdentifier"`
	Team       team   `json:"appleTeam"`
}

type Data struct {
	Id         string
	Identifier string
	TeamId     string
}

type CreateData struct {
	AccountId  string
	Identifier string
	TeamId     string
}

type GetByIdentifierData struct {
	Identifier string
	AccountId  string
}

type Service interface {
	Create(data CreateData) (*Data, error)
	GetByIdentifier(GetByIdentifierData) (*Data, error)
}

type service struct {
	graphql graphql.GraphQL
}

var _ Service = (*service)(nil)

func NewService(graphql graphql.GraphQL) Service {
	return &service{
		graphql: graphql,
	}
}
