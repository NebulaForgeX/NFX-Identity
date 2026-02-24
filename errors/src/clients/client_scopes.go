package clients

import "nfxid/pkgs/errx"

const (
	CodeClientScopeNotFound       = "CLIENT_SCOPE_NOT_FOUND"
	CodeScopeRequired             = "SCOPE_REQUIRED"
	CodeClientScopeAlreadyExists  = "CLIENT_SCOPE_ALREADY_EXISTS"
	CodeClientScopeAlreadyRevoked = "CLIENT_SCOPE_ALREADY_REVOKED"
	CodeClientScopeExpired        = "CLIENT_SCOPE_EXPIRED"
)

var (
	ErrClientScopeNotFound       = errx.NotFound(CodeClientScopeNotFound, "client scope not found")
	ErrScopeRequired             = errx.InvalidArg(CodeScopeRequired, "scope is required")
	ErrClientScopeAlreadyExists  = errx.Conflict(CodeClientScopeAlreadyExists, "client scope already exists")
	ErrClientScopeAlreadyRevoked = errx.Conflict(CodeClientScopeAlreadyRevoked, "client scope already revoked")
	ErrClientScopeExpired        = errx.Expired(CodeClientScopeExpired, "client scope expired")
)

/*
!CLIENT_SCOPE_NOT_FOUND
*en<client scope not found>
*zh<客户端范围不存在>
*fr<portée client introuvable>

!SCOPE_REQUIRED
*en<scope required>
*zh<范围为必填>
*fr<portée requise>

!CLIENT_SCOPE_ALREADY_EXISTS
*en<client scope already exists>
*zh<客户端范围已存在>
*fr<portée client existe déjà>

!CLIENT_SCOPE_ALREADY_REVOKED
*en<client scope already revoked>
*zh<客户端范围已撤销>
*fr<portée client déjà révoquée>

!CLIENT_SCOPE_EXPIRED
*en<client scope expired>
*zh<客户端范围已过期>
*fr<portée client expirée>

*/
