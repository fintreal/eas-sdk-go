package provisioningprofile

import "fmt"

type byId struct {
	Data []provisioningProfileData `json:"appleDistributionCertificates"`
}

type getProvisioningProfiles struct {
	ById byId `json:"byId"`
}

type getProvisioningProfilesResponse struct {
	Account getProvisioningProfiles `json:"account"`
}

const getQuery = `
	query ($accountId: String!, $id: ID!) {
        account {
            byId(accountId: $accountId) {
                appleProvisioningProfiles(appleAppIdentifierId: $id) {
                    id
                    provisioningProfile
                    appleAppIdentifier {
                        id
                    }
                }
            }
        }
	}`

func (service *provisioningProfileService) Get(id string, accountId string) (*ProvisioningProfileData, error) {
	variables := map[string]any{
		"accountId": accountId,
	}

	var response getProvisioningProfilesResponse
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
