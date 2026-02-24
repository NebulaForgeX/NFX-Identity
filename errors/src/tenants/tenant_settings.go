package tenants

import "nfxid/pkgs/errx"

const (
	CodeTenantSettingNotFound      = "TENANT_SETTING_NOT_FOUND"
	CodeTenantSettingAlreadyExists = "TENANT_SETTING_ALREADY_EXISTS"
)

var (
	ErrTenantSettingNotFound      = errx.NotFound(CodeTenantSettingNotFound, "tenant setting not found")
	ErrTenantSettingAlreadyExists = errx.Conflict(CodeTenantSettingAlreadyExists, "tenant setting already exists")
)
