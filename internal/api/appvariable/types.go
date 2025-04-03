package appvariable

import (
	"github.com/fintreal/expo-eas-sdk-go/internal/graphql"
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

type CreateAppVariableData struct {
	AppId string `json:"appId"`
	Name  string `json:"name"`
	Value string `json:"value"`
	// PUBLIC, SENSITIVE, SECRET
	Visibility string `json:"visibility"`
	// DEVELOPMENT, PREVIEW, PRODUCTION
	Environments []string `json:"environments"`
}

type UpdateAppVariableData struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
	// PUBLIC, SENSITIVE, SECRET
	Visibility string `json:"visibility"`
	// DEVELOPMENT, PREVIEW, PRODUCTION
	Environments []string `json:"environments"`
}

type AppVariableService interface {
	Get(id string, appId string) (*AppVariableData, error)
	GetByName(name string, appId string) (*AppVariableData, error)
	Create(createData CreateAppVariableData) (*AppVariableData, error)
	Update(data UpdateAppVariableData) (*AppVariableData, error)
	Delete(id string) error
}

type appVariableService struct {
	graphql graphql.GraphQL
}

var _ AppVariableService = (*appVariableService)(nil)

func NewAppVariableService(graphql graphql.GraphQL) AppVariableService {
	return &appVariableService{graphql: graphql}
}
