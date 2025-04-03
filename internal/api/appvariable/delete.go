package appvariable

const deleteAppVariableMutation = `
	mutation ($id: ID!) {
		environmentVariable {
			deleteEnvironmentVariable(id: $id) {
				id
			}
		}
	}`

// Deletes an App Environment Variable from EAS
func (service *appVariableService) Delete(id string) error {
	variables := map[string]any{"id": id}

	var response any

	return service.graphql.Query(deleteAppVariableMutation, variables, &response)
}
