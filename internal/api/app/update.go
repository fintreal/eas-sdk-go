package app

type updateApp struct {
	Data *AppData `json:"setAppInfo"`
}

type updateAppResponse struct {
	UpdateApp updateApp `json:"app"`
}

const updateAppMutation = `
	mutation ($id: ID!, $name: String!) {
		app {
			setAppInfo(appId: $id, appInfo: { displayName: $name }) {
				name
				slug
				id
			}
		}
	}`

// Creates an App in EAS
func (service *appService) Update(data UpdateAppData) (*AppData, error) {
	variables := map[string]any{
		"id":   data.Id,
		"name": data.Name,
	}

	var response updateAppResponse
	err := service.graphql.Query(updateAppMutation, variables, &response)

	if err != nil {
		return nil, err
	}

	return response.UpdateApp.Data, nil
}
