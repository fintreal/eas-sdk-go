package eas

import (
	"github.com/fintreal/eas-sdk-go/internal/api/account"
	"github.com/fintreal/eas-sdk-go/internal/api/android/googleserviceaccountkey"
	"github.com/fintreal/eas-sdk-go/internal/api/app"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/appbuildcredentials"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/appcredentials"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/appidentifier"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/appstoreapikey"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/certificate"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/provisioningprofile"
	"github.com/fintreal/eas-sdk-go/internal/api/apple/team"
	"github.com/fintreal/eas-sdk-go/internal/api/appvariable"
	"github.com/fintreal/eas-sdk-go/internal/api/me"
)

type MeData = me.Data
type AccountData = account.Data

type AppData = app.Data
type CreateAppData = app.CreateData
type UpdateAppData = app.UpdateData

type AppVariableData = appvariable.Data
type CreateAppVariableData = appvariable.CreateData
type UpdateAppVariableData = appvariable.UpdateData
type GetAppVariableData = appvariable.GetData
type GetByNameAppVariableData = appvariable.GetByNameData

type AppleTeamData = team.Data
type CreateAppleTeamData = team.CreateData
type UpdateAppleTeamData = team.UpdateData
type GetByIdentifierAppleTeamData = team.GetByIdentifierData

type AppleAppIdentifierData = appidentifier.Data
type CreateAppleAppIdentifierData = appidentifier.CreateData
type GetByIdentifierAppleAppIdentifierData = appidentifier.GetByIdentifierData

type AppleCertificateData = certificate.Data
type GetBySerialNumberAppleCertificateData = certificate.GetBySerialNumberData

type ProvisioningProfileData = provisioningprofile.Data
type CreateProvisioningProfileData = provisioningprofile.CreateData
type GetProvisioningProfileData = provisioningprofile.GetData

type AppStoreApiKeyData = appstoreapikey.Data
type GetByIdentifierAppStoreApiKeyData = appstoreapikey.GeyByIdentifierData

type AppCredentialsData = appcredentials.Data
type CreateAppCredentialsData = appcredentials.CreateData
type UpdateAppCredentialsData = appcredentials.UpdateData
type GetAppCredentialsData = appcredentials.GetData

type AppBuildCredentialsData = appbuildcredentials.Data
type CreateAppBuildCredentialsData = appbuildcredentials.CreateData
type GetAppBuildCredentialsData = appbuildcredentials.GetData

type GetByProjectIdentifierGoogleServiceAccountKeyData = googleserviceaccountkey.GetByProjectIdentifierData
type GoogleServiceAccountKeyData = googleserviceaccountkey.Data
