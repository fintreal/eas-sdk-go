package bundleidentifier

type createAppBundleIdentifier struct {
	Data data `json:"createAppleAppIdentifier"`
}

type createResponse struct {
	CreateAppIdentifier createAppBundleIdentifier `json:"appleAppIdentifier"`
}

const createQuery = `
	mutation ($accountId: ID!, $identifier: String!, $teamId: ID!) {
			appleAppIdentifier {
					createAppleAppIdentifier(
							accountId: $accountId
							appleAppIdentifierInput: { bundleIdentifier: $identifier, appleTeamId: $teamId }
					) {
							bundleIdentifier
							id
							appleTeam {
									id
							}
					}
			}
	}`

func (s *service) Create(data CreateData) (*Data, error) {
	variables := map[string]any{
		"accountId":  data.AccountId,
		"identifier": data.Identifier,
		"teamId":     data.TeamId,
	}

	var response createResponse
	err := s.graphql.Query(createQuery, variables, &response)
	if err != nil {
		return nil, err
	}

	return &Data{
		Id:         response.CreateAppIdentifier.Data.Id,
		Identifier: response.CreateAppIdentifier.Data.Identifier,
		TeamId:     response.CreateAppIdentifier.Data.Team.Id,
	}, nil
}
