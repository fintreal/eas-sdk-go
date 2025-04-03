package app

import "github.com/fintreal/expo-eas-sdk-go/internal/graphql"

type AppData struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type CreateAppData struct {
	Name      string
	Slug      string
	AccountId string
}

type UpdateAppData struct {
	Id   string
	Name string
}

type AppService interface {
	Create(data CreateAppData) (*AppData, error)
	Update(data UpdateAppData) (*AppData, error)
	Get(id string) (*AppData, error)
}

type appService struct {
	graphql graphql.GraphQL
}

var _ AppService = (*appService)(nil)

func NewAppService(graphql graphql.GraphQL) AppService {
	return &appService{
		graphql: graphql,
	}
}
