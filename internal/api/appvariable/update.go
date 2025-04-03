package appvariable

type updateAppVariable struct {
	Data *AppVariableData `json:"updateEnvironmentVariable"`
}

type upddateAppVariableResponse struct {
	UpdateAppVariable updateAppVariable `json:"environmentVariable"`
}

const updateAppVariableMutation = `
	mutation ($id: ID!, $name: String!, $value: String!, $visibility: EnvironmentVariableVisibility!, $environments: [EnvironmentVariableEnvironment!]) {
		environmentVariable {
			updateEnvironmentVariable(
				environmentVariableData: { environments: $environments, id: $id, name: $name, value: $value, visibility: $visibility }
			) {
				environments
				id
				name
				value
				visibility
			}
		}
	}`

// Updates an App Environment Variable in EAS
func (service *appVariableService) Update(data UpdateAppVariableData) (*AppVariableData, error) {
	variables := map[string]any{
		"id":           data.Id,
		"name":         data.Name,
		"value":        data.Value,
		"visibility":   data.Visibility,
		"environments": data.Environments,
	}

	var response upddateAppVariableResponse
	err := service.graphql.Query(updateAppVariableMutation, variables, &response)

	if err != nil {
		return nil, err
	}

	return response.UpdateAppVariable.Data, err
}
