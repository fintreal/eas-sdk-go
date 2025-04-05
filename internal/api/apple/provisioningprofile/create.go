package provisioningprofile

const createQuery = `
    mutation ($accountId: ID!, $appleAppIdentifierId: ID!, $base64: String!) {
        appleProvisioningProfile {
            createAppleProvisioningProfile(
                accountId: $accountId
                appleAppIdentifierId: $appleAppIdentifierId
                appleProvisioningProfileInput: { appleProvisioningProfile: $base64 }
            ) {
                id
                provisioningProfile
                appleAppIdentifier {
                    id
                }
            }
        }
    }
`

func (service *provisioningProfileService) Create(data CreateProvisioningProfileData) (*ProvisioningProfileData, error) {
	return nil, nil
}
