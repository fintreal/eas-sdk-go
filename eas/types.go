package eas

import (
	"github.com/fintreal/eas-sdk-go/internal/api/account"
	"github.com/fintreal/eas-sdk-go/internal/api/app"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/bundleidentifier"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/certificate"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/provisioningprofile"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/team"
	"github.com/fintreal/eas-sdk-go/internal/api/appvariable"
	"github.com/fintreal/eas-sdk-go/internal/api/me"
)

type MeData = me.MeData
type AccountData = account.AccountData

type AppData = app.AppData
type CreateAppData = app.CreateAppData
type UpdateAppData = app.UpdateAppData

type AppVariableData = appvariable.AppVariableData
type CreateAppVariableData = appvariable.CreateAppVariableData
type UpdateAppVariableData = appvariable.UpdateAppVariableData

type AppleTeamData = team.TeamData
type CreateAppleTeamData = team.CreateTeamData
type UpdateAppleTeamData = team.UpdateTeamData

type AppleAppBundleIdentifierData = bundleidentifier.AppBundleIdentifierData
type CreateAppleAppBundleIdentifierData = bundleidentifier.CreateAppBundleIdentifierData

type AppleCertificateData = certificate.CertificateData

type ProvisioningProfileData = provisioningprofile.ProvisioningProfileData
type CreateProvisioningProfileData = provisioningprofile.CreateProvisioningProfileData
