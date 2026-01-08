package domain_verifications

import "errors"

var (
	ErrDomainVerificationNotFound   = errors.New("domain verification not found")
	ErrTenantIDRequired             = errors.New("tenant id is required")
	ErrDomainRequired               = errors.New("domain is required")
	ErrDomainVerificationAlreadyExists = errors.New("domain verification already exists")
	ErrInvalidVerificationMethod    = errors.New("invalid verification method")
	ErrInvalidVerificationStatus    = errors.New("invalid verification status")
	ErrDomainVerificationExpired    = errors.New("domain verification expired")
)
