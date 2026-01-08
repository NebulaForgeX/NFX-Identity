package tenant_apps

import "errors"

var (
	ErrTenantAppNotFound      = errors.New("tenant app not found")
	ErrTenantIDRequired       = errors.New("tenant id is required")
	ErrAppIDRequired          = errors.New("app id is required")
	ErrTenantAppAlreadyExists = errors.New("tenant app already exists")
	ErrInvalidTenantAppStatus = errors.New("invalid tenant app status")
)
