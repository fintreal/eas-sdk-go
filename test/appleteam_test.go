package test

import (
	"os"
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/api/appleteam"
	"github.com/stretchr/testify/assert"
)

func TestGetAppleTeamByIdentifier(t *testing.T) {
	expectedData := &appleteam.AppleTeamData{
		Id:         os.Getenv("EXPO_TEST_GET_APPLE_TEAM_ID"),
		AccountId:  os.Getenv("EXPO_ACCOUNT_ID"),
		Identifier: os.Getenv("EXPO_TEST_GET_APPLE_TEAM_IDENTIFIER"),
		Name:       os.Getenv("EXPO_TEST_GET_APPLE_TEAM_NAME"),
		Type:       "COMPANY_OR_ORGANIZATION",
	}

	actualData, actualErr := client.AppleTeam.GetByIdentifier(expectedData.Identifier, expectedData.AccountId)

	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}

func TestCreateAppleTeamByIdentifier(t *testing.T) {
	expectedData := &appleteam.AppleTeamData{
		AccountId:  os.Getenv("EXPO_ACCOUNT_ID"),
		Identifier: generateRandomString(10),
		Name:       generateRandomString(10),
		Type:       "COMPANY_OR_ORGANIZATION", // RANDOMIZE
	}

	input := appleteam.CreateAppleTeamData{
		Identifier: expectedData.Identifier,
		Name:       expectedData.Name,
		Type:       expectedData.Type,
		AccountId:  os.Getenv("EXPO_ACCOUNT_ID"),
	}

	actualData, actualErr := client.AppleTeam.Create(input)

	assert.Equal(t, expectedData.AccountId, actualData.AccountId)
	assert.Equal(t, expectedData.Identifier, actualData.Identifier)
	assert.Equal(t, expectedData.Name, actualData.Name)
	assert.Equal(t, expectedData.Type, actualData.Type)
	assert.Equal(t, nil, actualErr)
}

func TestUpdateAppleTeamByIdentifier(t *testing.T) {
	expectedData := &appleteam.AppleTeamData{
		Id:         "cc513f5d-366d-42d6-aa40-e8132c3c78a3",
		AccountId:  os.Getenv("EXPO_ACCOUNT_ID"),
		Identifier: "TEST_UPDATE_TEAM",
		Name:       generateRandomString(10),
		Type:       "COMPANY_OR_ORGANIZATION", // RANDOMIZE
	}

	input := appleteam.UpdateAppleTeamData{
		Id:   expectedData.Id,
		Name: expectedData.Name,
		Type: expectedData.Type,
	}

	actualData, actualErr := client.AppleTeam.Update(input)

	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}
