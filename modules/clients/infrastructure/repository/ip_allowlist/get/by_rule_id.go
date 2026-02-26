package get

import (
	"context"
	"errors"
	clientsErr "nfxidentity/errors/src/clients"
	"nfxidentity/modules/clients/domain/ip_allowlist"
	"nfxidentity/modules/clients/infrastructure/rdb/models"
	"nfxidentity/modules/clients/infrastructure/repository/ip_allowlist/mapper"

	"gorm.io/gorm"
)

// ByRuleID 根据 RuleID 获取 IPAllowlist，实现 ip_allowlist.Get 接口
func (h *Handler) ByRuleID(ctx context.Context, ruleID string) (*ip_allowlist.IPAllowlist, error) {
	var m models.IpAllowlist
	if err := h.db.WithContext(ctx).Where("rule_id = ?", ruleID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, clientsErr.ErrIPAllowlistNotFound
		}
		return nil, err
	}
	return mapper.IPAllowlistModelToDomain(&m), nil
}
