package appcredentials

type GetIosAppCredentials struct {
	Data data `json:"getIosAppCredentials"`
}

type getResponse struct {
	IosAppCredentials GetIosAppCredentials `json:"iosAppCredentials"`
}

const getQuery = `
	query ($appId: String!) {
    appByAppId(appId: $appId) {
      iosAppCredentials {
        id
        appleAppIdentifier {
          id
        }
      }
      id
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

	return &Data{
		Id:              response.IosAppCredentials.Data.Id,
		AppId:           response.IosAppCredentials.Data.App.Id,
		AppIdentifierId: response.IosAppCredentials.Data.AppIdentifier.Id,
	}, nil
}
