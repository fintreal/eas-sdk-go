package app

type getAppResponse struct {
	Data *AppData `json:"appByAppId"`
}

const getAppQuery = `
	query App ($id: String!) {
    	appByAppId(appId: $id) {
        	id
       		name
        	slug
    	}
	}`

// Retrieves an App from EAS by it's id
func (service *appService) Get(id string) (*AppData, error) {
	variables := map[string]any{"id": id}

	var response getAppResponse
	err := service.graphql.Query(getAppQuery, variables, &response)

	return response.Data, err
}
