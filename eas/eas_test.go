package eas

import (
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/fintreal/expo-eas-sdk-go/internal/api/account"
	"github.com/fintreal/expo-eas-sdk-go/internal/api/app"
	"github.com/fintreal/expo-eas-sdk-go/internal/api/appvariable"
	"github.com/fintreal/expo-eas-sdk-go/internal/api/me"
	"github.com/stretchr/testify/assert"
)

var token = os.Getenv("EXPO_TOKEN")

var client = NewEASClient(token)

// Me

func TestGetMe(t *testing.T) {
	expectedData := &me.MeData{
		Id:          os.Getenv("EXPO_ME_ID"),
		DisplayName: "integration-test",
	}

	actualData, actualErr := client.Me.Get()
	fmt.Printf("%+v", actualData)
	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}

// Account

func TestGetAccountByName(t *testing.T) {
	expectedData := &account.AccountData{
		Id:   os.Getenv("EXPO_ACCOUNT_ID"),
		Name: os.Getenv("EXPO_ACCOUNT_NAME"),
	}

	actualData, actualErr := client.Account.GetByName(expectedData.Name)

	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}

// App

func TestAppGet(t *testing.T) {
	expectedData := &app.AppData{
		Id:   os.Getenv("EXPO_TEST_GET_APP_ID"),
		Name: "Test App",
		Slug: "test-app",
	}

	actualData, actualErr := client.App.Get(expectedData.Id)

	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}

func TestAppCreate(t *testing.T) {
	inputData := app.CreateAppData{
		AccountId: os.Getenv("EXPO_ACCOUNT_ID"),
		Name:      generateRandomString(10),
		Slug:      generateRandomString(10),
	}

	// Create
	expectedData, actualErr := client.App.Create(inputData)

	assert.Equal(t, inputData.Name, expectedData.Name)
	assert.Equal(t, inputData.Slug, expectedData.Slug)
	assert.Equal(t, nil, actualErr)
}

func TestAppUpdate(t *testing.T) {
	expectedData := &app.AppData{
		Id:   os.Getenv("EXPO_TEST_UPDATE_APP_ID"),
		Name: generateRandomString(10),
		Slug: "test-app-update",
	}

	updateData := app.UpdateAppData{
		Id:   expectedData.Id,
		Name: expectedData.Name,
	}

	actualData, actualErr := client.App.Update(updateData)

	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}

// AppVariable

func TestAppEnvironmentVariableGet(t *testing.T) {
	expectedData := &appvariable.AppVariableData{
		Id:           os.Getenv("EXPO_TEST_GET_ENVIRONMENT_VARIABLE_ID"),
		Name:         "TEST_ENVIRONMENT_VARIABLE",
		Value:        "VALUE",
		Visibility:   "PUBLIC",
		Environments: []string{"PREVIEW"},
	}

	actualData, actualErr := client.AppVariable.Get(expectedData.Id, os.Getenv("EXPO_TEST_GET_APP_ID"))

	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}

func TestAppEnvironmentVariableGetByName(t *testing.T) {
	expectedData := &appvariable.AppVariableData{
		Id:           os.Getenv("EXPO_TEST_GET_ENVIRONMENT_VARIABLE_ID"),
		Name:         "TEST_ENVIRONMENT_VARIABLE",
		Value:        "VALUE",
		Visibility:   "PUBLIC",
		Environments: []string{"PREVIEW"},
	}

	actualData, actualErr := client.AppVariable.GetByName(expectedData.Name, os.Getenv("EXPO_TEST_GET_APP_ID"))

	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}

func TestAppEnvironmentVariableUpdate(t *testing.T) {
	updateData := appvariable.UpdateAppVariableData{
		Id:           os.Getenv("EXPO_TEST_UPDATE_ENVIRONMENT_VARIABLE_ID"),
		Name:         generateRandomString(10),
		Value:        generateRandomString(10),
		Visibility:   "PUBLIC",
		Environments: []string{"PRODUCTION"},
	}

	actualData, actualErr := client.AppVariable.Update(updateData)

	assert.Equal(t, actualData.Id, updateData.Id)
	assert.Equal(t, actualData.Name, updateData.Name)
	assert.Equal(t, actualData.Value, updateData.Value)
	assert.Equal(t, actualData.Visibility, updateData.Visibility)
	assert.Equal(t, actualData.Environments, updateData.Environments)
	assert.Equal(t, nil, actualErr)
}

func TestAppEnvironmentVariableCreateAndDelete(t *testing.T) {
	inputData := appvariable.CreateAppVariableData{
		AppId:        os.Getenv("EXPO_TEST_GET_APP_ID"),
		Name:         generateRandomString(10),
		Value:        generateRandomString(10),
		Visibility:   "PUBLIC",
		Environments: []string{"DEVELOPMENT", "PREVIEW"},
	}

	actualData, actualErr := client.AppVariable.Create(inputData)

	assert.Equal(t, inputData.Name, actualData.Name)
	assert.Equal(t, inputData.Value, actualData.Value)
	assert.Equal(t, inputData.Visibility, actualData.Visibility)
	assert.Equal(t, inputData.Environments, actualData.Environments)
	assert.Equal(t, nil, actualErr)

	// Delete
	actualErr = client.AppVariable.Delete(actualData.Id)

	assert.Equal(t, nil, actualErr)
}

const charset = "abcdefghijklmnopqrstuvwxyz"

func generateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
