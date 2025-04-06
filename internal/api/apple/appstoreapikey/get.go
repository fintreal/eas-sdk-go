package appstoreapikey

import (
	"fmt"

	"github.com/fintreal/eas-sdk-go/internal/utils"
)

type appStoreApiKeysResponse struct {
	Data []AppStoreApiKeyData `json:"appStoreConnectApiKeys"`
}

const getQuery = `
	query ($accountId: String!) {
			account {
					byId(accountId: $accountId) {
							appStoreConnectApiKeys {
									id
									issuerIdentifier
									keyIdentifier
									name
							}
					}
			}
	}
`

func (s *appStoreApiKeyService) GetByIdentifier(identifier string, accountId string) (*AppStoreApiKeyData, error) {
	variables := map[string]any{"accountId": accountId}

	var response utils.AccountResponse[appStoreApiKeysResponse]

	err := s.graphql.Query(getQuery, variables, &response)
	if err != nil {
		return nil, err
	}

	return findAppStoreApiKeyByIdentifier(response.Account.ById.Data, identifier)
}

func findAppStoreApiKeyByIdentifier(apiKeys []AppStoreApiKeyData, identifier string) (*AppStoreApiKeyData, error) {
	for _, apiKey := range apiKeys {
		if apiKey.Identifier == identifier {
			return &apiKey, nil
		}
	}
	return nil, fmt.Errorf("couldn't find app store api key with identifier %s", identifier)
}
