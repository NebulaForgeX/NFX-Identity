package mapper

import (
	"nfxid/enums"
	"nfxid/modules/directory/domain/users"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/pkgs/utils/timex"
)

// UserDomainToModel 将 Domain User 转换为 Model User
func UserDomainToModel(u *users.User) *models.User {
	if u == nil {
		return nil
	}

	return &models.User{
		ID:          u.ID(),
		TenantID:    u.TenantID(),
		Username:    u.Username(),
		Status:      userStatusDomainToEnum(u.Status()),
		IsVerified:  u.IsVerified(),
		LastLoginAt: u.LastLoginAt(),
		CreatedAt:   u.CreatedAt(),
		UpdatedAt:   u.UpdatedAt(),
		DeletedAt:   timex.TimeToGormDeletedAt(u.DeletedAt()),
	}
}

// UserModelToDomain 将 Model User 转换为 Domain User
func UserModelToDomain(m *models.User) *users.User {
	if m == nil {
		return nil
	}

	state := users.UserState{
		ID:          m.ID,
		TenantID:    m.TenantID,
		Username:    m.Username,
		Status:      userStatusEnumToDomain(m.Status),
		IsVerified:  m.IsVerified,
		LastLoginAt: m.LastLoginAt,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
		DeletedAt:   timex.GormDeletedAtToTime(m.DeletedAt),
	}

	return users.NewUserFromState(state)
}

// UserModelToUpdates 将 Model User 转换为更新字段映射
func UserModelToUpdates(m *models.User) map[string]any {
	return map[string]any{
		models.UserCols.TenantID:    m.TenantID,
		models.UserCols.Username:    m.Username,
		models.UserCols.Status:      m.Status,
		models.UserCols.IsVerified:  m.IsVerified,
		models.UserCols.LastLoginAt: m.LastLoginAt,
		models.UserCols.UpdatedAt:   m.UpdatedAt,
		models.UserCols.DeletedAt:   m.DeletedAt,
	}
}

// 枚举转换辅助函数

// UserStatusDomainToEnum 将 Domain UserStatus 转换为 Enum UserStatus
func UserStatusDomainToEnum(us users.UserStatus) enums.DirectoryUserStatus {
	return userStatusDomainToEnum(us)
}

func userStatusDomainToEnum(us users.UserStatus) enums.DirectoryUserStatus {
	switch us {
	case users.UserStatusPending:
		return enums.DirectoryUserStatusPending
	case users.UserStatusActive:
		return enums.DirectoryUserStatusActive
	case users.UserStatusDeactive:
		return enums.DirectoryUserStatusDeactive
	default:
		return enums.DirectoryUserStatusPending
	}
}

func userStatusEnumToDomain(us enums.DirectoryUserStatus) users.UserStatus {
	switch us {
	case enums.DirectoryUserStatusPending:
		return users.UserStatusPending
	case enums.DirectoryUserStatusActive:
		return users.UserStatusActive
	case enums.DirectoryUserStatusDeactive:
		return users.UserStatusDeactive
	default:
		return users.UserStatusPending
	}
}
