package system

import "nfxid/pkgs/errx"

const (
	CodeSystemStateNotFound      = "SYSTEM_STATE_NOT_FOUND"
	CodeSystemAlreadyInitialized = "SYSTEM_ALREADY_INITIALIZED"
	CodeSystemNotInitialized     = "SYSTEM_NOT_INITIALIZED"
)

var (
	ErrSystemStateNotFound = errx.NotFound(CodeSystemStateNotFound, "system state not found")
	ErrAlreadyInitialized  = errx.Conflict(CodeSystemAlreadyInitialized, "system already initialized")
	ErrNotInitialized      = errx.FailedPrecond(CodeSystemNotInitialized, "system not initialized")
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

*/
