package app

type getResponse struct {
	Data *Data `json:"appByAppId"`
}

const getQuery = `
	query App ($id: String!) {
    	appByAppId(appId: $id) {
        	id
       		name
        	slug
    	}
	}`

// Retrieves an App from EAS by it's id
func (service *service) Get(id string) (*Data, error) {
	variables := map[string]any{"id": id}

	var response getResponse
	err := service.graphql.Query(getQuery, variables, &response)

	return response.Data, err
}
