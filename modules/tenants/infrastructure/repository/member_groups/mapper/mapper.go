package mapper

import (
	"nfxid/modules/tenants/domain/member_groups"
	"nfxid/modules/tenants/infrastructure/rdb/models"
)

// MemberGroupDomainToModel 将 Domain MemberGroup 转换为 Model MemberGroup
func MemberGroupDomainToModel(mg *member_groups.MemberGroup) *models.MemberGroup {
	if mg == nil {
		return nil
	}

	return &models.MemberGroup{
		ID:         mg.ID(),
		MemberID:   mg.MemberID(),
		GroupID:    mg.GroupID(),
		AssignedAt: mg.AssignedAt(),
		AssignedBy: mg.AssignedBy(),
		RevokedAt:  mg.RevokedAt(),
		RevokedBy:  mg.RevokedBy(),
	}
}

// MemberGroupModelToDomain 将 Model MemberGroup 转换为 Domain MemberGroup
func MemberGroupModelToDomain(m *models.MemberGroup) *member_groups.MemberGroup {
	if m == nil {
		return nil
	}

	state := member_groups.MemberGroupState{
		ID:         m.ID,
		MemberID:   m.MemberID,
		GroupID:    m.GroupID,
		AssignedAt: m.AssignedAt,
		AssignedBy: m.AssignedBy,
		RevokedAt:  m.RevokedAt,
		RevokedBy:  m.RevokedBy,
	}

	return member_groups.NewMemberGroupFromState(state)
}

// MemberGroupModelToUpdates 将 Model MemberGroup 转换为更新字段映射
func MemberGroupModelToUpdates(m *models.MemberGroup) map[string]any {
	return map[string]any{
		models.MemberGroupCols.MemberID:   m.MemberID,
		models.MemberGroupCols.GroupID:    m.GroupID,
		models.MemberGroupCols.AssignedAt: m.AssignedAt,
		models.MemberGroupCols.AssignedBy: m.AssignedBy,
		models.MemberGroupCols.RevokedAt:  m.RevokedAt,
		models.MemberGroupCols.RevokedBy:  m.RevokedBy,
	}
}
