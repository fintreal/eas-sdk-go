package bundleidentifier

import (
	"fmt"

	"github.com/fintreal/eas-sdk-go/internal/utils"
)

type getResponse struct {
	Data []data `json:"appleAppIdentifiers"`
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

func (s *service) GetByIdentifier(getData GetByIdentifierData) (*Data, error) {
	variables := map[string]any{
		"identifier": getData.Identifier,
		"accountId":  getData.AccountId,
	}

	var response utils.AccountResponse[getResponse]

	err := s.graphql.Query(getQuery, variables, &response)
	if err != nil {
		return nil, err
	}
	return findBundleIdentifierByIdentifier(response.Account.ById.Data, getData.Identifier)
}

func findBundleIdentifierByIdentifier(identifiers []data, identifier string) (*Data, error) {
	for _, bundleIdentifier := range identifiers {
		if bundleIdentifier.Identifier == identifier {
			return &Data{
				Id:         bundleIdentifier.Id,
				Identifier: bundleIdentifier.Identifier,
				TeamId:     bundleIdentifier.Team.Id,
			}, nil
		}
	}
	return nil, fmt.Errorf("couldn't find bundle identifier with identifier %s", identifier)
}
