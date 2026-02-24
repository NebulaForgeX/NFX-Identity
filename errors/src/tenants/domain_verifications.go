package tenants

import "nfxid/pkgs/errx"

const (
	CodeDomainVerificationNotFound      = "DOMAIN_VERIFICATION_NOT_FOUND"
	CodeDomainRequired                  = "DOMAIN_REQUIRED"
	CodeDomainVerificationAlreadyExists = "DOMAIN_VERIFICATION_ALREADY_EXISTS"
	CodeInvalidVerificationMethod       = "INVALID_VERIFICATION_METHOD"
	CodeInvalidVerificationStatus       = "INVALID_VERIFICATION_STATUS"
	CodeDomainVerificationExpired       = "DOMAIN_VERIFICATION_EXPIRED"
)

var (
	ErrDomainVerificationNotFound      = errx.NotFound(CodeDomainVerificationNotFound, "domain verification not found")
	ErrDomainRequired                  = errx.InvalidArg(CodeDomainRequired, "domain is required")
	ErrDomainVerificationAlreadyExists = errx.Conflict(CodeDomainVerificationAlreadyExists, "domain verification already exists")
	ErrInvalidVerificationMethod       = errx.InvalidArg(CodeInvalidVerificationMethod, "invalid verification method")
	ErrInvalidVerificationStatus       = errx.InvalidArg(CodeInvalidVerificationStatus, "invalid verification status")
	ErrDomainVerificationExpired       = errx.Expired(CodeDomainVerificationExpired, "domain verification expired")
)
