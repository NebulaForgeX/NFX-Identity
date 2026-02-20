package tenant_role_assignments

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
	New(ctx context.Context, a *TenantRoleAssignment) error
}

type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*TenantRoleAssignment, error)
	ByUserIDAndTenantID(ctx context.Context, userID, tenantID uuid.UUID) (*TenantRoleAssignment, error)
	ByTenantID(ctx context.Context, tenantID uuid.UUID) ([]*TenantRoleAssignment, error)
	ByUserID(ctx context.Context, userID uuid.UUID) ([]*TenantRoleAssignment, error)
}

type Check interface {
	ByUserIDAndTenantID(ctx context.Context, userID, tenantID uuid.UUID) (bool, error)
}

type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
}
