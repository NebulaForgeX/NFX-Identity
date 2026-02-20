package tenant_role_assignments

import (
	"time"

	"github.com/google/uuid"
)

type NewTenantRoleAssignmentParams struct {
	UserID       uuid.UUID
	TenantID     uuid.UUID
	TenantRoleID uuid.UUID
	AssignedBy   *uuid.UUID
}

func NewTenantRoleAssignment(p NewTenantRoleAssignmentParams) (*TenantRoleAssignment, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	at := time.Now().UTC()
	return &TenantRoleAssignment{
		state: TenantRoleAssignmentState{
			ID:           id,
			UserID:       p.UserID,
			TenantID:     p.TenantID,
			TenantRoleID: p.TenantRoleID,
			AssignedAt:   at,
			AssignedBy:   p.AssignedBy,
		},
	}, nil
}

func NewTenantRoleAssignmentFromState(st TenantRoleAssignmentState) *TenantRoleAssignment {
	return &TenantRoleAssignment{state: st}
}
