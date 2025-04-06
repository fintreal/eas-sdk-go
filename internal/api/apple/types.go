package apple

import (
	"github.com/fintreal/eas-sdk-go/internal/api/apple/bundleidentifier"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/certificate"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/provisioningprofile"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/team"
	"github.com/fintreal/eas-sdk-go/internal/graphql"
)

type AppleService struct {
	Team                team.TeamService
	Certificate         certificate.CertificateService
	BundleIdentifier    bundleidentifier.AppBundleIdentifierService
	ProvisioningProfile provisioningprofile.ProvisioningProfileService
}

func NewAppleService(graphQL graphql.GraphQL) AppleService {
	return AppleService{
		Team:                team.NewTeamService(graphQL),
		Certificate:         certificate.NewCertificateService(graphQL),
		BundleIdentifier:    bundleidentifier.NewAppBundleIdentifierService(graphQL),
		ProvisioningProfile: provisioningprofile.NewProvisioningProfileService(graphQL),
	}
}
