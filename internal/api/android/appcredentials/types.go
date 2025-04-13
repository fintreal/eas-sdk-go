package appcredentials

import "github.com/fintreal/eas-sdk-go/internal/graphql"

type objWithId struct {
	Id string `json:"id"`
}

type data struct {
	Id                      string    `json:"id"`
	Identifier              string    `json:"applicationIdentifier"`
	App                     objWithId `json:"app"`
	GoogleServiceAccountKey objWithId `json:"googleServiceAccountKeyForSubmissions"`
}

type Data struct {
	Id                        string
	AppId                     string
	Identifier                string
	GoogleServiceAccountKeyId string
}

type GetData struct {
	Id    string
	AppId string
}

type CreateData struct {
	AppId                     string
	Identifier                string
	GoogleServiceAccountKeyId string
}

type Service interface {
	Get(GetData) (*Data, error)
	Create(CreateData) (*Data, error)
	Delete(string) (*any, error)
}

type service struct {
	graphql graphql.GraphQL
}

var _ Service = (*service)(nil)

func NewService(graphql graphql.GraphQL) Service {
	return &service{graphql: graphql}
}
