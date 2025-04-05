package provisioningprofile

const deleteQuery = `
	mutation ($id: ID!) {
    appleProvisioningProfile {
      deleteAppleProvisioningProfile(id: $id) {
        id
      }
    }
	}`

func (service *provisioningProfileService) Delete(id string) error {
	variables := map[string]any{"id": id}

	var response any

	return service.graphql.Query(deleteQuery, variables, &response)
}
