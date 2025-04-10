package appcredentials

import "github.com/fintreal/eas-sdk-go/internal/api/apple/appbuildcredentials"

type createIosAppCredentials struct {
	Data data `json:"createIosAppCredentials"`
}

type createResponse struct {
	IosAppCredentials createIosAppCredentials `json:"iosAppCredentials"`
}

const createQuery = `
mutation ($appId: ID!, $appIdentifierId: ID!) {
  iosAppCredentials {
    createIosAppCredentials(
      appId: $appId
      appleAppIdentifierId: $appIdentifierId
      iosAppCredentialsInput: {  }
    ) {
      id
      app {
        id
      }
      appleAppIdentifier {
        id
      }
      iosAppBuildCredentialsArray {
      	id
      	iosDistributionType
        provisioningProfile {
          id
        }
        distributionCertificate {
          id
        }
        iosAppCredentials {
          id
        }
      }
    }
  }
}
`

func (service *service) Create(data CreateData) (*Data, error) {
	variables := map[string]any{
		"appId":           data.AppId,
		"appIdentifierId": data.AppIdentifierId,
	}

	var response createResponse

	err := service.graphql.Query(createQuery, variables, &response)

	if err != nil {
		return nil, err
	}

	buildCredentials := []appbuildcredentials.Data{}

	for _, b := range response.IosAppCredentials.Data.BuildCredentials {
		buildCredential := appbuildcredentials.Data{
			Id:                    b.Id,
			DistributionType:      b.DistributionType,
			AppCredentialsId:      b.AppCredentials.Id,
			ProvisioningProfileId: b.ProvisioningProfile.Id,
			CertificateId:         b.Certificate.Id,
		}
		buildCredentials = append(buildCredentials, buildCredential)
	}

	return &Data{
		Id:               response.IosAppCredentials.Data.Id,
		AppId:            response.IosAppCredentials.Data.App.Id,
		AppIdentifierId:  response.IosAppCredentials.Data.AppIdentifier.Id,
		BuildCredentials: buildCredentials,
	}, nil
}
