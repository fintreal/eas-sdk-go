package test

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/api/appleteam"
	"github.com/stretchr/testify/assert"
)

func TestGetAppleTeamByIdentifier(t *testing.T) {
	expectedData := &appleteam.AppleTeamData{
		Id:         immutableAppleTeamId,
		AccountId:  accountId,
		Identifier: "TEST_IDENTIFIER",
		Name:       "TEST_NAME",
		Type:       "COMPANY_OR_ORGANIZATION",
	}

	actualData, actualErr := client.AppleTeam.GetByIdentifier(expectedData.Identifier, expectedData.AccountId)

	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}

func TestCreateAppleTeamByIdentifier(t *testing.T) {
	expectedData := &appleteam.AppleTeamData{
		AccountId:  accountId,
		Identifier: generateRandomString(10),
		Name:       generateRandomString(10),
		Type:       "COMPANY_OR_ORGANIZATION", // RANDOMIZE
	}

	input := appleteam.CreateAppleTeamData{
		Identifier: expectedData.Identifier,
		Name:       expectedData.Name,
		Type:       expectedData.Type,
		AccountId:  accountId,
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
		Id:         mutableAppleTeamId,
		AccountId:  accountId,
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
