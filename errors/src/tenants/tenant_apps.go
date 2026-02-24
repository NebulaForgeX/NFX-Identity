package tenants

import "nfxid/pkgs/errx"

const (
	CodeTenantAppNotFound      = "TENANT_APP_NOT_FOUND"
	CodeTenantAppAlreadyExists = "TENANT_APP_ALREADY_EXISTS"
	CodeInvalidTenantAppStatus = "INVALID_TENANT_APP_STATUS"
)

var (
	ErrTenantAppNotFound      = errx.NotFound(CodeTenantAppNotFound, "tenant app not found")
	ErrTenantAppAlreadyExists = errx.Conflict(CodeTenantAppAlreadyExists, "tenant app already exists")
	ErrInvalidTenantAppStatus = errx.InvalidArg(CodeInvalidTenantAppStatus, "invalid tenant app status")
)

/*
!TENANT_APP_NOT_FOUND
*en<tenant app not found>
*zh<租户应用不存在>
*fr<application tenant introuvable>

!TENANT_APP_ALREADY_EXISTS
*en<tenant app already exists>
*zh<租户应用已存在>
*fr<application tenant existe déjà>

!INVALID_TENANT_APP_STATUS
*en<invalid tenant app status>
*zh<无效的租户应用状态>
*fr<statut d'application tenant invalide>

*/
