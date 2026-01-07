package mapper

import (
	"nfxid/modules/permission/domain/user_permission"
	"nfxid/modules/permission/infrastructure/rdb/models"
	"nfxid/pkgs/utils/timex"
)

func UserPermissionDomainToModel(up *user_permission.UserPermission) *models.UserPermission {
	if up == nil {
		return nil
	}

	return &models.UserPermission{
		ID:           up.ID(),
		UserID:       up.UserID(),
		PermissionID: up.PermissionID(),
		CreatedAt:    up.CreatedAt(),
		DeletedAt:    timex.TimeToGormDeletedAt(up.DeletedAt()),
	}
}

func UserPermissionModelToDomain(m *models.UserPermission) *user_permission.UserPermission {
	if m == nil {
		return nil
	}

	state := user_permission.UserPermissionState{
		ID:           m.ID,
		UserID:       m.UserID,
		PermissionID: m.PermissionID,
		CreatedAt:    m.CreatedAt,
		DeletedAt:    timex.GormDeletedAtToTime(m.DeletedAt),
	}

	return user_permission.NewUserPermissionFromState(state)
}

