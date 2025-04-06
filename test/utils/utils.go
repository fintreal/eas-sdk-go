package utils

import (
	"math/rand"
	"os"

	"github.com/fintreal/eas-sdk-go/eas"
)

var Token = os.Getenv("EXPO_TOKEN")
var Client = eas.NewEASClient(Token)
var ImmutableProvisioningProfileBase64 = os.Getenv("IMMUTABLE_PROVISIONING_PROFILE_BASE64")

// TEST IDs

var AccountId = "c705db5b-7312-4d9e-baeb-65bb640b0888"
var MeId = "e013cf58-3bf8-4522-a4da-cde29b9fbc95"

var ImmutableAppId = "953ed82f-4ac7-47be-ab46-d9c7a1169fe6"
var MutableAppId = "976aa3ac-c3c7-47f5-a94e-2f3f0f75409d"

var ImmutableAppVariableId = "955ef61a-b78c-43cc-bfa7-6da1a51e76e3"
var MutableAppVariableId = "2d90f56c-aead-41d8-8a76-dcbcb92415ac"

var ImmutableAppleTeamId = "cf4830de-3c1b-4594-ad73-20ca9ac901b7"
var MutableAppleTeamId = "cc513f5d-366d-42d6-aa40-e8132c3c78a3"

const charset = "abcdefghijklmnopqrstuvwxyz"

func GenerateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
