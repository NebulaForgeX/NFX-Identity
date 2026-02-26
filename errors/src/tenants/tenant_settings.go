package tenants

import "nfxidentity/pkgs/errx"

const (
	CodeTenantSettingNotFound      = "TENANT_SETTING_NOT_FOUND"
	CodeTenantSettingAlreadyExists = "TENANT_SETTING_ALREADY_EXISTS"
)

var (
	ErrTenantSettingNotFound      = errx.NotFound(CodeTenantSettingNotFound, "tenant setting not found")
	ErrTenantSettingAlreadyExists = errx.Conflict(CodeTenantSettingAlreadyExists, "tenant setting already exists")
)

/*
!TENANT_SETTING_NOT_FOUND
*en<tenant setting not found>
*zh<租户设置不存在>
*fr<paramètre tenant introuvable>

!TENANT_SETTING_ALREADY_EXISTS
*en<tenant setting already exists>
*zh<租户设置已存在>
*fr<paramètre tenant existe déjà>

*/
