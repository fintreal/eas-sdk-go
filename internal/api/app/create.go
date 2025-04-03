package app

type createApp struct {
	Data *AppData `json:"createApp"`
}

type createAppResponse struct {
	CreateApp createApp `json:"app"`
}

const createAppMutation = `
	mutation ($accountId: ID!, $name: String!, $slug: String!) {
		app {
			createApp(
				appInput: {
					accountId: $accountId,
					appInfo: { displayName: $name }
					projectName: $slug
				}
			) {
				id
				name
				slug
			}
		}
	}`

// Creates an App in EAS
func (service *appService) Create(data CreateAppData) (*AppData, error) {
	variables := map[string]any{
		"accountId": data.AccountId,
		"name":      data.Name,
		"slug":      data.Slug,
	}

	var response createAppResponse
	err := service.graphql.Query(createAppMutation, variables, &response)

	if err != nil {
		return nil, err
	}

	return response.CreateApp.Data, nil
}
