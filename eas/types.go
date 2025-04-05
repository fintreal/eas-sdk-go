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
)

type MeData = me.MeData
type AccountData = account.AccountData

type AppData = app.AppData
type CreateAppData = app.CreateAppData
type UpdateAppData = app.UpdateAppData

type AppVariableData = appvariable.AppVariableData
type CreateAppVariableData = appvariable.CreateAppVariableData
type UpdateAppVariableData = appvariable.UpdateAppVariableData

type AppleTeamData = appleteam.AppleTeamData
type CreateAppleTeamData = appleteam.CreateAppleTeamData
type UpdateAppleTeamData = appleteam.UpdateAppleTeamData

type AppleAppBundleIdentifierData = appleappbundleidentifier.AppleAppBundleIdentifierData
type CreateAppleAppBundleIdentifierData = appleappbundleidentifier.CreateAppleAppBundleIdentifierData

type AppleCertificateData = applecertificate.AppleCertificateData

type ProvisioningProfileData = provisioningprofile.ProvisioningProfileData
type CreateProvisioningProfileData = provisioningprofile.CreateProvisioningProfileData
