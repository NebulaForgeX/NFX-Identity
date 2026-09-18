package system

import "nfxidentity/pkgs/errx"

const (
	CodeSystemStateNotFound      = "SYSTEM_STATE_NOT_FOUND"
	CodeSystemAlreadyInitialized = "SYSTEM_ALREADY_INITIALIZED"
	CodeSystemNotInitialized     = "SYSTEM_NOT_INITIALIZED"
)

var (
	ErrSystemStateNotFound  = errx.NotFound(CodeSystemStateNotFound, "system state not found")
	ErrAlreadyInitialized   = errx.Conflict(CodeSystemAlreadyInitialized, "system already initialized")
	ErrNotInitialized       = errx.FailedPrecond(CodeSystemNotInitialized, "system not initialized")
	ErrAdminEmailRequired   = errx.InvalidArg("ADMIN_EMAIL_REQUIRED", "administrator email is required")
	ErrBootstrapOwnerFailed = errx.Internal("BOOTSTRAP_OWNER_FAILED", "failed to create bootstrap owner")
)

/*
!SYSTEM_STATE_NOT_FOUND
*en<system state not found>
*zh<系统状态不存在>
*fr<état du système introuvable>

!SYSTEM_ALREADY_INITIALIZED
*en<system already initialized>
*zh<系统已初始化>
*fr<système déjà initialisé>

!SYSTEM_NOT_INITIALIZED
*en<system not initialized>
*zh<系统未初始化>
*fr<système non initialisé>

!ADMIN_EMAIL_REQUIRED
*en<administrator email is required>
*zh<必须填写管理员邮箱>
*fr<l'e-mail administrateur est requis>

!BOOTSTRAP_OWNER_FAILED
*en<failed to create bootstrap owner>
*zh<创建初始 owner 失败>
*fr<échec de la création du propriétaire initial>

*/
