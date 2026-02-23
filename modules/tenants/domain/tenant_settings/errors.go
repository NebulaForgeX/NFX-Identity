package tenant_settings

import "nfxid/pkgs/errx"

var (
	ErrTenantSettingNotFound     = errx.NotFound("TENANT_SETTING_NOT_FOUND", "tenant setting not found")
	ErrTenantIDRequired          = errx.InvalidArg("TENANT_ID_REQUIRED", "tenant id is required")
	ErrTenantSettingAlreadyExists = errx.Conflict("TENANT_SETTING_ALREADY_EXISTS", "tenant setting already exists")
)
