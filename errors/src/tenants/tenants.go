package tenants

import "nfxidentity/pkgs/errx"

const (
	CodeTenantNotFound        = "TENANT_NOT_FOUND"
	CodeTenantIDAlreadyExists = "TENANT_ID_ALREADY_EXISTS"
	CodeInvalidTenantStatus   = "INVALID_TENANT_STATUS"
)

var (
	ErrTenantNotFound        = errx.NotFound(CodeTenantNotFound, "tenant not found")
	ErrTenantIDAlreadyExists = errx.Conflict(CodeTenantIDAlreadyExists, "tenant id already exists")
	ErrInvalidTenantStatus   = errx.InvalidArg(CodeInvalidTenantStatus, "invalid tenant status")
)

/*
!TENANT_NOT_FOUND
*en<tenant not found>
*zh<租户不存在>
*fr<tenant introuvable>

!TENANT_ID_ALREADY_EXISTS
*en<tenant id already exists>
*zh<租户 ID 已存在>
*fr<id tenant existe déjà>

!INVALID_TENANT_STATUS
*en<invalid tenant status>
*zh<无效的租户状态>
*fr<statut tenant invalide>

*/
