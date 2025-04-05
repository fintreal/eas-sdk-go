package provisioningprofile

import "github.com/fintreal/eas-sdk-go/internal/graphql"

type appleAppIdentifier struct {
	Id string `json:"id"`
}

type provisioningProfileData struct {
	Id                  string             `json:"id"`
	Base64              string             `json:"provisioningProfile"`
	AppBundleIdentifier appleAppIdentifier `json:"appleAppIdentifier"`
}

type ProvisioningProfileData struct {
	Id                    string
	Base64                string
	AppBundleIdentifierId string
}

type CreateProvisioningProfileData struct {
	Base64                string
	AppBundleIdentifierId string
	AccountId             string
}

type ProvisioningProfileService interface {
	Get(id string, accountId string) (*ProvisioningProfileData, error)
	Create(data CreateProvisioningProfileData) (*ProvisioningProfileData, error)
	Delete(id string) error
}

type provisioningProfileService struct {
	graphql graphql.GraphQL
}

var _ ProvisioningProfileService = (*provisioningProfileService)(nil)

func NewProvisioningProfileService(graphql graphql.GraphQL) ProvisioningProfileService {
	return &provisioningProfileService{
		graphql: graphql,
	}
}
