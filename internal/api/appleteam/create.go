package appleteam

type createAppleTeam struct {
	Data *appleTeamData `json:"createAppleTeam"`
}

type createAppleTeamResponse struct {
	CreateAppleTeam createAppleTeam `json:"appleTeam"`
}

const createAppleTeamMutation = `
	mutation ($accountId: ID!, $name: String!, $identifier: String!, $type: AppleTeamType!) {
		appleTeam {
			createAppleTeam(
				accountId: $accountId
				appleTeamInput: {
					appleTeamName: $name
					appleTeamType: $type
					appleTeamIdentifier: $identifier
				}
			) {
				appleTeamIdentifier
				appleTeamName
				appleTeamType
				id
				account {
					id
				}
			}
		}
	}`

func (service *appleTeamService) Create(data CreateAppleTeamData) (*AppleTeamData, error) {
	variables := map[string]any{
		"identifier": data.Identifier,
		"name":       data.Name,
		"accountId":  data.AccountId,
		"type":       data.Type,
	}

	var response createAppleTeamResponse

	err := service.graphql.Query(createAppleTeamMutation, variables, &response)

	if err != nil {
		return nil, err
	}

	return &AppleTeamData{
		Id:         response.CreateAppleTeam.Data.Id,
		Name:       response.CreateAppleTeam.Data.Name,
		Identifier: response.CreateAppleTeam.Data.Identifier,
		Type:       response.CreateAppleTeam.Data.Type,
		AccountId:  response.CreateAppleTeam.Data.Account.Id,
	}, nil
}
