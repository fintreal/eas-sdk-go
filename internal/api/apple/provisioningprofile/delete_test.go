package provisioningprofile

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/graphql"
	"github.com/stretchr/testify/assert"
)

func TestDele(t *testing.T) {
	id := "test-id"
	expectedVariables := map[string]any{"id": id}

	graphQLMock := graphql.NewGraphQLMock[any](nil)

	service := NewProvisioningProfileService(graphQLMock)

	actualErr := service.Delete(id)

	actualQuery := graphQLMock.Calls[0].Arguments.Get(0).(string)
	actualVariables := graphQLMock.Calls[0].Arguments.Get(1).(map[string]any)

	assert.Equal(t, deleteQuery, actualQuery)
	assert.Equal(t, expectedVariables, actualVariables)
	assert.Equal(t, nil, actualErr)
}
