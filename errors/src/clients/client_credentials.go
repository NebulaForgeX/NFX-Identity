package clients

import "nfxid/pkgs/errx"

const (
	CodeClientCredentialNotFound = "CLIENT_CREDENTIAL_NOT_FOUND"
	CodeClientIDRequired         = "CLIENT_ID_REQUIRED"
	CodeSecretHashRequired       = "SECRET_HASH_REQUIRED"
	CodeClientIDAlreadyExists    = "CLIENT_ID_ALREADY_EXISTS"
	CodeInvalidCredentialStatus  = "INVALID_CREDENTIAL_STATUS"
	CodeCredentialAlreadyRevoked = "CREDENTIAL_ALREADY_REVOKED"
	CodeCredentialExpired        = "CREDENTIAL_EXPIRED"
)

var (
	ErrClientCredentialNotFound = errx.NotFound(CodeClientCredentialNotFound, "client credential not found")
	ErrClientIDRequired         = errx.InvalidArg(CodeClientIDRequired, "client id is required")
	ErrSecretHashRequired       = errx.InvalidArg(CodeSecretHashRequired, "secret hash is required")
	ErrClientIDAlreadyExists    = errx.Conflict(CodeClientIDAlreadyExists, "client id already exists")
	ErrInvalidCredentialStatus  = errx.InvalidArg(CodeInvalidCredentialStatus, "invalid credential status")
	ErrCredentialAlreadyRevoked = errx.Conflict(CodeCredentialAlreadyRevoked, "credential already revoked")
	ErrCredentialExpired        = errx.Expired(CodeCredentialExpired, "credential expired")
)

/*
!CLIENT_CREDENTIAL_NOT_FOUND
*en<client credential not found>
*zh<客户端凭据不存在>
*fr<identifiant client introuvable>

!CLIENT_ID_REQUIRED
*en<client id required>
*zh<客户端 ID 为必填>
*fr<id client requis>

!SECRET_HASH_REQUIRED
*en<secret hash required>
*zh<密钥哈希为必填>
*fr<hachage du secret requis>

!CLIENT_ID_ALREADY_EXISTS
*en<client id already exists>
*zh<客户端 ID 已存在>
*fr<id client existe déjà>

!INVALID_CREDENTIAL_STATUS
*en<invalid credential status>
*zh<无效的凭据状态>
*fr<statut d'identifiant invalide>

!CREDENTIAL_ALREADY_REVOKED
*en<credential already revoked>
*zh<凭据已撤销>
*fr<identifiant déjà révoqué>

!CREDENTIAL_EXPIRED
*en<credential expired>
*zh<凭据已过期>
*fr<identifiant expiré>

*/
