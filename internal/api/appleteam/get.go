package appleteam

type getAppleTeam struct {
	Data *appleTeamData `json:"byAppleTeamIdentifier"`
}

type getAppleTeamResponse struct {
	GetAppleTeam getAppleTeam `json:"appleTeam"`
}

const getAppleTeamByIdentifierQuery = `
	query ($identifier: String!, $accountId: ID!) {
		appleTeam {
			byAppleTeamIdentifier(accountId: $accountId, identifier: $identifier) {
				id
				appleTeamName
				appleTeamType
				appleTeamIdentifier
				account {
					id
				}
			}
		}
	}
`

func (service *appleTeamService) GetByIdentifier(identifier string, accountId string) (*AppleTeamData, error) {
	variables := map[string]any{
		"accountId":  accountId,
		"identifier": identifier,
	}

	var response getAppleTeamResponse

	err := service.graphql.Query(getAppleTeamByIdentifierQuery, variables, &response)

	if err != nil {
		return nil, err
	}

	return &AppleTeamData{
		Id:         response.GetAppleTeam.Data.Id,
		Name:       response.GetAppleTeam.Data.Name,
		Identifier: response.GetAppleTeam.Data.Identifier,
		Type:       response.GetAppleTeam.Data.Type,
		AccountId:  response.GetAppleTeam.Data.Account.Id,
	}, nil
}
