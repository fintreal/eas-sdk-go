package appcredentials

import (
	"fmt"

	"github.com/fintreal/eas-sdk-go/internal/api/apple/appbuildcredentials"
)

type appByAppId struct {
	Data []data `json:"iosAppCredentials"`
}

type getResponse struct {
	AppByAppId appByAppId `json:"appByAppId"`
}

const getQuery = `
	query ($appId: String!) {
    appByAppId(appId: $appId) {
      iosAppCredentials {
        id
        appleAppIdentifier {
          id
        }
        app {
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

func (service *service) Get(data GetData) (*Data, error) {
	variables := map[string]any{
		"appId": data.AppId,
		"id":    data.Id,
	}

	var response getResponse

	err := service.graphql.Query(getQuery, variables, &response)

	if err != nil {
		return nil, err
	}

	return findById(response.AppByAppId.Data, data.Id)
}

func findById(data []data, id string) (*Data, error) {
	for _, d := range data {
		if d.Id == id {
			buildCredentials := []appbuildcredentials.Data{}
			for _, b := range d.BuildCredentials {
				buildCredential := appbuildcredentials.Data{
					Id:                    b.Id,
					AppCredentialsId:      b.AppCredentials.Id,
					CertificateId:         b.Certificate.Id,
					DistributionType:      b.DistributionType,
					ProvisioningProfileId: b.ProvisioningProfile.Id,
				}
				buildCredentials = append(buildCredentials, buildCredential)
			}

			return &Data{
				Id:               d.Id,
				AppId:            d.App.Id,
				AppIdentifierId:  d.AppIdentifier.Id,
				BuildCredentials: buildCredentials,
			}, nil
		}
	}
	return nil, fmt.Errorf("couldn't find app credentials with id %s", id)
}
