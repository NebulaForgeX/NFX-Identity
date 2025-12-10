package profile

import profileErrors "nebulaid/modules/auth/domain/profile/errors"

func (e *ProfileEditable) Validate() error {
	// Profile 的验证规则相对宽松，因为大部分字段都是可选的
	// 只需要确保关键字段的格式正确即可

	if e.Nickname != nil && *e.Nickname != "" {
		if len(*e.Nickname) > 50 {
			return profileErrors.ErrProfileNotFound // 可以添加更具体的错误
		}
	}

	if e.DisplayName != nil && *e.DisplayName != "" {
		if len(*e.DisplayName) > 100 {
			return profileErrors.ErrProfileNotFound
		}
	}

	if e.Bio != nil && *e.Bio != "" {
		if len(*e.Bio) > 1000 {
			return profileErrors.ErrProfileNotFound
		}
	}

	return nil
}
