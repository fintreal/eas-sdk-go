package provisioningprofile

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
	return nil, nil
}
