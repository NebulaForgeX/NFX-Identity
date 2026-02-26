package auth

import "nfxidentity/pkgs/errx"

const (
	CodeUserCredentialNotFound      = "USER_CREDENTIAL_NOT_FOUND"
	CodeCredentialTypeRequired      = "CREDENTIAL_TYPE_REQUIRED"
	CodeInvalidCredentialType       = "INVALID_CREDENTIAL_TYPE"
	CodeInvalidCredentialStatus     = "INVALID_CREDENTIAL_STATUS"
	CodePasswordHashRequired        = "PASSWORD_HASH_REQUIRED"
	CodeUserCredentialAlreadyExists = "USER_CREDENTIAL_ALREADY_EXISTS"
)

var (
	ErrUserCredentialNotFound      = errx.NotFound(CodeUserCredentialNotFound, "user credential not found")
	ErrCredentialTypeRequired      = errx.InvalidArg(CodeCredentialTypeRequired, "credential type is required")
	ErrInvalidCredentialType       = errx.InvalidArg(CodeInvalidCredentialType, "invalid credential type")
	ErrInvalidCredentialStatus     = errx.InvalidArg(CodeInvalidCredentialStatus, "invalid credential status")
	ErrPasswordHashRequired        = errx.InvalidArg(CodePasswordHashRequired, "password hash is required")
	ErrUserCredentialAlreadyExists = errx.Conflict(CodeUserCredentialAlreadyExists, "user credential already exists")
)

/*
!USER_CREDENTIAL_NOT_FOUND
*en<user credential not found>
*zh<用户凭据不存在>
*fr<identifiant utilisateur introuvable>

!CREDENTIAL_TYPE_REQUIRED
*en<credential type required>
*zh<凭据类型为必填>
*fr<type d'identifiant requis>

!INVALID_CREDENTIAL_TYPE
*en<invalid credential type>
*zh<无效的凭据类型>
*fr<type d'identifiant invalide>

!INVALID_CREDENTIAL_STATUS
*en<invalid credential status>
*zh<无效的凭据状态>
*fr<statut d'identifiant invalide>

!PASSWORD_HASH_REQUIRED
*en<password hash required>
*zh<密码哈希为必填>
*fr<hachage du mot de passe requis>

!USER_CREDENTIAL_ALREADY_EXISTS
*en<user credential already exists>
*zh<用户凭据已存在>
*fr<identifiant utilisateur existe déjà>

*/
