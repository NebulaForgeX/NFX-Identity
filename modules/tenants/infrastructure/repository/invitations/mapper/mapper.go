package mapper

import (
	"encoding/json"
	"nfxid/enums"
	"nfxid/modules/tenants/domain/invitations"
	"nfxid/modules/tenants/infrastructure/rdb/models"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

// InvitationDomainToModel 将 Domain Invitation 转换为 Model Invitation
func InvitationDomainToModel(i *invitations.Invitation) *models.Invitation {
	if i == nil {
		return nil
	}

	// Domain 使用 RoleIDs 切片，Model 使用单个 TenantRoleID，取第一个
	var tenantRoleID uuid.UUID
	if len(i.RoleIDs()) > 0 {
		tenantRoleID = i.RoleIDs()[0]
	}

	var metadata *datatypes.JSON
	if i.Metadata() != nil && len(i.Metadata()) > 0 {
		metadataBytes, _ := json.Marshal(i.Metadata())
		jsonData := datatypes.JSON(metadataBytes)
		metadata = &jsonData
	}

	return &models.Invitation{
		ID:               i.ID(),
		InviteID:         i.InviteID(),
		TenantID:         i.TenantID(),
		Email:            i.Email(),
		TokenHash:        i.TokenHash(),
		ExpiresAt:        i.ExpiresAt(),
		Status:           invitationStatusDomainToEnum(i.Status()),
		InvitedBy:        i.InvitedBy(),
		InvitedAt:        i.InvitedAt(),
		AcceptedByUserID: i.AcceptedByUserID(),
		AcceptedAt:       i.AcceptedAt(),
		RevokedBy:        i.RevokedBy(),
		RevokedAt:        i.RevokedAt(),
		RevokeReason:     i.RevokeReason(),
		TenantRoleID:     tenantRoleID,
		Metadata:         metadata,
	}
}

// InvitationModelToDomain 将 Model Invitation 转换为 Domain Invitation
func InvitationModelToDomain(m *models.Invitation) *invitations.Invitation {
	if m == nil {
		return nil
	}

	// Model 使用单个 TenantRoleID，Domain 使用 RoleIDs 切片
	var roleIDs []uuid.UUID
	if m.TenantRoleID != uuid.Nil {
		roleIDs = []uuid.UUID{m.TenantRoleID}
	}

	var metadata map[string]interface{}
	if m.Metadata != nil {
		json.Unmarshal(*m.Metadata, &metadata)
	}

	state := invitations.InvitationState{
		ID:               m.ID,
		InviteID:         m.InviteID,
		TenantID:         m.TenantID,
		Email:            m.Email,
		TokenHash:        m.TokenHash,
		ExpiresAt:        m.ExpiresAt,
		Status:           invitationStatusEnumToDomain(m.Status),
		InvitedBy:        m.InvitedBy,
		InvitedAt:        m.InvitedAt,
		AcceptedByUserID: m.AcceptedByUserID,
		AcceptedAt:       m.AcceptedAt,
		RevokedBy:        m.RevokedBy,
		RevokedAt:        m.RevokedAt,
		RevokeReason:     m.RevokeReason,
		RoleIDs:          roleIDs,
		Metadata:         metadata,
	}

	return invitations.NewInvitationFromState(state)
}

// InvitationModelToUpdates 将 Model Invitation 转换为更新字段映射
func InvitationModelToUpdates(m *models.Invitation) map[string]any {
	var metadata any
	if m.Metadata != nil {
		metadata = m.Metadata
	}

	return map[string]any{
		models.InvitationCols.InviteID:         m.InviteID,
		models.InvitationCols.TenantID:         m.TenantID,
		models.InvitationCols.Email:            m.Email,
		models.InvitationCols.TokenHash:        m.TokenHash,
		models.InvitationCols.ExpiresAt:        m.ExpiresAt,
		models.InvitationCols.Status:           m.Status,
		models.InvitationCols.InvitedBy:        m.InvitedBy,
		models.InvitationCols.InvitedAt:        m.InvitedAt,
		models.InvitationCols.AcceptedByUserID: m.AcceptedByUserID,
		models.InvitationCols.AcceptedAt:       m.AcceptedAt,
		models.InvitationCols.RevokedBy:        m.RevokedBy,
		models.InvitationCols.RevokedAt:        m.RevokedAt,
		models.InvitationCols.RevokeReason:     m.RevokeReason,
		models.InvitationCols.TenantRoleID:     m.TenantRoleID,
		models.InvitationCols.Metadata:         metadata,
	}
}

// 枚举转换辅助函数

// InvitationStatusDomainToEnum 将 Domain InvitationStatus 转换为 Enum InvitationStatus
func InvitationStatusDomainToEnum(is invitations.InvitationStatus) enums.TenantsInvitationStatus {
	return invitationStatusDomainToEnum(is)
}

func invitationStatusDomainToEnum(is invitations.InvitationStatus) enums.TenantsInvitationStatus {
	switch is {
	case invitations.InvitationStatusPending:
		return enums.TenantsInvitationStatusPending
	case invitations.InvitationStatusAccepted:
		return enums.TenantsInvitationStatusAccepted
	case invitations.InvitationStatusExpired:
		return enums.TenantsInvitationStatusExpired
	case invitations.InvitationStatusRevoked:
		return enums.TenantsInvitationStatusRevoked
	default:
		return enums.TenantsInvitationStatusPending
	}
}

func invitationStatusEnumToDomain(is enums.TenantsInvitationStatus) invitations.InvitationStatus {
	switch is {
	case enums.TenantsInvitationStatusPending:
		return invitations.InvitationStatusPending
	case enums.TenantsInvitationStatusAccepted:
		return invitations.InvitationStatusAccepted
	case enums.TenantsInvitationStatusExpired:
		return invitations.InvitationStatusExpired
	case enums.TenantsInvitationStatusRevoked:
		return invitations.InvitationStatusRevoked
	default:
		return invitations.InvitationStatusPending
	}
}
