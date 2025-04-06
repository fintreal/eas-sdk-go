package provisioningprofile

import (
	"fmt"

	"github.com/fintreal/eas-sdk-go/internal/utils"
)

type getProvisioningProfilesResponse struct {
	Data []provisioningProfileData `json:"appleProvisioningProfiles"`
}

const getQuery = `
	query ($accountId: String!) {
			account {
					byId(accountId: $accountId) {
							appleProvisioningProfiles {
									id
									provisioningProfile
									appleAppIdentifier {
											id
									}
							}
					}
			}
	}
`

func (service *provisioningProfileService) Get(id string, accountId string) (*ProvisioningProfileData, error) {
	variables := map[string]any{
		"accountId": accountId,
	}

	var response utils.AccountResponse[getProvisioningProfilesResponse]
	err := service.graphql.Query(getQuery, variables, &response)
	if err != nil {
		return nil, err
	}
	return findProvisioningProfileById(response.Account.ById.Data, id)
}

func findProvisioningProfileById(provisioningProfiles []provisioningProfileData, id string) (*ProvisioningProfileData, error) {
	for _, provisioningProfile := range provisioningProfiles {
		if provisioningProfile.Id == id {
			return &ProvisioningProfileData{
				Id:                    provisioningProfile.Id,
				AppBundleIdentifierId: provisioningProfile.AppBundleIdentifier.Id,
				Base64:                provisioningProfile.Base64,
			}, nil
		}
	}
	return nil, fmt.Errorf("couldn't find provisioning profile with id %s", id)
}
