package clients

import "nfxid/pkgs/errx"

// Shared codes and errors used by multiple client domains.
const (
	CodeAppIDRequired   = "APP_ID_REQUIRED"
	CodeNameRequired    = "NAME_REQUIRED"
	CodeHashAlgRequired = "HASH_ALG_REQUIRED"
)

var (
	ErrAppIDRequired   = errx.InvalidArg(CodeAppIDRequired, "app id is required")
	ErrNameRequired    = errx.InvalidArg(CodeNameRequired, "name is required")
	ErrHashAlgRequired = errx.InvalidArg(CodeHashAlgRequired, "hash alg is required")
)

/*
!APP_ID_REQUIRED
*en<app id required>
*zh<应用 ID 为必填>
*fr<id d'application requis>

!NAME_REQUIRED
*en<name required>
*zh<名称为必填>
*fr<nom requis>

!HASH_ALG_REQUIRED
*en<hash alg required>
*zh<哈希算法为必填>
*fr<algorithme de hachage requis>

*/
