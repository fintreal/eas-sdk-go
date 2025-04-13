package appcredentials

import "fmt"

type appByAppId struct {
	Data []data `json:"androidAppCredentials"`
}

type getResponse struct {
	AppByAppId appByAppId `json:"appByAppId"`
}

const getQuery = `
	query ($appId: String!) {
    appByAppId(appId: $appId) {
      androidAppCredentials {
        app { id }
        applicationIdentifier
        googleServiceAccountKeyForSubmissions { id }
        id
      }
    }
	}
`

func (service *service) Get(data GetData) (*Data, error) {
	variables := map[string]any{
		"appId": data.AppId,
	}

	var response getResponse

	err := service.graphql.Query(getQuery, variables, &response)

	if err != nil {
		return nil, err
	}

	if len(response.AppByAppId.Data) == 0 {
		return nil, nil
	}

	return findById(response.AppByAppId.Data, data.Id)
}

func findById(data []data, id string) (*Data, error) {
	for _, d := range data {
		if d.Id == id {
			data := &Data{
				Id:                        d.Id,
				AppId:                     d.App.Id,
				Identifier:                d.Identifier,
				GoogleServiceAccountKeyId: d.GoogleServiceAccountKey.Id,
			}
			return data, nil
		}
	}
	return nil, fmt.Errorf("couldn't find app credentials with id %s", id)
}
