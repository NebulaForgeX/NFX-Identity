package application_role_assignments

import (
	"context"

	"github.com/google/uuid"
)

type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Delete Delete
}

type Create interface {
	New(ctx context.Context, a *ApplicationRoleAssignment) error
}

type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*ApplicationRoleAssignment, error)
	ByUserIDAndApplicationID(ctx context.Context, userID, applicationID uuid.UUID) (*ApplicationRoleAssignment, error)
	ByApplicationID(ctx context.Context, applicationID uuid.UUID) ([]*ApplicationRoleAssignment, error)
	ByUserID(ctx context.Context, userID uuid.UUID) ([]*ApplicationRoleAssignment, error)
}

type Check interface {
	ByUserIDAndApplicationID(ctx context.Context, userID, applicationID uuid.UUID) (bool, error)
}

type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
}
