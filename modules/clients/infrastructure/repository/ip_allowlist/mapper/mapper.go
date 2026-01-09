package mapper

import (
	"nfxid/enums"
	"nfxid/modules/clients/domain/ip_allowlist"
	"nfxid/modules/clients/infrastructure/rdb/models"
)

// IPAllowlistDomainToModel 将 Domain IPAllowlist 转换为 Model IpAllowlist
func IPAllowlistDomainToModel(ip *ip_allowlist.IPAllowlist) *models.IpAllowlist {
	if ip == nil {
		return nil
	}

	return &models.IpAllowlist{
		ID:           ip.ID(),
		RuleID:       ip.RuleID(),
		AppID:        ip.AppID(),
		Cidr:         ip.CIDR(), // Model 使用 Cidr，Domain 使用 CIDR
		Description:  ip.Description(),
		Status:       allowlistStatusDomainToEnum(ip.Status()),
		CreatedAt:    ip.CreatedAt(),
		CreatedBy:    ip.CreatedBy(),
		UpdatedAt:    ip.UpdatedAt(),
		UpdatedBy:    ip.UpdatedBy(),
		RevokedAt:    ip.RevokedAt(),
		RevokedBy:    ip.RevokedBy(),
		RevokeReason: ip.RevokeReason(),
	}
}

// IPAllowlistModelToDomain 将 Model IpAllowlist 转换为 Domain IPAllowlist
func IPAllowlistModelToDomain(m *models.IpAllowlist) *ip_allowlist.IPAllowlist {
	if m == nil {
		return nil
	}

	state := ip_allowlist.IPAllowlistState{
		ID:           m.ID,
		RuleID:       m.RuleID,
		AppID:        m.AppID,
		CIDR:         m.Cidr, // Model 使用 Cidr，Domain 使用 CIDR
		Description:  m.Description,
		Status:       allowlistStatusEnumToDomain(m.Status),
		CreatedAt:    m.CreatedAt,
		CreatedBy:    m.CreatedBy,
		UpdatedAt:    m.UpdatedAt,
		UpdatedBy:    m.UpdatedBy,
		RevokedAt:    m.RevokedAt,
		RevokedBy:    m.RevokedBy,
		RevokeReason: m.RevokeReason,
	}

	return ip_allowlist.NewIPAllowlistFromState(state)
}

// IPAllowlistModelToUpdates 将 Model IpAllowlist 转换为更新字段映射
func IPAllowlistModelToUpdates(m *models.IpAllowlist) map[string]any {
	return map[string]any{
		models.IpAllowlistCols.Cidr:         m.Cidr,
		models.IpAllowlistCols.Description: m.Description,
		models.IpAllowlistCols.Status:      m.Status,
		models.IpAllowlistCols.UpdatedAt:   m.UpdatedAt,
		models.IpAllowlistCols.UpdatedBy:   m.UpdatedBy,
		models.IpAllowlistCols.RevokedAt:   m.RevokedAt,
		models.IpAllowlistCols.RevokedBy:   m.RevokedBy,
		models.IpAllowlistCols.RevokeReason: m.RevokeReason,
	}
}

// 枚举转换辅助函数

// AllowlistStatusDomainToEnum 将 Domain AllowlistStatus 转换为 Enum AllowlistStatus（导出供其他包使用）
func AllowlistStatusDomainToEnum(as ip_allowlist.AllowlistStatus) enums.ClientsAllowlistStatus {
	return allowlistStatusDomainToEnum(as)
}

func allowlistStatusDomainToEnum(as ip_allowlist.AllowlistStatus) enums.ClientsAllowlistStatus {
	switch as {
	case ip_allowlist.AllowlistStatusActive:
		return enums.ClientsAllowlistStatusActive
	case ip_allowlist.AllowlistStatusDisabled:
		return enums.ClientsAllowlistStatusDisabled
	case ip_allowlist.AllowlistStatusRevoked:
		return enums.ClientsAllowlistStatusRevoked
	default:
		return enums.ClientsAllowlistStatusActive
	}
}

func allowlistStatusEnumToDomain(as enums.ClientsAllowlistStatus) ip_allowlist.AllowlistStatus {
	switch as {
	case enums.ClientsAllowlistStatusActive:
		return ip_allowlist.AllowlistStatusActive
	case enums.ClientsAllowlistStatusDisabled:
		return ip_allowlist.AllowlistStatusDisabled
	case enums.ClientsAllowlistStatusRevoked:
		return ip_allowlist.AllowlistStatusRevoked
	default:
		return ip_allowlist.AllowlistStatusActive
	}
}
