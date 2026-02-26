package application_roles

import (
	accessErr "nfxidentity/errors/src/access"
	"time"

	"github.com/google/uuid"
)

func nowUTC() time.Time { return time.Now().UTC() }

func (r *ApplicationRole) Validate() error {
	if r.RoleKey() == "" {
		return accessErr.ErrApplicationRoleRoleKeyRequired
	}
	if r.ApplicationID() == uuid.Nil {
		return accessErr.ErrApplicationRoleApplicationIDRequired
	}
	return nil
}
