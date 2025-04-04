package eas

import (
	"github.com/fintreal/eas-sdk-go/internal/api/account"
	"github.com/fintreal/eas-sdk-go/internal/api/app"
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
