package get

import (
	"context"
	"errors"
	"nfxid/modules/directory/domain/user_phones"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/modules/directory/infrastructure/repository/user_phones/mapper"

	"gorm.io/gorm"
)

// ByCountryCodeAndPhone 根据国家代码和手机号获取 UserPhone
// countryCode 可能是 "+1", "1", "+86", "86" 等格式
// phone 是纯手机号，不包含国家代码
func (h *Handler) ByCountryCodeAndPhone(ctx context.Context, countryCode, phone string) (*user_phones.UserPhone, error) {
	var m models.UserPhone

	// 移除 countryCode 中可能存在的 + 号
	normalizedCC := countryCode
	if len(normalizedCC) > 0 && normalizedCC[0] == '+' {
		normalizedCC = normalizedCC[1:]
	}

	// 查询：country_code 匹配（支持带+号或不带+号），phone 精确匹配
	query := h.db.WithContext(ctx).Where(
		"phone = ? AND (country_code = ? OR country_code = ? OR REPLACE(country_code, '+', '') = ?)",
		phone, countryCode, normalizedCC, normalizedCC,
	)

	if err := query.First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, user_phones.ErrUserPhoneNotFound
		}
		return nil, err
	}

	return mapper.UserPhoneModelToDomain(&m), nil
}
