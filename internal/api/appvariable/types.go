package appvariable

import (
	"github.com/fintreal/eas-sdk-go/internal/graphql"
)

type AppVariableData struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
	// PUBLIC, SENSITIVE, SECRET
	Visibility string `json:"visibility"`
	// DEVELOPMENT, PREVIEW, PRODUCTION
	Environments []string `json:"environments"`
}

type GetByNameAppVariableData struct {
	Name  string
	AppId string
}

type GetAppVariableData struct {
	Id    string
	AppId string
}

type CreateAppVariableData struct {
	AppId        string
	Name         string
	Value        string
	Visibility   string
	Environments []string
}

type UpdateAppVariableData struct {
	Id           string
	Name         string
	Value        string
	Visibility   string
	Environments []string
}

type AppVariableService interface {
	Get(GetAppVariableData) (*AppVariableData, error)
	GetByName(GetByNameAppVariableData) (*AppVariableData, error)
	Create(CreateAppVariableData) (*AppVariableData, error)
	Update(UpdateAppVariableData) (*AppVariableData, error)
	Delete(string) (*any, error)
}

type appVariableService struct {
	graphql graphql.GraphQL
}

var _ AppVariableService = (*appVariableService)(nil)

func NewAppVariableService(graphql graphql.GraphQL) AppVariableService {
	return &appVariableService{graphql: graphql}
}
