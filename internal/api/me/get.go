package me

type getResponse struct {
	Data *Data `json:"meActor"`
}

const getQuery = `
	query {
		meActor {
			displayName
			id
		}
	}`

// Retrieves the current user name and id from EAS
func (service *service) Get() (*Data, error) {
	variables := map[string]any{}

	var response getResponse
	err := service.graphql.Query(getQuery, variables, &response)
	return response.Data, err
}
