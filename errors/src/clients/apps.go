package clients

import "nfxidentity/pkgs/errx"

const (
	CodeAppNotFound        = "APP_NOT_FOUND"
	CodeTenantIDRequired   = "TENANT_ID_REQUIRED"
	CodeAppIDAlreadyExists = "APP_ID_ALREADY_EXISTS"
	CodeInvalidAppType     = "INVALID_APP_TYPE"
	CodeInvalidAppStatus   = "INVALID_APP_STATUS"
	CodeInvalidEnvironment = "INVALID_ENVIRONMENT"
)

var (
	ErrAppNotFound        = errx.NotFound(CodeAppNotFound, "app not found")
	ErrTenantIDRequired   = errx.InvalidArg(CodeTenantIDRequired, "tenant id is required")
	ErrAppIDAlreadyExists = errx.Conflict(CodeAppIDAlreadyExists, "app id already exists")
	ErrInvalidAppType     = errx.InvalidArg(CodeInvalidAppType, "invalid app type")
	ErrInvalidAppStatus   = errx.InvalidArg(CodeInvalidAppStatus, "invalid app status")
	ErrInvalidEnvironment = errx.InvalidArg(CodeInvalidEnvironment, "invalid environment")
)

/*
!APP_NOT_FOUND
*en<app not found>
*zh<应用不存在>
*fr<application introuvable>

!TENANT_ID_REQUIRED
*en<tenant id required>
*zh<租户 ID 为必填>
*fr<id tenant requis>

!APP_ID_ALREADY_EXISTS
*en<app id already exists>
*zh<应用 ID 已存在>
*fr<id d'application existe déjà>

!INVALID_APP_TYPE
*en<invalid app type>
*zh<无效的应用类型>
*fr<type d'application invalide>

!INVALID_APP_STATUS
*en<invalid app status>
*zh<无效的应用状态>
*fr<statut d'application invalide>

!INVALID_ENVIRONMENT
*en<invalid environment>
*zh<无效的环境>
*fr<environnement invalide>

*/
