package android

import (
	"github.com/fintreal/eas-sdk-go/internal/api/android/appcredentials"
	"github.com/fintreal/eas-sdk-go/internal/api/android/googleserviceaccountkey"
	"github.com/fintreal/eas-sdk-go/internal/graphql"
)

type Service struct {
	GoogleServiceAccountKey googleserviceaccountkey.Service
	AppCredentials          appcredentials.Service
}

func NewService(graphQL graphql.GraphQL) Service {
	return Service{
		GoogleServiceAccountKey: googleserviceaccountkey.NewService(graphQL),
		AppCredentials:          appcredentials.NewService(graphQL),
	}
}
