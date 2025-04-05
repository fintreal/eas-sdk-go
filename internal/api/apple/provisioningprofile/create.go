package provisioningprofile

type createProvisioningProfile struct {
	Data provisioningProfileData `json:"createAppleProvisioningProfile"`
}

type createProvisioningProfileResponse struct {
	CreateProvisioningProfile createProvisioningProfile `json:"appleProvisioningProfile"`
}

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
	variables := map[string]any{
		"accountId":            data.AccountId,
		"appleAppIdentifierId": data.AppBundleIdentifierId,
		"base64":               data.Base64,
	}

	var response createProvisioningProfileResponse

	err := service.graphql.Query(createQuery, variables, &response)

	if err != nil {
		return nil, err
	}

	return &ProvisioningProfileData{
		Id:                    response.CreateProvisioningProfile.Data.Id,
		AppBundleIdentifierId: response.CreateProvisioningProfile.Data.AppBundleIdentifier.Id,
		Base64:                response.CreateProvisioningProfile.Data.Base64,
	}, nil
}
