package appcredentials

import (
	"github.com/fintreal/eas-sdk-go/internal/api/android/appbuildcredentials"
	"github.com/fintreal/eas-sdk-go/internal/graphql"
)

type objWithId struct {
	Id string `json:"id"`
}

type buildCredentialsData struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Keystore objWithId `json:"androidKeystore"`
}

type data struct {
	Id                      string                 `json:"id"`
	Identifier              string                 `json:"applicationIdentifier"`
	App                     objWithId              `json:"app"`
	GoogleServiceAccountKey objWithId              `json:"googleServiceAccountKeyForSubmissions"`
	BuildCredentials        []buildCredentialsData `json:"androidAppBuildCredentialsArray"`
}

type Data struct {
	Id                        string
	AppId                     string
	Identifier                string
	GoogleServiceAccountKeyId string
	BuildCredentials          []appbuildcredentials.Data
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
