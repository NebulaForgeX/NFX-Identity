package application_role_assignments

import "github.com/google/uuid"

func (a *ApplicationRoleAssignment) Validate() error {
	if a.UserID() == uuid.Nil || a.ApplicationID() == uuid.Nil || a.ApplicationRoleID() == uuid.Nil {
		return ErrApplicationRoleAssignmentNotFound
	}
	return nil
}
