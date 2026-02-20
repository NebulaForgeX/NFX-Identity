package mapper

import (
	"nfxid/modules/access/domain/tenant_role_assignments"
	"nfxid/modules/access/infrastructure/rdb/models"
)

func TenantRoleAssignmentDomainToModel(a *tenant_role_assignments.TenantRoleAssignment) *models.TenantRoleAssignment {
	if a == nil {
		return nil
	}
	return &models.TenantRoleAssignment{
		ID: a.ID(), UserID: a.UserID(), TenantID: a.TenantID(), TenantRoleID: a.TenantRoleID(),
		AssignedAt: a.AssignedAt(), AssignedBy: a.AssignedBy(),
	}
}

func TenantRoleAssignmentModelToDomain(m *models.TenantRoleAssignment) *tenant_role_assignments.TenantRoleAssignment {
	if m == nil {
		return nil
	}
	return tenant_role_assignments.NewTenantRoleAssignmentFromState(tenant_role_assignments.TenantRoleAssignmentState{
		ID: m.ID, UserID: m.UserID, TenantID: m.TenantID, TenantRoleID: m.TenantRoleID,
		AssignedAt: m.AssignedAt, AssignedBy: m.AssignedBy,
	})
}
