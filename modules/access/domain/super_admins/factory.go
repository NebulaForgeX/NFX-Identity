package super_admins

import (
	"time"

	"github.com/google/uuid"
)

type NewSuperAdminParams struct {
	UserID uuid.UUID
}

func NewSuperAdmin(p NewSuperAdminParams) *SuperAdmin {
	return &SuperAdmin{
		state: SuperAdminState{
			UserID:    p.UserID,
			CreatedAt: time.Now().UTC(),
		},
	}
}

func NewSuperAdminFromState(st SuperAdminState) *SuperAdmin {
	return &SuperAdmin{state: st}
}
