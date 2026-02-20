package mapper

import (
	"nfxid/modules/access/domain/application_roles"
	"nfxid/modules/access/infrastructure/rdb/models"
)

func ApplicationRoleDomainToModel(r *application_roles.ApplicationRole) *models.ApplicationRole {
	if r == nil {
		return nil
	}
	return &models.ApplicationRole{
		ID: r.ID(), ApplicationID: r.ApplicationID(), RoleKey: r.RoleKey(),
		Name: r.Name(), CreatedAt: r.CreatedAt(),
	}
}

func ApplicationRoleModelToDomain(m *models.ApplicationRole) *application_roles.ApplicationRole {
	if m == nil {
		return nil
	}
	return application_roles.NewApplicationRoleFromState(application_roles.ApplicationRoleState{
		ID: m.ID, ApplicationID: m.ApplicationID, RoleKey: m.RoleKey,
		Name: m.Name, CreatedAt: m.CreatedAt,
	})
}
