package apple

import (
	"github.com/fintreal/eas-sdk-go/internal/api/apple/appbuildcredentials"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/appcredentials"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/appstoreapikey"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/bundleidentifier"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/certificate"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/provisioningprofile"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/team"
	"github.com/fintreal/eas-sdk-go/internal/graphql"
)

type Service struct {
	Team                team.Service
	Certificate         certificate.Service
	BundleIdentifier    bundleidentifier.Service
	ProvisioningProfile provisioningprofile.Service
	AppStoreApiKey      appstoreapikey.Service
	AppCredentials      appcredentials.Service
	AppBuildCredentials appbuildcredentials.Service
}

func NewService(graphQL graphql.GraphQL) Service {
	return Service{
		Team:                team.NewService(graphQL),
		Certificate:         certificate.NewService(graphQL),
		BundleIdentifier:    bundleidentifier.NewService(graphQL),
		ProvisioningProfile: provisioningprofile.NewService(graphQL),
		AppStoreApiKey:      appstoreapikey.NewService(graphQL),
		AppCredentials:      appcredentials.NewService(graphQL),
		AppBuildCredentials: appbuildcredentials.NewService(graphQL),
	}
}
