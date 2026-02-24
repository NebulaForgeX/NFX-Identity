package application_role_assignments

import (
	accessErr "nfxid/errors/src/access"

	"github.com/google/uuid"
)

func (a *ApplicationRoleAssignment) Validate() error {
	if a.UserID() == uuid.Nil || a.ApplicationID() == uuid.Nil || a.ApplicationRoleID() == uuid.Nil {
		return accessErr.ErrApplicationRoleAssignmentNotFound
	}
	return nil
}
