package auth

import "nfxidentity/pkgs/errx"

var (
	ErrIdentityNotFound                = errx.NotFound("IDENTITY_NOT_FOUND", "identity not found")
	ErrIdentityAccountIDInvalid        = errx.InvalidArg("IDENTITY_ACCOUNT_ID_INVALID", "invalid identity account id")
	ErrIdentityProviderInvalid         = errx.InvalidArg("IDENTITY_PROVIDER_INVALID", "invalid identity provider")
	ErrIdentityProviderSubjectRequired = errx.InvalidArg("IDENTITY_PROVIDER_SUBJECT_REQUIRED", "identity provider subject is required")
	ErrIdentityPasswordHashRequired    = errx.InvalidArg("IDENTITY_PASSWORD_HASH_REQUIRED", "identity password hash is required")
)

/*
!IDENTITY_NOT_FOUND
*en<Identity not found>
*zh<登录身份不存在>
*fr<Identité introuvable>

!IDENTITY_ACCOUNT_ID_INVALID
*en<Invalid identity account id>
*zh<登录身份账号 ID 无效>
*fr<Identifiant de compte d'identité invalide>

!IDENTITY_PROVIDER_INVALID
*en<Invalid identity provider>
*zh<登录方式无效>
*fr<Fournisseur d'identité invalide>

!IDENTITY_PROVIDER_SUBJECT_REQUIRED
*en<Identity provider subject is required>
*zh<登录标识（subject）必填>
*fr<Sujet du fournisseur d'identité requis>

!IDENTITY_PASSWORD_HASH_REQUIRED
*en<Identity password hash is required>
*zh<密码哈希必填>
*fr<Hash de mot de passe requis>
*/
