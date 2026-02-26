package directory

import "nfxidentity/pkgs/errx"

const (
	CodeUserPhoneNotFound       = "USER_PHONE_NOT_FOUND"
	CodePhoneRequired           = "PHONE_REQUIRED"
	CodePhoneAlreadyExists      = "PHONE_ALREADY_EXISTS"
	CodeInvalidPhone            = "INVALID_PHONE"
	CodeVerificationCodeExpired = "VERIFICATION_CODE_EXPIRED"
)

var (
	ErrUserPhoneNotFound       = errx.NotFound(CodeUserPhoneNotFound, "user phone not found")
	ErrPhoneRequired           = errx.InvalidArg(CodePhoneRequired, "phone is required")
	ErrPhoneAlreadyExists      = errx.Conflict(CodePhoneAlreadyExists, "phone already exists")
	ErrInvalidPhone            = errx.InvalidArg(CodeInvalidPhone, "invalid phone format")
	ErrVerificationCodeExpired = errx.Expired(CodeVerificationCodeExpired, "verification code expired")
)

/*
!USER_PHONE_NOT_FOUND
*en<user phone not found>
*zh<用户手机号不存在>
*fr<téléphone utilisateur introuvable>

!PHONE_REQUIRED
*en<phone required>
*zh<手机号为必填>
*fr<téléphone requis>

!PHONE_ALREADY_EXISTS
*en<phone already exists>
*zh<手机号已存在>
*fr<téléphone existe déjà>

!INVALID_PHONE
*en<invalid phone>
*zh<无效的手机号格式>
*fr<téléphone invalide>

!VERIFICATION_CODE_EXPIRED
*en<verification code expired>
*zh<验证码已过期>
*fr<code de vérification expiré>

*/
