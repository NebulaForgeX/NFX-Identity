package auth

import "nfxidentity/pkgs/errx"

var (
	ErrPhoneBindingNotFound     = errx.NotFound("PHONE_BINDING_NOT_FOUND", "phone binding not found")
	ErrPhoneAccountIDInvalid    = errx.InvalidArg("PHONE_ACCOUNT_ID_INVALID", "invalid phone account id")
	ErrPhoneNumberRequired      = errx.InvalidArg("PHONE_NUMBER_REQUIRED", "phone number is required")
	ErrPhoneNumberFormatInvalid = errx.InvalidArg("PHONE_NUMBER_FORMAT_INVALID", "invalid phone number format")
)

/*
!PHONE_BINDING_NOT_FOUND
*en<Phone binding not found>
*zh<手机号绑定不存在>
*fr<Liaison téléphone introuvable>

!PHONE_ACCOUNT_ID_INVALID
*en<Invalid phone account id>
*zh<手机号绑定账号 ID 无效>
*fr<Identifiant de compte du téléphone invalide>

!PHONE_NUMBER_REQUIRED
*en<Phone number is required>
*zh<手机号必填>
*fr<Numéro de téléphone requis>

!PHONE_NUMBER_FORMAT_INVALID
*en<Invalid phone number format>
*zh<手机号格式无效>
*fr<Format de numéro de téléphone invalide>
*/
