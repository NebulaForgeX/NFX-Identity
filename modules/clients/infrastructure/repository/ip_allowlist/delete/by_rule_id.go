package delete

import (
	"context"
	clientsErr "nfxid/errors/src/clients"
	"nfxid/modules/clients/infrastructure/rdb/models"
)

// ByRuleID 根据 RuleID 删除 IPAllowlist，实现 ip_allowlist.Delete 接口
func (h *Handler) ByRuleID(ctx context.Context, ruleID string) error {
	result := h.db.WithContext(ctx).
		Where("rule_id = ?", ruleID).
		Delete(&models.IpAllowlist{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return clientsErr.ErrIPAllowlistNotFound
	}
	return nil
}
