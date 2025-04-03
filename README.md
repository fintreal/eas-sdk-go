# expo-eas-sdk-go

Expo EAS SDK for Go.
- Manage Expo EAS app and environment variables from code
- Uses Expo EAS GraphQL API
- Builds on [machinebox/graphql](https://github.com/machinebox/graphql). Big Thanks!


## Installation
```
go get github.com/fintreal/expo-eas-sdk-go
```

## Usage

### Initialize
```go
client := eas.NewEASClient(token)

data, _ := client.Me.Get()

fmt.Printf("%+v\n", data)
```

### Account

##### Get By Name
```go
client := eas.NewEASClient(token)

data, _ := client.Account.GetByName(accountName)

fmt.Printf("%+v\n", data)
```

### App

##### Get
```go
client := NewEASClient(token)

data, err := client.App.Get(appId)

fmt.Printf("%+v", data)
```

##### Create
```go
client := eas.NewEASClient(token)

inputData := eas.CreateAppData{
	AccountId: accountId,
	Name:      "Test App Name",
	Slug:      "Test Slug",
}

data, _ := client.App.Create(inputData)

fmt.Printf("%+v\n", data)
```

##### Update
```go
client := eas.NewEASClient(token)

inputData := eas.UpdateAppData{
	AccountId: accountId,
	Name:      "Test App Name",
}

data, _ := client.App.Update(inputData)

fmt.Printf("%+v\n", data)
```

### App Environment Variable

##### Get
```go
client := eas.NewEASClient(token)
	
data, _ := client.AppVariable.Get(appVariableId, appId)

fmt.Printf("%+v\n", data)
```

##### Get By Name
```go
client := eas.NewEASClient(token)

data, _ := client.AppVariable.GetByName(appVariableName, appId)

fmt.Printf("%+v\n", data)
```

##### Create
```go
client := eas.NewEASClient(token)
inputData := eas.CreateAppVariableData{
	AppId:        appId,
	Name:         "API_URL",
	Value:        "http://localhost:3000/api",
	Visibility:   "PUBLIC",
	Environments: []string{"PREVIEW", "PRODUCTION"},
}

data, _ := client.AppVariable.Create(inputData)

fmt.Printf("%+v\n", data)
```

##### Update
```go
client := eas.NewEASClient(token)
inputData := eas.UpdateAppVariableData{
	Id:           id,
	Name:         "API_URL",
	Value:        "http://localhost:3000/api",
	Visibility:   "PUBLIC",
	Environments: []string{"DEVELOPMENT", "PREVIEW"},
}

data, _ := client.AppVariable.Create(inputData)

fmt.Printf("%+v\n", data)
```

##### Delete
```go
client := eas.NewEASClient(token)

data, _ := client.AppVariable.Delete(appVariableId)

fmt.Printf("%+v\n", data)
```
