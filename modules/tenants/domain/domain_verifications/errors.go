package domain_verifications

import "nfxid/pkgs/errx"

var (
	ErrDomainVerificationNotFound   = errx.NotFound("DOMAIN_VERIFICATION_NOT_FOUND", "domain verification not found")
	ErrTenantIDRequired             = errx.InvalidArg("TENANT_ID_REQUIRED", "tenant id is required")
	ErrDomainRequired               = errx.InvalidArg("DOMAIN_REQUIRED", "domain is required")
	ErrDomainVerificationAlreadyExists = errx.Conflict("DOMAIN_VERIFICATION_ALREADY_EXISTS", "domain verification already exists")
	ErrInvalidVerificationMethod    = errx.InvalidArg("INVALID_VERIFICATION_METHOD", "invalid verification method")
	ErrInvalidVerificationStatus    = errx.InvalidArg("INVALID_VERIFICATION_STATUS", "invalid verification status")
	ErrDomainVerificationExpired    = errx.Expired("DOMAIN_VERIFICATION_EXPIRED", "domain verification expired")
)
