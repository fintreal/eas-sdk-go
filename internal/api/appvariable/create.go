package appvariable

type createAppVariable struct {
	Data *AppVariableData `json:"createEnvironmentVariableForApp"`
}

type createAppVariableResponse struct {
	CreateAppVariable createAppVariable `json:"environmentVariable"`
}

const createAppVariableMutation = `
	mutation ($appId: ID!, $name: String!, $value: String!, $visibility: EnvironmentVariableVisibility!, $environments: [EnvironmentVariableEnvironment!]) {
		environmentVariable {
			createEnvironmentVariableForApp(
				appId: $appId
				environmentVariableData: {
					name: $name
					value: $value
					type: STRING
					visibility: $visibility
					environments: $environments
				}
			) {
				environments
				id
				name
				scope
				type
				value
				visibility
			}
		}
	}`

// Creates an App Environment Variable in EAS
func (service *appVariableService) Create(data CreateAppVariableData) (*AppVariableData, error) {
	variables := map[string]any{
		"appId":        data.AppId,
		"name":         data.Name,
		"value":        data.Value,
		"visibility":   data.Visibility,
		"environments": data.Environments,
	}

	var response createAppVariableResponse
	err := service.graphql.Query(createAppVariableMutation, variables, &response)

	if err != nil {
		return nil, err
	}

	return response.CreateAppVariable.Data, err
}
