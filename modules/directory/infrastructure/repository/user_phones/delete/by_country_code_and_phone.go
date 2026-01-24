package delete

import (
	"context"
	"nfxid/modules/directory/infrastructure/rdb/models"
)

// ByCountryCodeAndPhone 根据国家代码和手机号删除 UserPhone，实现 user_phones.Delete 接口
func (h *Handler) ByCountryCodeAndPhone(ctx context.Context, countryCode, phone string) error {
	// 移除 countryCode 中可能存在的 + 号
	normalizedCC := countryCode
	if len(normalizedCC) > 0 && normalizedCC[0] == '+' {
		normalizedCC = normalizedCC[1:]
	}

	return h.db.WithContext(ctx).
		Where("phone = ? AND (country_code = ? OR country_code = ? OR REPLACE(country_code, '+', '') = ?)",
			phone, countryCode, normalizedCC, normalizedCC).
		Delete(&models.UserPhone{}).Error
}
