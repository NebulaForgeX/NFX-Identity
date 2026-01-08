package groups

import (
	"time"
)

func (g *Group) Update(name string, groupType GroupType, parentGroupID *uuid.UUID, description *string, metadata map[string]interface{}) error {
	if g.DeletedAt() != nil {
		return ErrGroupNotFound
	}
	if name == "" {
		return ErrNameRequired
	}
	if groupType != "" {
		validTypes := map[GroupType]struct{}{
			GroupTypeDepartment: {},
			GroupTypeTeam:       {},
			GroupTypeGroup:      {},
			GroupTypeOther:      {},
		}
		if _, ok := validTypes[groupType]; !ok {
			return ErrInvalidGroupType
		}
		g.state.Type = groupType
	}

	g.state.Name = name
	if parentGroupID != nil {
		g.state.ParentGroupID = parentGroupID
	}
	if description != nil {
		g.state.Description = description
	}
	if metadata != nil {
		g.state.Metadata = metadata
	}

	g.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (g *Group) Delete() error {
	if g.DeletedAt() != nil {
		return nil // idempotent
	}

	now := time.Now().UTC()
	g.state.DeletedAt = &now
	g.state.UpdatedAt = now
	return nil
}
