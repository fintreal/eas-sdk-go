package appleappbundleidentifier

import "fmt"

type byId struct {
	Data []appleAppBundleIdentifierData `json:"appleAppIdentifiers"`
}

type getBundleIdentifiers struct {
	ById byId `json:"byId"`
}

type getBundleIdentifiersResponse struct {
	Account getBundleIdentifiers `json:"account"`
}

const getAppleTeamByIdentifierQuery = `
	query ($accountId: String!, $identifier: String!) {
		account {
			byId(accountId: $accountId) {
				appleAppIdentifiers(bundleIdentifier: $identifier) {
					bundleIdentifier
					id
					appleTeam {
							id
					}
				}
			}
		}
	}`

func (s *appleAppBundleIdentifierService) GetByIdentifier(identifier string, accountId string) (*AppleAppBundleIdentifierData, error) {
	variables := map[string]any{
		"identifier": identifier,
		"accountId":  accountId,
	}

	var response getBundleIdentifiersResponse

	err := s.graphql.Query(getAppleTeamByIdentifierQuery, variables, &response)
	if err != nil {
		return nil, err
	}
	return findBundleIdentifierByIdentifier(response.Account.ById.Data, identifier)
}

func findBundleIdentifierByIdentifier(identifiers []appleAppBundleIdentifierData, identifier string) (*AppleAppBundleIdentifierData, error) {
	for _, bundleIdentifier := range identifiers {
		if bundleIdentifier.Identifier == identifier {
			return &AppleAppBundleIdentifierData{
				Id:         bundleIdentifier.Id,
				Identifier: bundleIdentifier.Identifier,
				TeamId:     bundleIdentifier.AppleTeam.Id,
			}, nil
		}
	}
	return nil, fmt.Errorf("couldn't find bundle identifier with identifier %s", identifier)
}
