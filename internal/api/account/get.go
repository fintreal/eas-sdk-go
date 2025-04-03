package account

type getAccount struct {
	ByName *AccountData `json:"byName"`
}

type getAccountResponse struct {
	Account getAccount `json:"account"`
}

const getAccountByNameQuery = `
	query ($name: String!) {
		account {
			byName(accountName: $name) {
				id
				name
			}
		}
	}`

// Retrieves the currents user's organization account by it's name
//
// @name organization name
func (service *accountService) GetByName(name string) (*AccountData, error) {
	variables := map[string]any{"name": name}

	var response getAccountResponse

	err := service.graphql.Query(getAccountByNameQuery, variables, &response)

	return response.Account.ByName, err
}
