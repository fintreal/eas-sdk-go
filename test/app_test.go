package test

import (
	"os"
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/api/app"
	"github.com/stretchr/testify/assert"
)

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
