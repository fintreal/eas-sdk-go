package appcredentials

type createAndroidAppCredentials struct {
	Data data `json:"createAndroidAppCredentials"`
}

type createResponse struct {
	AndroidAppCredentials createAndroidAppCredentials `json:"androidAppCredentials"`
}

const createQuery = `
	mutation ($appId: ID!, $identifier: String!, $googleServiceAccountKeyId: ID!) {
    androidAppCredentials {
      createAndroidAppCredentials(
        appId: $appId
        applicationIdentifier: $identifier
        androidAppCredentialsInput: { googleServiceAccountKeyForSubmissionsId: $googleServiceAccountKeyId }
      ) {
        id
        applicationIdentifier
        googleServiceAccountKeyForSubmissions { id }
        app { id }
      }
    }
	}
`

func (service *service) Create(data CreateData) (*Data, error) {
	variables := map[string]any{
		"appId":                     data.AppId,
		"identifier":                data.Identifier,
		"googleServiceAccountKeyId": data.GoogleServiceAccountKeyId,
	}

	var response createResponse

	err := service.graphql.Query(createQuery, variables, &response)

	if err != nil {
		return nil, err
	}
	return &Data{
		Id:                        response.AndroidAppCredentials.Data.Id,
		AppId:                     response.AndroidAppCredentials.Data.App.Id,
		Identifier:                response.AndroidAppCredentials.Data.Identifier,
		GoogleServiceAccountKeyId: response.AndroidAppCredentials.Data.GoogleServiceAccountKey.Id,
	}, nil
}
