package account

import "github.com/fintreal/eas-sdk-go/internal/graphql"

type AccountData struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type AccountService interface {
	GetByName(name string) (*AccountData, error)
}

type accountService struct {
	graphql graphql.GraphQL
}

var _ AccountService = (*accountService)(nil)

func NewAccountService(graphql graphql.GraphQL) AccountService {
	return &accountService{graphql: graphql}
}
