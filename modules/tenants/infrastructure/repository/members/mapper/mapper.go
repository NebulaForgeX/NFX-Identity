package mapper

import (
	"encoding/json"
	"nfxid/enums"
	"nfxid/modules/tenants/domain/members"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/pkgs/utils/timex"

	"gorm.io/datatypes"
)

// MemberDomainToModel 将 Domain Member 转换为 Model Member
func MemberDomainToModel(m *members.Member) *models.Member {
	if m == nil {
		return nil
	}

	var metadata *datatypes.JSON
	if m.Metadata() != nil && len(m.Metadata()) > 0 {
		metadataBytes, _ := json.Marshal(m.Metadata())
		jsonData := datatypes.JSON(metadataBytes)
		metadata = &jsonData
	}

	return &models.Member{
		ID:          m.ID(),
		MemberID:    m.MemberID(),
		TenantID:    m.TenantID(),
		UserID:      m.UserID(),
		Status:      memberStatusDomainToEnum(m.Status()),
		Source:      memberSourceDomainToEnum(m.Source()),
		JoinedAt:    m.JoinedAt(),
		LeftAt:      m.LeftAt(),
		CreatedAt:   m.CreatedAt(),
		CreatedBy:   m.CreatedBy(),
		UpdatedAt:   m.UpdatedAt(),
		ExternalRef: m.ExternalRef(),
		Metadata:    metadata,
	}
}

// MemberModelToDomain 将 Model Member 转换为 Domain Member
func MemberModelToDomain(m *models.Member) *members.Member {
	if m == nil {
		return nil
	}

	var metadata map[string]interface{}
	if m.Metadata != nil {
		json.Unmarshal(*m.Metadata, &metadata)
	}

	state := members.MemberState{
		ID:          m.ID,
		MemberID:    m.MemberID,
		TenantID:    m.TenantID,
		UserID:      m.UserID,
		Status:      memberStatusEnumToDomain(m.Status),
		Source:      memberSourceEnumToDomain(m.Source),
		JoinedAt:    m.JoinedAt,
		LeftAt:      m.LeftAt,
		CreatedAt:   m.CreatedAt,
		CreatedBy:   m.CreatedBy,
		UpdatedAt:   m.UpdatedAt,
		ExternalRef: m.ExternalRef,
		Metadata:    metadata,
	}

	return members.NewMemberFromState(state)
}

// MemberModelToUpdates 将 Model Member 转换为更新字段映射
func MemberModelToUpdates(m *models.Member) map[string]any {
	var metadata any
	if m.Metadata != nil {
		metadata = m.Metadata
	}

	return map[string]any{
		models.MemberCols.TenantID:    m.TenantID,
		models.MemberCols.UserID:      m.UserID,
		models.MemberCols.Status:      m.Status,
		models.MemberCols.Source:      m.Source,
		models.MemberCols.JoinedAt:    m.JoinedAt,
		models.MemberCols.LeftAt:      m.LeftAt,
		models.MemberCols.CreatedBy:   m.CreatedBy,
		models.MemberCols.UpdatedAt:   m.UpdatedAt,
		models.MemberCols.ExternalRef: m.ExternalRef,
		models.MemberCols.Metadata:    metadata,
	}
}

// 枚举转换辅助函数

// MemberStatusDomainToEnum 将 Domain MemberStatus 转换为 Enum MemberStatus
func MemberStatusDomainToEnum(ms members.MemberStatus) enums.TenantsMemberStatus {
	return memberStatusDomainToEnum(ms)
}

func memberStatusDomainToEnum(ms members.MemberStatus) enums.TenantsMemberStatus {
	switch ms {
	case members.MemberStatusInvited:
		return enums.TenantsMemberStatusInvited
	case members.MemberStatusActive:
		return enums.TenantsMemberStatusActive
	case members.MemberStatusSuspended:
		return enums.TenantsMemberStatusSuspended
	case members.MemberStatusRemoved:
		return enums.TenantsMemberStatusRemoved
	default:
		return enums.TenantsMemberStatusInvited
	}
}

func memberStatusEnumToDomain(ms enums.TenantsMemberStatus) members.MemberStatus {
	switch ms {
	case enums.TenantsMemberStatusInvited:
		return members.MemberStatusInvited
	case enums.TenantsMemberStatusActive:
		return members.MemberStatusActive
	case enums.TenantsMemberStatusSuspended:
		return members.MemberStatusSuspended
	case enums.TenantsMemberStatusRemoved:
		return members.MemberStatusRemoved
	default:
		return members.MemberStatusInvited
	}
}

func memberSourceDomainToEnum(ms members.MemberSource) enums.TenantsMemberSource {
	switch ms {
	case members.MemberSourceManual:
		return enums.TenantsMemberSourceManual
	case members.MemberSourceInvite:
		return enums.TenantsMemberSourceInvite
	case members.MemberSourceSCIM:
		return enums.TenantsMemberSourceScim
	case members.MemberSourceSSO:
		return enums.TenantsMemberSourceSso
	case members.MemberSourceHRSync:
		return enums.TenantsMemberSourceHrSync
	case members.MemberSourceImport:
		return enums.TenantsMemberSourceImport
	default:
		return enums.TenantsMemberSourceManual
	}
}

func memberSourceEnumToDomain(ms enums.TenantsMemberSource) members.MemberSource {
	switch ms {
	case enums.TenantsMemberSourceManual:
		return members.MemberSourceManual
	case enums.TenantsMemberSourceInvite:
		return members.MemberSourceInvite
	case enums.TenantsMemberSourceScim:
		return members.MemberSourceSCIM
	case enums.TenantsMemberSourceSso:
		return members.MemberSourceSSO
	case enums.TenantsMemberSourceHrSync:
		return members.MemberSourceHRSync
	case enums.TenantsMemberSourceImport:
		return members.MemberSourceImport
	default:
		return members.MemberSourceManual
	}
}
