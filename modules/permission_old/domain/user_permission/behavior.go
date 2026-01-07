package user_permission

import (
	"time"
)

func (up *UserPermission) Delete() error {
	if up.DeletedAt() != nil {
		return nil // idempotent
	}

	now := time.Now().UTC()
	up.state.DeletedAt = &now
	return nil
}

