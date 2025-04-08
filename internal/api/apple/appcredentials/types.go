package appcredentials

import (
	"github.com/fintreal/eas-sdk-go/internal/graphql"
)

type objWithId struct {
	Id string `json:"id"`
}

type data struct {
	Id            string    `json:"id"`
	App           objWithId `json:"app"`
	AppIdentifier objWithId `json:"appleAppIdentifier"`
}

type Data struct {
	Id              string `json:"id"`
	AppIdentifierId string `json:"appIdentifier"`
	AppId           string `json:"appId"`
}

type GetData struct {
	Id    string
	AppId string
}

type CreateData struct {
	AppIdentifierId string
	AppId           string
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
