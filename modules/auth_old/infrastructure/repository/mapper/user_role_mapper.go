package mapper

import (
	"nfxid/modules/auth/domain/user_role"
	"nfxid/modules/auth/infrastructure/rdb/models"
)

func UserRoleDomainToModel(ur *user_role.UserRole) *models.UserRole {
	if ur == nil {
		return nil
	}

	return &models.UserRole{
		ID:        ur.ID(),
		UserID:    ur.UserID(),
		RoleID:    ur.RoleID(),
		CreatedAt: ur.CreatedAt(),
	}
}

func UserRoleModelToDomain(m *models.UserRole) *user_role.UserRole {
	if m == nil {
		return nil
	}

	state := user_role.UserRoleState{
		ID:        m.ID,
		UserID:    m.UserID,
		RoleID:    m.RoleID,
		CreatedAt: m.CreatedAt,
	}

	return user_role.NewUserRoleFromState(state)
}

