package appvariable

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/testutils"
)

func TestDelete(t *testing.T) {
	id := "test-id"
	expectedVariables := map[string]any{"id": id}

	config := testutils.TestConfig[string, any, any, Service]{
		NewServiceFunction: NewService,
		FunctionUnderTest:  "Delete",
		Input:              &id,
		MockResponse:       mockResponse,
		ExpectedQuery:      deleteQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       nil,
	}
	testutils.Test(t, config)
}
