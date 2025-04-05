package appleappbundleidentifier

type createAppleAppBundleIdentifier struct {
	Data appleAppBundleIdentifierData `json:"createAppleAppIdentifier"`
}

type createAppleAppBundleIdentifierResponse struct {
	CreateAppleAppIdentifier createAppleAppBundleIdentifier `json:"appleAppIdentifier"`
}

const createAppleAppIdentifierMutation = `
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

func (s *appleAppBundleIdentifierService) Create(data CreateAppleAppBundleIdentifierData) (*AppleAppBundleIdentifierData, error) {
	variables := map[string]any{
		"accountId":   data.AccountId,
		"identifier":  data.Identifier,
		"teamId": data.TeamId,
	}

	var response createAppleAppBundleIdentifierResponse
	err := s.graphql.Query(createAppleAppIdentifierMutation, variables, &response)
	if err != nil {
		return nil, err
	}

	return &AppleAppBundleIdentifierData{
		Id:         response.CreateAppleAppIdentifier.Data.Id,
		Identifier: response.CreateAppleAppIdentifier.Data.Identifier,
		TeamId:     response.CreateAppleAppIdentifier.Data.AppleTeam.Id,
	}, nil
}
