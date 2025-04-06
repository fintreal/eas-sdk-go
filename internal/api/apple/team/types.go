package team

import "github.com/fintreal/eas-sdk-go/internal/graphql"

type account struct {
	Id string `json:"id"`
}

type teamData struct {
	Id         string  `json:"id"`
	Name       string  `json:"appleTeamName"`
	Identifier string  `json:"appleTeamIdentifier"`
	Type       string  `json:"appleTeamType"`
	Account    account `json:"account"`
}

type TeamData struct {
	Id         string
	Name       string
	Identifier string
	Type       string
	AccountId  string
}

type CreateTeamData struct {
	Name       string
	Identifier string
	Type       string
	AccountId  string
}

type UpdateTeamData struct {
	Id   string
	Name string
	Type string
}

type TeamService interface {
	Create(data CreateTeamData) (*TeamData, error)
	Update(data UpdateTeamData) (*TeamData, error)
	GetByIdentifier(identifier string, accountId string) (*TeamData, error)
}

type teamService struct {
	graphql graphql.GraphQL
}

var _ TeamService = (*teamService)(nil)

func NewTeamService(graphql graphql.GraphQL) TeamService {
	return &teamService{
		graphql: graphql,
	}
}
