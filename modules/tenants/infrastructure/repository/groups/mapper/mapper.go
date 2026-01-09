package mapper

import (
	"encoding/json"
	"nfxid/enums"
	"nfxid/modules/tenants/domain/groups"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/pkgs/utils/timex"

	"gorm.io/datatypes"
)

// GroupDomainToModel 将 Domain Group 转换为 Model Group
func GroupDomainToModel(g *groups.Group) *models.Group {
	if g == nil {
		return nil
	}

	var metadata *datatypes.JSON
	if g.Metadata() != nil && len(g.Metadata()) > 0 {
		metadataBytes, _ := json.Marshal(g.Metadata())
		jsonData := datatypes.JSON(metadataBytes)
		metadata = &jsonData
	}

	return &models.Group{
		ID:            g.ID(),
		GroupID:       g.GroupID(),
		TenantID:      g.TenantID(),
		Name:          g.Name(),
		Type:          groupTypeDomainToEnum(g.Type()),
		ParentGroupID: g.ParentGroupID(),
		Description:   g.Description(),
		CreatedAt:     g.CreatedAt(),
		UpdatedAt:     g.UpdatedAt(),
		CreatedBy:     g.CreatedBy(),
		DeletedAt:     timex.TimeToGormDeletedAt(g.DeletedAt()),
		Metadata:      metadata,
	}
}

// GroupModelToDomain 将 Model Group 转换为 Domain Group
func GroupModelToDomain(m *models.Group) *groups.Group {
	if m == nil {
		return nil
	}

	var metadata map[string]interface{}
	if m.Metadata != nil {
		json.Unmarshal(*m.Metadata, &metadata)
	}

	state := groups.GroupState{
		ID:            m.ID,
		GroupID:       m.GroupID,
		TenantID:      m.TenantID,
		Name:          m.Name,
		Type:          groupTypeEnumToDomain(m.Type),
		ParentGroupID: m.ParentGroupID,
		Description:   m.Description,
		CreatedAt:     m.CreatedAt,
		UpdatedAt:     m.UpdatedAt,
		CreatedBy:     m.CreatedBy,
		DeletedAt:     timex.GormDeletedAtToTime(m.DeletedAt),
		Metadata:      metadata,
	}

	return groups.NewGroupFromState(state)
}

// GroupModelToUpdates 将 Model Group 转换为更新字段映射
func GroupModelToUpdates(m *models.Group) map[string]any {
	var metadata any
	if m.Metadata != nil {
		metadata = m.Metadata
	}

	return map[string]any{
		models.GroupCols.GroupID:       m.GroupID,
		models.GroupCols.TenantID:      m.TenantID,
		models.GroupCols.Name:          m.Name,
		models.GroupCols.Type:          m.Type,
		models.GroupCols.ParentGroupID: m.ParentGroupID,
		models.GroupCols.Description:   m.Description,
		models.GroupCols.UpdatedAt:     m.UpdatedAt,
		models.GroupCols.CreatedBy:     m.CreatedBy,
		models.GroupCols.DeletedAt:     m.DeletedAt,
		models.GroupCols.Metadata:      metadata,
	}
}

// 枚举转换辅助函数

// GroupTypeDomainToEnum 将 Domain GroupType 转换为 Enum GroupType
func GroupTypeDomainToEnum(gt groups.GroupType) enums.TenantsGroupType {
	return groupTypeDomainToEnum(gt)
}

func groupTypeDomainToEnum(gt groups.GroupType) enums.TenantsGroupType {
	switch gt {
	case groups.GroupTypeDepartment:
		return enums.TenantsGroupTypeDepartment
	case groups.GroupTypeTeam:
		return enums.TenantsGroupTypeTeam
	case groups.GroupTypeGroup:
		return enums.TenantsGroupTypeGroup
	case groups.GroupTypeOther:
		return enums.TenantsGroupTypeOther
	default:
		return enums.TenantsGroupTypeGroup
	}
}

func groupTypeEnumToDomain(gt enums.TenantsGroupType) groups.GroupType {
	switch gt {
	case enums.TenantsGroupTypeDepartment:
		return groups.GroupTypeDepartment
	case enums.TenantsGroupTypeTeam:
		return groups.GroupTypeTeam
	case enums.TenantsGroupTypeGroup:
		return groups.GroupTypeGroup
	case enums.TenantsGroupTypeOther:
		return groups.GroupTypeOther
	default:
		return groups.GroupTypeGroup
	}
}
