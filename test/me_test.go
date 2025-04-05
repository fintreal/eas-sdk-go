package test

import (
	"os"
	"testing"

	"github.com/fintreal/eas-sdk-go/internal/api/me"
	"github.com/stretchr/testify/assert"
)

func TestGetMe(t *testing.T) {
	expectedData := &me.MeData{
		Id:          os.Getenv("EXPO_ME_ID"),
		DisplayName: "integration-test",
	}

	actualData, actualErr := client.Me.Get()
	assert.Equal(t, expectedData, actualData)
	assert.Equal(t, nil, actualErr)
}
