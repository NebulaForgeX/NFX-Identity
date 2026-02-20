package application_roles

import (
	"time"

	"github.com/google/uuid"
)

func nowUTC() time.Time { return time.Now().UTC() }

func (r *ApplicationRole) Validate() error {
	if r.RoleKey() == "" {
		return ErrApplicationRoleRoleKeyRequired
	}
	if r.ApplicationID() == uuid.Nil {
		return ErrApplicationRoleApplicationIDRequired
	}
	return nil
}
