package mapper

import (
	"nfxid/modules/access/domain/application_role_assignments"
	"nfxid/modules/access/infrastructure/rdb/models"
)

func ApplicationRoleAssignmentDomainToModel(a *application_role_assignments.ApplicationRoleAssignment) *models.ApplicationRoleAssignment {
	if a == nil {
		return nil
	}
	return &models.ApplicationRoleAssignment{
		ID: a.ID(), UserID: a.UserID(), ApplicationID: a.ApplicationID(),
		ApplicationRoleID: a.ApplicationRoleID(), AssignedAt: a.AssignedAt(), AssignedBy: a.AssignedBy(),
	}
}

func ApplicationRoleAssignmentModelToDomain(m *models.ApplicationRoleAssignment) *application_role_assignments.ApplicationRoleAssignment {
	if m == nil {
		return nil
	}
	return application_role_assignments.NewApplicationRoleAssignmentFromState(application_role_assignments.ApplicationRoleAssignmentState{
		ID: m.ID, UserID: m.UserID, ApplicationID: m.ApplicationID,
		ApplicationRoleID: m.ApplicationRoleID, AssignedAt: m.AssignedAt, AssignedBy: m.AssignedBy,
	})
}
