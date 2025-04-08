package team

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/testutils"
)

func TestCreate(t *testing.T) {
	expectedData := &TeamData{
		Id:         "test-id",
		Name:       "test-name",
		Identifier: "test-identifier",
		Type:       "test-type",
		AccountId:  "test-account-id",
	}

	expectedVariables := map[string]any{
		"identifier": expectedData.Identifier,
		"name":       expectedData.Name,
		"accountId":  expectedData.AccountId,
		"type":       expectedData.Type,
	}

	input := &CreateTeamData{
		Name:       expectedData.Name,
		Identifier: expectedData.Identifier,
		Type:       expectedData.Type,
		AccountId:  expectedData.AccountId,
	}

	mockResponse := createTeamResponse{
		CreateTeam: createTeam{
			Data: &teamData{
				Id:         expectedData.Id,
				Name:       expectedData.Name,
				Identifier: expectedData.Identifier,
				Type:       expectedData.Type,
				Account:    account{Id: expectedData.AccountId},
			},
		},
	}

	config := testutils.TestConfig[CreateTeamData, TeamData, createTeamResponse, TeamService]{
		NewServiceFunction: NewTeamService,
		FunctionUnderTest:  "Create",
		Input:              input,
		MockResponse:       mockResponse,
		ExpectedQuery:      createQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       expectedData,
	}
	testutils.Test(t, config)
}
