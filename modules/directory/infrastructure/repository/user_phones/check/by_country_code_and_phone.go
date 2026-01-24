package check

import (
	"context"
	"nfxid/modules/directory/infrastructure/rdb/models"
)

// ByCountryCodeAndPhone 根据国家代码和手机号检查 UserPhone 是否存在，实现 user_phones.Check 接口
func (h *Handler) ByCountryCodeAndPhone(ctx context.Context, countryCode, phone string) (bool, error) {
	// 移除 countryCode 中可能存在的 + 号
	normalizedCC := countryCode
	if len(normalizedCC) > 0 && normalizedCC[0] == '+' {
		normalizedCC = normalizedCC[1:]
	}

	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.UserPhone{}).
		Where("phone = ? AND (country_code = ? OR country_code = ? OR REPLACE(country_code, '+', '') = ?)",
			phone, countryCode, normalizedCC, normalizedCC).
		Count(&count).Error
	return count > 0, err
}
