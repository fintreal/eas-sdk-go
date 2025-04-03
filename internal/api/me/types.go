package me

import (
	"github.com/fintreal/expo-eas-sdk-go/internal/graphql"
)

type MeData struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName"`
}

type MeService interface {
	Get() (*MeData, error)
}

type meService struct {
	graphql graphql.GraphQL
}

var _ MeService = (*meService)(nil)

func NewMeService(graphql graphql.GraphQL) MeService {
	return &meService{
		graphql: graphql,
	}
}
