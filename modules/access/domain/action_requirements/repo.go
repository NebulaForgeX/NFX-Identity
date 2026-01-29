package action_requirements

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
	New(ctx context.Context, ar *ActionRequirement) error
}

type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*ActionRequirement, error)
	ByActionIDAndPermissionID(ctx context.Context, actionID, permissionID uuid.UUID) (*ActionRequirement, error)
	ByActionID(ctx context.Context, actionID uuid.UUID) ([]*ActionRequirement, error)
	ByPermissionID(ctx context.Context, permissionID uuid.UUID) ([]*ActionRequirement, error)
}

type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
	ByActionIDAndPermissionID(ctx context.Context, actionID, permissionID uuid.UUID) (bool, error)
}

type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
	ByActionIDAndPermissionID(ctx context.Context, actionID, permissionID uuid.UUID) error
	ByActionID(ctx context.Context, actionID uuid.UUID) error
}
