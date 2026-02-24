package auth

import "nfxid/pkgs/errx"

// Shared codes and errors used by multiple auth domains.
const (
	CodeUserIDRequired      = "USER_ID_REQUIRED"
	CodeTenantIDRequired    = "TENANT_ID_REQUIRED"
	CodeExpiresAtRequired   = "EXPIRES_AT_REQUIRED"
	CodeInvalidRevokeReason = "INVALID_REVOKE_REASON"
)

var (
	ErrUserIDRequired      = errx.InvalidArg(CodeUserIDRequired, "user id is required")
	ErrTenantIDRequired    = errx.InvalidArg(CodeTenantIDRequired, "tenant id is required")
	ErrExpiresAtRequired   = errx.InvalidArg(CodeExpiresAtRequired, "expires at is required")
	ErrInvalidRevokeReason = errx.InvalidArg(CodeInvalidRevokeReason, "invalid revoke reason")
)

/*
!USER_ID_REQUIRED
*en<user id required>
*zh<用户 ID 为必填>
*fr<identifiant utilisateur requis>

!TENANT_ID_REQUIRED
*en<tenant id required>
*zh<租户 ID 为必填>
*fr<identifiant tenant requis>

!EXPIRES_AT_REQUIRED
*en<expires at required>
*zh<过期时间为必填>
*fr<date d'expiration requise>

!INVALID_REVOKE_REASON
*en<invalid revoke reason>
*zh<无效的撤销原因>
*fr<raison de révocation invalide>

*/
