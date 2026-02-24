package clients

import "nfxid/pkgs/errx"

const (
	CodeAPIKeyNotFound       = "API_KEY_NOT_FOUND"
	CodeKeyIDRequired        = "KEY_ID_REQUIRED"
	CodeKeyHashRequired      = "KEY_HASH_REQUIRED"
	CodeKeyIDAlreadyExists   = "KEY_ID_ALREADY_EXISTS"
	CodeInvalidAPIKeyStatus  = "INVALID_API_KEY_STATUS"
	CodeAPIKeyAlreadyRevoked = "API_KEY_ALREADY_REVOKED"
	CodeAPIKeyExpired        = "API_KEY_EXPIRED"
)

var (
	ErrAPIKeyNotFound       = errx.NotFound(CodeAPIKeyNotFound, "api key not found")
	ErrKeyIDRequired        = errx.InvalidArg(CodeKeyIDRequired, "key id is required")
	ErrKeyHashRequired      = errx.InvalidArg(CodeKeyHashRequired, "key hash is required")
	ErrKeyIDAlreadyExists   = errx.Conflict(CodeKeyIDAlreadyExists, "key id already exists")
	ErrInvalidAPIKeyStatus  = errx.InvalidArg(CodeInvalidAPIKeyStatus, "invalid api key status")
	ErrAPIKeyAlreadyRevoked = errx.Conflict(CodeAPIKeyAlreadyRevoked, "api key already revoked")
	ErrAPIKeyExpired        = errx.Expired(CodeAPIKeyExpired, "api key expired")
)

/*
!API_KEY_NOT_FOUND
*en<api key not found>
*zh<API 密钥不存在>
*fr<clé API introuvable>

!KEY_ID_REQUIRED
*en<key id required>
*zh<密钥 ID 为必填>
*fr<id de clé requis>

!KEY_HASH_REQUIRED
*en<key hash required>
*zh<密钥哈希为必填>
*fr<hachage de clé requis>

!KEY_ID_ALREADY_EXISTS
*en<key id already exists>
*zh<密钥 ID 已存在>
*fr<id de clé existe déjà>

!INVALID_API_KEY_STATUS
*en<invalid api key status>
*zh<无效的 API 密钥状态>
*fr<statut de clé API invalide>

!API_KEY_ALREADY_REVOKED
*en<api key already revoked>
*zh<API 密钥已撤销>
*fr<clé API déjà révoquée>

!API_KEY_EXPIRED
*en<api key expired>
*zh<API 密钥已过期>
*fr<clé API expirée>

*/
