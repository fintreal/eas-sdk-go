package bundleidentifier

type createAppBundleIdentifier struct {
	Data appBundleIdentifierData `json:"createAppleAppIdentifier"`
}

type createAppBundleIdentifierResponse struct {
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

func (s *appBundleIdentifierService) Create(data CreateAppBundleIdentifierData) (*AppBundleIdentifierData, error) {
	variables := map[string]any{
		"accountId":  data.AccountId,
		"identifier": data.Identifier,
		"teamId":     data.TeamId,
	}

	var response createAppBundleIdentifierResponse
	err := s.graphql.Query(createQuery, variables, &response)
	if err != nil {
		return nil, err
	}

	return &AppBundleIdentifierData{
		Id:         response.CreateAppIdentifier.Data.Id,
		Identifier: response.CreateAppIdentifier.Data.Identifier,
		TeamId:     response.CreateAppIdentifier.Data.Team.Id,
	}, nil
}
