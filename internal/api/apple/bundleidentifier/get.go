package bundleidentifier

import "fmt"

type byId struct {
	Data []appBundleIdentifierData `json:"appleAppIdentifiers"`
}

type getBundleIdentifiers struct {
	ById byId `json:"byId"`
}

type getBundleIdentifiersResponse struct {
	Account getBundleIdentifiers `json:"account"`
}

const getQuery = `
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

func (s *appBundleIdentifierService) GetByIdentifier(identifier string, accountId string) (*AppBundleIdentifierData, error) {
	variables := map[string]any{
		"identifier": identifier,
		"accountId":  accountId,
	}

	var response getBundleIdentifiersResponse

	err := s.graphql.Query(getQuery, variables, &response)
	if err != nil {
		return nil, err
	}
	return findBundleIdentifierByIdentifier(response.Account.ById.Data, identifier)
}

func findBundleIdentifierByIdentifier(identifiers []appBundleIdentifierData, identifier string) (*AppBundleIdentifierData, error) {
	for _, bundleIdentifier := range identifiers {
		if bundleIdentifier.Identifier == identifier {
			return &AppBundleIdentifierData{
				Id:         bundleIdentifier.Id,
				Identifier: bundleIdentifier.Identifier,
				TeamId:     bundleIdentifier.Team.Id,
			}, nil
		}
	}
	return nil, fmt.Errorf("couldn't find bundle identifier with identifier %s", identifier)
}
