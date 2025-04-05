package appleteam

import "github.com/fintreal/eas-sdk-go/internal/graphql"

type account struct {
	Id string `json:"id"`
}

type appleTeamData struct {
	Id         string  `json:"id"`
	Name       string  `json:"appleTeamName"`
	Identifier string  `json:"appleTeamIdentifier"`
	Type       string  `json:"appleTeamType"`
	Account    account `json:"account"`
}

type AppleTeamData struct {
	Id         string
	Name       string
	Identifier string
	Type       string
	AccountId  string
}

type CreateAppleTeamData struct {
	Name       string
	Identifier string
	Type       string
	AccountId  string
}

type UpdateAppleTeamData struct {
	Id   string
	Name string
	Type string
}

type AppleTeamService interface {
	Create(data CreateAppleTeamData) (*AppleTeamData, error)
	Update(data UpdateAppleTeamData) (*AppleTeamData, error)
	GetByIdentifier(identifier string, accountId string) (*AppleTeamData, error)
}

type appleTeamService struct {
	graphql graphql.GraphQL
}

var _ AppleTeamService = (*appleTeamService)(nil)

func NewAppleTeamService(graphql graphql.GraphQL) AppleTeamService {
	return &appleTeamService{
		graphql: graphql,
	}
}
