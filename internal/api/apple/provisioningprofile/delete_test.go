package provisioningprofile

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/testutils"
)

func TestDele(t *testing.T) {
	id := "test-id"
	expectedVariables := map[string]any{"id": id}

	config := testutils.TestConfig[string, any, any, ProvisioningProfileService]{
		NewServiceFunction: NewProvisioningProfileService,
		FunctionUnderTest:  "Delete",
		Input:              &id,
		MockResponse:       nil,
		ExpectedQuery:      deleteQuery,
		ExpectedVariables:  expectedVariables,
		ExpectedData:       nil,
	}
	testutils.Test(t, config)
}
