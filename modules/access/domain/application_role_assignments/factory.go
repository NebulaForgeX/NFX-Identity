package application_role_assignments

import (
	"time"

	"github.com/google/uuid"
)

type NewApplicationRoleAssignmentParams struct {
	UserID            uuid.UUID
	ApplicationID     uuid.UUID
	ApplicationRoleID uuid.UUID
	AssignedBy        *uuid.UUID
}

func NewApplicationRoleAssignment(p NewApplicationRoleAssignmentParams) (*ApplicationRoleAssignment, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	at := time.Now().UTC()
	return &ApplicationRoleAssignment{
		state: ApplicationRoleAssignmentState{
			ID:                id,
			UserID:            p.UserID,
			ApplicationID:     p.ApplicationID,
			ApplicationRoleID: p.ApplicationRoleID,
			AssignedAt:        at,
			AssignedBy:        p.AssignedBy,
		},
	}, nil
}

func NewApplicationRoleAssignmentFromState(st ApplicationRoleAssignmentState) *ApplicationRoleAssignment {
	return &ApplicationRoleAssignment{state: st}
}
