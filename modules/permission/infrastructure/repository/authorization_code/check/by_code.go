package check

import (
	"context"
)

// ByCode 根据 Code 检查 AuthorizationCode 是否存在，实现 authorizationCodeDomain.Check 接口
func (h *Handler) ByCode(ctx context.Context, code string) (bool, error) {
	var count int64
	if err := h.db.WithContext(ctx).Table("permission.authorization_codes").Where("code = ?", code).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
