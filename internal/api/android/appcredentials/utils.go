package appcredentials

import "github.com/fintreal/eas-sdk-go/internal/api/android/appbuildcredentials"

func mapData(data data) Data {
	buildCredentials := []appbuildcredentials.Data{}
	for _, b := range data.BuildCredentials {
		buildCredential := appbuildcredentials.Data{
			Id:               b.Id,
			Name:             b.Name,
			AppCredentialsId: data.Id,
			KeystoreId:       b.Keystore.Id,
		}
		buildCredentials = append(buildCredentials, buildCredential)
	}
	return Data{
		Id:                        data.Id,
		AppId:                     data.App.Id,
		Identifier:                data.Identifier,
		GoogleServiceAccountKeyId: data.GoogleServiceAccountKey.Id,
		BuildCredentials:          buildCredentials,
	}
}
