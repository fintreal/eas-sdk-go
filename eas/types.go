package eas

import (
	"github.com/fintreal/eas-sdk-go/internal/api/account"
	"github.com/fintreal/eas-sdk-go/internal/api/app"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/appstoreapikey"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/bundleidentifier"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/certificate"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/provisioningprofile"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/team"
	"github.com/fintreal/eas-sdk-go/internal/api/appvariable"
	"github.com/fintreal/eas-sdk-go/internal/api/me"
)

type MeData = me.Data
type AccountData = account.Data

type AppData = app.AppData
type CreateAppData = app.CreateAppData
type UpdateAppData = app.UpdateAppData

type AppVariableData = appvariable.AppVariableData
type CreateAppVariableData = appvariable.CreateAppVariableData
type UpdateAppVariableData = appvariable.UpdateAppVariableData
type GetAppVariableData = appvariable.GetAppVariableData
type GetByNameAppVariableData = appvariable.GetByNameAppVariableData

type AppleTeamData = team.TeamData
type CreateAppleTeamData = team.CreateTeamData
type UpdateAppleTeamData = team.UpdateTeamData
type GetByIdentifierAppleTeamData = team.GetByIdentifierData

type AppleAppBundleIdentifierData = bundleidentifier.AppBundleIdentifierData
type CreateAppleAppBundleIdentifierData = bundleidentifier.CreateAppBundleIdentifierData
type GetByIdentifierAppleBundleIdentifierData = bundleidentifier.GetByIdentifierData

type AppleCertificateData = certificate.CertificateData
type GetBySerialNumberAppleCertificateData = certificate.GetBySerialNumberCertificateData

type ProvisioningProfileData = provisioningprofile.ProvisioningProfileData
type CreateProvisioningProfileData = provisioningprofile.CreateProvisioningProfileData
type GetProvisioningProfileData = provisioningprofile.GetProvisioningProfileData

type AppStoreApiKeyData = appstoreapikey.AppStoreApiKeyData
type GetByIdentifierAppStoreApiKeyData = appstoreapikey.GeyByIdentifierAppStoreApiKeyData
