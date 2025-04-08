package appvariable

import "fmt"

type getAppVariables struct {
	Data []AppVariableData `json:"environmentVariablesIncludingSensitive"`
}

type getAppVariablesResponse struct {
	AppByAppId getAppVariables `json:"appByAppId"`
}

const getAppVariablesQuery = `
	query ($appId: String!) {
		appByAppId(appId: $appId) {
			environmentVariablesIncludingSensitive {
				environments
				id
				name
				value
				visibility
			}
		}
	}`

func (service *appVariableService) getAppEnvironmentVariables(appId string) ([]AppVariableData, error) {
	variables := map[string]any{"appId": appId}

	var response getAppVariablesResponse

	err := service.graphql.Query(getAppVariablesQuery, variables, &response)
	return response.AppByAppId.Data, err
}

// Retrieves an App Environment Variable from EAS by it's name and appId
func (service *appVariableService) GetByName(getByName GetByNameAppVariableData) (*AppVariableData, error) {
	data, err := service.getAppEnvironmentVariables(getByName.AppId)
	if err != nil {
		return nil, err
	}
	return findByName(data, getByName.Name)
}

// Retrieves an App Environment Variable from EAS by it' id and appId
func (service *appVariableService) Get(getData GetAppVariableData) (*AppVariableData, error) {
	data, err := service.getAppEnvironmentVariables(getData.AppId)
	if err != nil {
		return nil, err
	}

	return findById(data, getData.Id)
}

func findByName(variables []AppVariableData, name string) (*AppVariableData, error) {
	for _, variable := range variables {
		if variable.Name == name {
			return &variable, nil
		}
	}
	return nil, fmt.Errorf("couldn't find variable with name %s", name)
}

func findById(variables []AppVariableData, id string) (*AppVariableData, error) {
	for _, variable := range variables {
		if variable.Id == id {
			return &variable, nil
		}
	}
	return nil, fmt.Errorf("couldn't find variable with id %s", id)
}
