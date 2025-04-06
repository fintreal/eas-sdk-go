package test

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/api/app"
	"github.com/fintreal/eas-sdk-go/test/utils"
	"github.com/stretchr/testify/assert"
)

func TestAppGet(t *testing.T) {
	expectedData := &app.AppData{
		Id:   utils.ImmutableAppId,
		Name: "Test App",
		Slug: "test-app",
	}

	actualData, actualErr := utils.Client.App.Get(expectedData.Id)

	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}

func TestAppCreate(t *testing.T) {
	inputData := app.CreateAppData{
		AccountId: utils.AccountId,
		Name:      utils.GenerateRandomString(10),
		Slug:      utils.GenerateRandomString(10),
	}

	// Create
	expectedData, actualErr := utils.Client.App.Create(inputData)

	assert.Equal(t, inputData.Name, expectedData.Name)
	assert.Equal(t, inputData.Slug, expectedData.Slug)
	assert.Equal(t, nil, actualErr)
}

func TestAppUpdate(t *testing.T) {
	expectedData := &app.AppData{
		Id:   utils.MutableAppId,
		Name: utils.GenerateRandomString(10),
		Slug: "test-app-update",
	}

	updateData := app.UpdateAppData{
		Id:   expectedData.Id,
		Name: expectedData.Name,
	}

	actualData, actualErr := utils.Client.App.Update(updateData)

	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}
