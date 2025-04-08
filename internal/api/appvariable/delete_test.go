package appvariable

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/testutils"
)

func TestDelete(t *testing.T) {
	id := "test-id"
	expectedVariables := map[string]any{"id": id}

	config := testutils.TestConfig[string, any, any, AppVariableService]{
		NewServiceFunction: NewAppVariableService,
		FunctionUnderTest:  "Delete",
		Input:              &id,
		MockResponse:       mockResponse,
		ExpectedQuery:      deleteAppVariableMutation,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       nil,
	}
	testutils.Test(t, config)
}
