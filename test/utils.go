package test

import (
	"math/rand"
	"os"

	"github.com/fintreal/eas-sdk-go/eas"
)

var token = os.Getenv("EXPO_TOKEN")
var client = eas.NewEASClient(token)
var immutableProvisioningProfileBase64 = os.Getenv("IMMUTABLE_PROVISIONING_PROFILE_BASE64")

// TEST IDs

var accountId = "c705db5b-7312-4d9e-baeb-65bb640b0888"
var meId = "e013cf58-3bf8-4522-a4da-cde29b9fbc95"

var immutableAppId = "953ed82f-4ac7-47be-ab46-d9c7a1169fe6"
var mutableAppId = "976aa3ac-c3c7-47f5-a94e-2f3f0f75409d"

var immutableAppVariableId = "955ef61a-b78c-43cc-bfa7-6da1a51e76e3"
var mutableAppVariableId = "2d90f56c-aead-41d8-8a76-dcbcb92415ac"

var immutableAppleTeamId = "cf4830de-3c1b-4594-ad73-20ca9ac901b7"
var mutableAppleTeamId = "cc513f5d-366d-42d6-aa40-e8132c3c78a3"

const charset = "abcdefghijklmnopqrstuvwxyz"

func generateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
