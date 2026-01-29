package action_requirements

import "github.com/google/uuid"

func (ar *ActionRequirement) Validate() error {
	if ar.ActionID() == uuid.Nil {
		return ErrActionIDRequired
	}
	if ar.PermissionID() == uuid.Nil {
		return ErrPermissionIDRequired
	}
	return nil
}
