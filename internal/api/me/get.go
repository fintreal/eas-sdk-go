package me

type getMeResponse struct {
	Data *MeData `json:"meActor"`
}

const meQuery = `
	query {
		meActor {
			displayName
			id
		}
	}`

// Retrieves the current user name and id from EAS
func (service *meService) Get() (*MeData, error) {
	variables := map[string]any{}

	var response getMeResponse
	err := service.graphql.Query(meQuery, variables, &response)
	return response.Data, err
}
