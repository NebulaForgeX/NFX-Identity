package super_admins

import (
	"time"

	"github.com/google/uuid"
)

type SuperAdmin struct {
	state SuperAdminState
}

type SuperAdminState struct {
	UserID    uuid.UUID
	CreatedAt time.Time
}

func (s *SuperAdmin) UserID() uuid.UUID     { return s.state.UserID }
func (s *SuperAdmin) CreatedAt() time.Time   { return s.state.CreatedAt }
