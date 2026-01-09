package mapper

import (
	"encoding/json"
	"nfxid/enums"
	"nfxid/modules/tenants/domain/domain_verifications"
	"nfxid/modules/tenants/infrastructure/rdb/models"

	"gorm.io/datatypes"
)

// DomainVerificationDomainToModel 将 Domain DomainVerification 转换为 Model DomainVerification
func DomainVerificationDomainToModel(dv *domain_verifications.DomainVerification) *models.DomainVerification {
	if dv == nil {
		return nil
	}

	var metadata *datatypes.JSON
	if dv.Metadata() != nil && len(dv.Metadata()) > 0 {
		metadataBytes, _ := json.Marshal(dv.Metadata())
		jsonData := datatypes.JSON(metadataBytes)
		metadata = &jsonData
	}

	return &models.DomainVerification{
		ID:                 dv.ID(),
		TenantID:           dv.TenantID(),
		Domain:             dv.Domain(),
		VerificationMethod: verificationMethodDomainToEnum(dv.VerificationMethod()),
		VerificationToken:  dv.VerificationToken(),
		Status:             verificationStatusDomainToEnum(dv.Status()),
		VerifiedAt:         dv.VerifiedAt(),
		ExpiresAt:          dv.ExpiresAt(),
		CreatedAt:          dv.CreatedAt(),
		CreatedBy:          dv.CreatedBy(),
		Metadata:           metadata,
	}
}

// DomainVerificationModelToDomain 将 Model DomainVerification 转换为 Domain DomainVerification
func DomainVerificationModelToDomain(m *models.DomainVerification) *domain_verifications.DomainVerification {
	if m == nil {
		return nil
	}

	var metadata map[string]interface{}
	if m.Metadata != nil {
		json.Unmarshal(*m.Metadata, &metadata)
	}

	state := domain_verifications.DomainVerificationState{
		ID:                 m.ID,
		TenantID:           m.TenantID,
		Domain:             m.Domain,
		VerificationMethod: verificationMethodEnumToDomain(m.VerificationMethod),
		VerificationToken:  m.VerificationToken,
		Status:             verificationStatusEnumToDomain(m.Status),
		VerifiedAt:         m.VerifiedAt,
		ExpiresAt:          m.ExpiresAt,
		CreatedAt:          m.CreatedAt,
		CreatedBy:          m.CreatedBy,
		Metadata:           metadata,
	}

	return domain_verifications.NewDomainVerificationFromState(state)
}

// DomainVerificationModelToUpdates 将 Model DomainVerification 转换为更新字段映射
func DomainVerificationModelToUpdates(m *models.DomainVerification) map[string]any {
	var metadata any
	if m.Metadata != nil {
		metadata = m.Metadata
	}

	return map[string]any{
		models.DomainVerificationCols.TenantID:           m.TenantID,
		models.DomainVerificationCols.Domain:             m.Domain,
		models.DomainVerificationCols.VerificationMethod: m.VerificationMethod,
		models.DomainVerificationCols.VerificationToken:  m.VerificationToken,
		models.DomainVerificationCols.Status:             m.Status,
		models.DomainVerificationCols.VerifiedAt:         m.VerifiedAt,
		models.DomainVerificationCols.ExpiresAt:          m.ExpiresAt,
		models.DomainVerificationCols.CreatedBy:          m.CreatedBy,
		models.DomainVerificationCols.Metadata:           metadata,
	}
}

// 枚举转换辅助函数

func verificationMethodDomainToEnum(vm domain_verifications.VerificationMethod) enums.TenantsVerificationMethod {
	switch vm {
	case domain_verifications.VerificationMethodDNS:
		return enums.TenantsVerificationMethodDns
	case domain_verifications.VerificationMethodTXT:
		return enums.TenantsVerificationMethodTxt
	case domain_verifications.VerificationMethodHTML:
		return enums.TenantsVerificationMethodHtml
	case domain_verifications.VerificationMethodFILE:
		return enums.TenantsVerificationMethodFile
	default:
		return enums.TenantsVerificationMethodDns
	}
}

func verificationMethodEnumToDomain(vm enums.TenantsVerificationMethod) domain_verifications.VerificationMethod {
	switch vm {
	case enums.TenantsVerificationMethodDns:
		return domain_verifications.VerificationMethodDNS
	case enums.TenantsVerificationMethodTxt:
		return domain_verifications.VerificationMethodTXT
	case enums.TenantsVerificationMethodHtml:
		return domain_verifications.VerificationMethodHTML
	case enums.TenantsVerificationMethodFile:
		return domain_verifications.VerificationMethodFILE
	default:
		return domain_verifications.VerificationMethodDNS
	}
}

// VerificationStatusDomainToEnum 将 Domain VerificationStatus 转换为 Enum VerificationStatus
func VerificationStatusDomainToEnum(vs domain_verifications.VerificationStatus) enums.TenantsVerificationStatus {
	return verificationStatusDomainToEnum(vs)
}

func verificationStatusDomainToEnum(vs domain_verifications.VerificationStatus) enums.TenantsVerificationStatus {
	switch vs {
	case domain_verifications.VerificationStatusPending:
		return enums.TenantsVerificationStatusPending
	case domain_verifications.VerificationStatusVerified:
		return enums.TenantsVerificationStatusVerified
	case domain_verifications.VerificationStatusFailed:
		return enums.TenantsVerificationStatusFailed
	case domain_verifications.VerificationStatusExpired:
		return enums.TenantsVerificationStatusExpired
	default:
		return enums.TenantsVerificationStatusPending
	}
}

func verificationStatusEnumToDomain(vs enums.TenantsVerificationStatus) domain_verifications.VerificationStatus {
	switch vs {
	case enums.TenantsVerificationStatusPending:
		return domain_verifications.VerificationStatusPending
	case enums.TenantsVerificationStatusVerified:
		return domain_verifications.VerificationStatusVerified
	case enums.TenantsVerificationStatusFailed:
		return domain_verifications.VerificationStatusFailed
	case enums.TenantsVerificationStatusExpired:
		return domain_verifications.VerificationStatusExpired
	default:
		return domain_verifications.VerificationStatusPending
	}
}
