package eas

import (
	"github.com/fintreal/eas-sdk-go/internal/api/account"
	"github.com/fintreal/eas-sdk-go/internal/api/app"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/provisioningprofile"
	"github.com/fintreal/eas-sdk-go/internal/api/appleappbundleidentifier"
	"github.com/fintreal/eas-sdk-go/internal/api/applecertificate"
	"github.com/fintreal/eas-sdk-go/internal/api/appleteam"
	"github.com/fintreal/eas-sdk-go/internal/api/appvariable"
	"github.com/fintreal/eas-sdk-go/internal/api/me"
	"github.com/fintreal/eas-sdk-go/internal/graphql"
)

// EASClient capable of interacting with Expo EAS GraphQL API
type EASClient struct {
	Me                       me.MeService
	App                      app.AppService
	AppVariable              appvariable.AppVariableService
	Account                  account.AccountService
	AppleTeam                appleteam.AppleTeamService
	AppleAppBundleIdentifier appleappbundleidentifier.AppleAppBundleIdentifierService
	AppleCertificate         applecertificate.AppleCertificateService
	AppleProvisioningProfile provisioningprofile.ProvisioningProfileService
}

// EASClient capable of interacting with Expo EAS GraphQL API
//
// @token Expo Personal Access Token or Robot Access Token
func NewEASClient(token string) *EASClient {
	if token == "" {
		panic("expo token can't be an empty string")
	}
	graphql := graphql.NewGraphQL(token)
	return &EASClient{
		Me:                       me.NewMeService(graphql),
		App:                      app.NewAppService(graphql),
		AppVariable:              appvariable.NewAppVariableService(graphql),
		Account:                  account.NewAccountService(graphql),
		AppleTeam:                appleteam.NewAppleTeamService(graphql),
		AppleAppBundleIdentifier: appleappbundleidentifier.NewAppleAppBundleIdentifierService(graphql),
		AppleCertificate:         applecertificate.NewAppleCertificateService(graphql),
		AppleProvisioningProfile: provisioningprofile.NewProvisioningProfileService(graphql),
	}
}
