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
