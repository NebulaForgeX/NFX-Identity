package auth

import "nfxid/pkgs/errx"

const (
	CodePasswordResetNotFound = "PASSWORD_RESET_NOT_FOUND"
	CodeResetIDRequired       = "RESET_ID_REQUIRED"
	CodeDeliveryRequired      = "DELIVERY_REQUIRED"
	CodeCodeHashRequired      = "CODE_HASH_REQUIRED"
	CodeResetIDAlreadyExists  = "RESET_ID_ALREADY_EXISTS"
	CodeInvalidResetDelivery  = "INVALID_RESET_DELIVERY"
	CodeInvalidResetStatus    = "INVALID_RESET_STATUS"
	CodeResetAlreadyUsed      = "RESET_ALREADY_USED"
	CodeResetExpired          = "RESET_EXPIRED"
)

var (
	ErrPasswordResetNotFound = errx.NotFound(CodePasswordResetNotFound, "password reset not found")
	ErrResetIDRequired       = errx.InvalidArg(CodeResetIDRequired, "reset id is required")
	ErrDeliveryRequired      = errx.InvalidArg(CodeDeliveryRequired, "delivery is required")
	ErrCodeHashRequired      = errx.InvalidArg(CodeCodeHashRequired, "code hash is required")
	ErrResetIDAlreadyExists  = errx.Conflict(CodeResetIDAlreadyExists, "reset id already exists")
	ErrInvalidResetDelivery  = errx.InvalidArg(CodeInvalidResetDelivery, "invalid reset delivery")
	ErrInvalidResetStatus    = errx.InvalidArg(CodeInvalidResetStatus, "invalid reset status")
	ErrResetAlreadyUsed      = errx.Conflict(CodeResetAlreadyUsed, "reset already used")
	ErrResetExpired          = errx.Expired(CodeResetExpired, "reset expired")
)

/*
!PASSWORD_RESET_NOT_FOUND
*en<password reset not found>
*zh<密码重置记录不存在>
*fr<réinitialisation du mot de passe introuvable>

!RESET_ID_REQUIRED
*en<reset id required>
*zh<重置 ID 为必填>
*fr<id de réinitialisation requis>

!DELIVERY_REQUIRED
*en<delivery required>
*zh<交付方式为必填>
*fr<livraison requise>

!CODE_HASH_REQUIRED
*en<code hash required>
*zh<验证码哈希为必填>
*fr<hachage du code requis>

!RESET_ID_ALREADY_EXISTS
*en<reset id already exists>
*zh<重置 ID 已存在>
*fr<id de réinitialisation existe déjà>

!INVALID_RESET_DELIVERY
*en<invalid reset delivery>
*zh<无效的重置交付方式>
*fr<livraison de réinitialisation invalide>

!INVALID_RESET_STATUS
*en<invalid reset status>
*zh<无效的重置状态>
*fr<statut de réinitialisation invalide>

!RESET_ALREADY_USED
*en<reset already used>
*zh<重置已使用>
*fr<réinitialisation déjà utilisée>

!RESET_EXPIRED
*en<reset expired>
*zh<重置已过期>
*fr<réinitialisation expirée>

*/
