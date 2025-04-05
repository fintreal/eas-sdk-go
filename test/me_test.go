package test

import (
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/api/me"
	"github.com/stretchr/testify/assert"
)

func TestGetMe(t *testing.T) {
	expectedData := &me.MeData{
		Id:          meId,
		DisplayName: "integration-test",
	}

	actualData, actualErr := client.Me.Get()
	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}
