package tenant_settings

import "errors"

var (
	ErrTenantSettingNotFound     = errors.New("tenant setting not found")
	ErrTenantIDRequired          = errors.New("tenant id is required")
	ErrTenantSettingAlreadyExists = errors.New("tenant setting already exists")
)
