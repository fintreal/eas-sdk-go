package bundleidentifier

import (
	"fmt"

	"github.com/fintreal/eas-sdk-go/internal/utils"
)

type appBundleIdentifierResponse struct {
	Data []appBundleIdentifierData `json:"appleAppIdentifiers"`
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

	var response utils.AccountResponse[appBundleIdentifierResponse]

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
