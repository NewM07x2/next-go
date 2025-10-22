package resolver

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"next-go-sample/internal/service"
)

type Resolver struct {
	SampleService *service.SampleService
}

func NewResolver(sampleService *service.SampleService) *Resolver {
	return &Resolver{
		SampleService: sampleService,
	}
}
