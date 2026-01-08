package tenants

import "errors"

var (
	ErrTenantNotFound      = errors.New("tenant not found")
	ErrTenantIDRequired    = errors.New("tenant id is required")
	ErrNameRequired        = errors.New("name is required")
	ErrTenantIDAlreadyExists = errors.New("tenant id already exists")
	ErrInvalidTenantStatus = errors.New("invalid tenant status")
)
