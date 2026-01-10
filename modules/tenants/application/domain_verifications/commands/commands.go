package commands

import (
	"nfxid/modules/tenants/domain/domain_verifications"

	"github.com/google/uuid"
)

// CreateDomainVerificationCmd 创建域名验证命令
type CreateDomainVerificationCmd struct {
	TenantID           uuid.UUID
	Domain             string
	VerificationMethod domain_verifications.VerificationMethod
	VerificationToken  *string
	Status             domain_verifications.VerificationStatus
	ExpiresAt          *string
	CreatedBy          *uuid.UUID
	Metadata           map[string]interface{}
}

// VerifyDomainCmd 验证域名命令
type VerifyDomainCmd struct {
	DomainVerificationID uuid.UUID
}

// FailDomainVerificationCmd 标记域名验证失败命令
type FailDomainVerificationCmd struct {
	DomainVerificationID uuid.UUID
}

// DeleteDomainVerificationCmd 删除域名验证命令
type DeleteDomainVerificationCmd struct {
	DomainVerificationID uuid.UUID
}
