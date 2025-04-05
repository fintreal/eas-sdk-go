package provisioningprofile

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/graphql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDele(t *testing.T) {
	id := "test-id"
	expectedVariables := map[string]any{"id": id}

	graphQLMock := newDeleteGraphQLMock()

	service := NewProvisioningProfileService(graphQLMock)

	actualErr := service.Delete(id)

	actualQuery := graphQLMock.Calls[0].Arguments.Get(0).(string)
	actualVariables := graphQLMock.Calls[0].Arguments.Get(1).(map[string]any)

	assert.Equal(t, deleteQuery, actualQuery)
	assert.Equal(t, expectedVariables, actualVariables)
	assert.Equal(t, nil, actualErr)
}

func newDeleteGraphQLMock() *graphql.GraphQLMock {
	graphQLMock := graphql.NewGraphQLMock()
	graphQLMock.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	return graphQLMock
}
