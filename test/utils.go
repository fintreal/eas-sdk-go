package test

import (
	"math/rand"
	"os"

	"github.com/fintreal/eas-sdk-go/eas"
)

var token = os.Getenv("EXPO_TOKEN")
var client = eas.NewEASClient(token)

const charset = "abcdefghijklmnopqrstuvwxyz"

func generateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
