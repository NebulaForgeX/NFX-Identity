package groups

import (
	"time"

	"github.com/google/uuid"
)

type NewGroupParams struct {
	GroupID       string
	TenantID      uuid.UUID
	Name          string
	Type          GroupType
	ParentGroupID *uuid.UUID
	Description   *string
	CreatedBy     *uuid.UUID
	Metadata      map[string]interface{}
}

func NewGroup(p NewGroupParams) (*Group, error) {
	if err := validateGroupParams(p); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	groupType := p.Type
	if groupType == "" {
		groupType = GroupTypeGroup
	}

	now := time.Now().UTC()
	return NewGroupFromState(GroupState{
		ID:            id,
		GroupID:       p.GroupID,
		TenantID:      p.TenantID,
		Name:          p.Name,
		Type:          groupType,
		ParentGroupID: p.ParentGroupID,
		Description:   p.Description,
		CreatedBy:     p.CreatedBy,
		Metadata:      p.Metadata,
		CreatedAt:     now,
		UpdatedAt:     now,
	}), nil
}

func NewGroupFromState(st GroupState) *Group {
	return &Group{state: st}
}

func validateGroupParams(p NewGroupParams) error {
	if p.GroupID == "" {
		return ErrGroupIDRequired
	}
	if p.Name == "" {
		return ErrNameRequired
	}
	if p.TenantID == uuid.Nil {
		return ErrTenantIDRequired
	}
	if p.Type != "" {
		validTypes := map[GroupType]struct{}{
			GroupTypeDepartment: {},
			GroupTypeTeam:       {},
			GroupTypeGroup:      {},
			GroupTypeOther:      {},
		}
		if _, ok := validTypes[p.Type]; !ok {
			return ErrInvalidGroupType
		}
	}
	return nil
}
