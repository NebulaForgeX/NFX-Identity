package tenants

import "nfxidentity/pkgs/errx"

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

/*
!DOMAIN_VERIFICATION_NOT_FOUND
*en<domain verification not found>
*zh<域名验证不存在>
*fr<vérification de domaine introuvable>

!DOMAIN_REQUIRED
*en<domain required>
*zh<域名为必填>
*fr<domaine requis>

!DOMAIN_VERIFICATION_ALREADY_EXISTS
*en<domain verification already exists>
*zh<域名验证已存在>
*fr<vérification de domaine existe déjà>

!INVALID_VERIFICATION_METHOD
*en<invalid verification method>
*zh<无效的验证方式>
*fr<méthode de vérification invalide>

!INVALID_VERIFICATION_STATUS
*en<invalid verification status>
*zh<无效的验证状态>
*fr<statut de vérification invalide>

!DOMAIN_VERIFICATION_EXPIRED
*en<domain verification expired>
*zh<域名验证已过期>
*fr<vérification de domaine expirée>

*/
