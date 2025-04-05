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
	return nil
}
