package appleteam

type updateAppleTeam struct {
	Data *appleTeamData `json:"updateAppleTeam"`
}

type updateAppleTeamResponse struct {
	UpdateAppleTeam updateAppleTeam `json:"appleTeam"`
}

const updateAppleTeamMutation = `
	mutation ($id: ID!, $name: String!, $type: AppleTeamType!) {
		appleTeam {
			updateAppleTeam(
				appleTeamUpdateInput: { appleTeamName: $name, appleTeamType: $type }
				id: $id
			) {
				appleTeamName
				appleTeamType
				id
				appleTeamIdentifier
				account {
					id
				}
			}
		}
	}`

func (service *appleTeamService) Update(data UpdateAppleTeamData) (*AppleTeamData, error) {
	variables := map[string]any{
		"id":   data.Id,
		"name": data.Name,
		"type": data.Type,
	}

	var response updateAppleTeamResponse

	err := service.graphql.Query(updateAppleTeamMutation, variables, &response)

	if err != nil {
		return nil, err
	}

	return &AppleTeamData{
		Id:         response.UpdateAppleTeam.Data.Id,
		Name:       response.UpdateAppleTeam.Data.Name,
		Identifier: response.UpdateAppleTeam.Data.Identifier,
		Type:       response.UpdateAppleTeam.Data.Type,
		AccountId:  response.UpdateAppleTeam.Data.Account.Id,
	}, nil
}
