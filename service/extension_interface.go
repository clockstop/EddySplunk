package service

import "context"

// ExtensionServiceInterface defines the methods for interacting with the AWS Lambda Extensions API
type ExtensionServiceInterface interface {
    Register(ctx context.Context, extensionName string) (*RegisterResponse, error)
}