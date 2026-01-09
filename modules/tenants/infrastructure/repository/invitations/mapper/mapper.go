package mapper

import (
	"encoding/json"
	"strings"
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

	// 序列化 RoleIDs 为 PostgreSQL UUID数组格式 "{uuid1,uuid2,...}"
	var roleIdsStr *string
	if len(i.RoleIDs()) > 0 {
		ids := make([]string, len(i.RoleIDs()))
		for j, id := range i.RoleIDs() {
			ids[j] = id.String()
		}
		idsStr := "{" + strings.Join(ids, ",") + "}"
		roleIdsStr = &idsStr
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
		RoleIds:          roleIdsStr, // Model 使用 RoleIds，Domain 使用 RoleIDs
		Metadata:         metadata,
	}
}

// InvitationModelToDomain 将 Model Invitation 转换为 Domain Invitation
func InvitationModelToDomain(m *models.Invitation) *invitations.Invitation {
	if m == nil {
		return nil
	}

	// 解析 RoleIDs 从 PostgreSQL UUID数组格式 "{uuid1,uuid2,...}"
	var roleIDs []uuid.UUID
	if m.RoleIds != nil && *m.RoleIds != "" {
		// 移除大括号并分割
		idsStr := strings.Trim(*m.RoleIds, "{}")
		if idsStr != "" {
			idStrs := strings.Split(idsStr, ",")
			roleIDs = make([]uuid.UUID, 0, len(idStrs))
			for _, idStr := range idStrs {
				if id, err := uuid.Parse(strings.TrimSpace(idStr)); err == nil {
					roleIDs = append(roleIDs, id)
				}
			}
		}
	}

	var metadata map[string]interface{}
	if m.Metadata != nil {
		json.Unmarshal(*m.Metadata, &metadata)
	}

	state := invitations.InvitationState{
		ID:              m.ID,
		InviteID:        m.InviteID,
		TenantID:        m.TenantID,
		Email:           m.Email,
		TokenHash:       m.TokenHash,
		ExpiresAt:       m.ExpiresAt,
		Status:          invitationStatusEnumToDomain(m.Status),
		InvitedBy:       m.InvitedBy,
		InvitedAt:       m.InvitedAt,
		AcceptedByUserID: m.AcceptedByUserID,
		AcceptedAt:      m.AcceptedAt,
		RevokedBy:       m.RevokedBy,
		RevokedAt:       m.RevokedAt,
		RevokeReason:    m.RevokeReason,
		RoleIDs:         roleIDs, // Model 使用 RoleIds，Domain 使用 RoleIDs
		Metadata:        metadata,
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
		models.InvitationCols.RoleIds:          m.RoleIds,
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
