package appcredentials

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

	return &Data{
		Id:               response.IosAppCredentials.Data.Id,
		AppId:            response.IosAppCredentials.Data.App.Id,
		AppIdentifierId:  response.IosAppCredentials.Data.AppIdentifier.Id,
		BuildCredentials: response.IosAppCredentials.Data.BuildCredentials,
	}, nil
}
