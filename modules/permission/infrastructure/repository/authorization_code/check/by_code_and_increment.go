package check

import (
	"context"
	authorizationCodeDomain "nfxid/modules/permission/domain/authorization_code"
	authorizationCodeDomainErrors "nfxid/modules/permission/domain/authorization_code/errors"
	"nfxid/modules/permission/infrastructure/rdb/models"
	authorizationCodeMapper "nfxid/modules/permission/infrastructure/repository/mapper"

	"gorm.io/gorm"
)

// ByCodeAndIncrement 根据 Code 检查 AuthorizationCode 是否有效，如果有效则增加 used_count
func (h *Handler) ByCodeAndIncrement(ctx context.Context, code string) (*authorizationCodeDomain.AuthorizationCode, error) {
	var model models.AuthorizationCode

	// 使用事务确保原子性操作
	err := h.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 先查找授权码
		if err := tx.Table("permission.authorization_codes").
			Where("code = ? AND deleted_at IS NULL", code).
			First(&model).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return authorizationCodeDomainErrors.ErrAuthorizationCodeNotFound
			}
			return err
		}

		// 转换为 domain entity
		domainEntity := authorizationCodeMapper.AuthorizationCodeModelToDomain(&model)

		// 检查是否可用（检查是否激活、是否过期、是否超过最大使用次数）
		if !domainEntity.IsAvailable() {
			// 根据具体情况返回不同的错误
			if !domainEntity.IsActive() {
				return authorizationCodeDomainErrors.ErrAuthorizationCodeInactive
			}
			if domainEntity.Editable().UsedCount >= domainEntity.Editable().MaxUses {
				return authorizationCodeDomainErrors.ErrAuthorizationCodeAlreadyUsed
			}
			// 可能已过期，但 IsAvailable 已经处理了
			return authorizationCodeDomainErrors.ErrAuthorizationCodeExpired
		}

		// 使用授权码（增加 used_count）
		if err := domainEntity.Use(); err != nil {
			return err
		}

		// 更新数据库
		updateModel := authorizationCodeMapper.AuthorizationCodeDomainToModel(domainEntity)
		if err := tx.Table("permission.authorization_codes").
			Where("id = ?", model.ID).
			Updates(map[string]interface{}{
				"used_count": updateModel.UsedCount,
				"updated_at": updateModel.UpdatedAt,
			}).Error; err != nil {
			return err
		}

		model = *updateModel
		return nil
	})

	if err != nil {
		return nil, err
	}

	// 返回更新后的 domain entity
	return authorizationCodeMapper.AuthorizationCodeModelToDomain(&model), nil
}
