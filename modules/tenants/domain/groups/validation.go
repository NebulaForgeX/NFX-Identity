package groups

import "github.com/google/uuid"

func (g *Group) Validate() error {
	if g.GroupID() == "" {
		return ErrGroupIDRequired
	}
	if g.Name() == "" {
		return ErrNameRequired
	}
	if g.TenantID() == uuid.Nil {
		return ErrTenantIDRequired
	}
	validTypes := map[GroupType]struct{}{
		GroupTypeDepartment: {},
		GroupTypeTeam:       {},
		GroupTypeGroup:      {},
		GroupTypeOther:      {},
	}
	if _, ok := validTypes[g.Type()]; !ok {
		return ErrInvalidGroupType
	}
	return nil
}
