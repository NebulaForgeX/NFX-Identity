package groups

import (
	tenantsErr "nfxid/errors/src/tenants"

	"github.com/google/uuid"
)

func (g *Group) Validate() error {
	if g.GroupID() == "" {
		return tenantsErr.ErrGroupIDRequired
	}
	if g.Name() == "" {
		return tenantsErr.ErrNameRequired
	}
	if g.TenantID() == uuid.Nil {
		return tenantsErr.ErrTenantIDRequired
	}
	validTypes := map[GroupType]struct{}{
		GroupTypeDepartment: {},
		GroupTypeTeam:       {},
		GroupTypeGroup:      {},
		GroupTypeOther:      {},
	}
	if _, ok := validTypes[g.Type()]; !ok {
		return tenantsErr.ErrInvalidGroupType
	}
	return nil
}
