package test

import (
	"os"
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/api/appvariable"
	"github.com/stretchr/testify/assert"
)

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
