package badge

import "github.com/google/uuid"

type NewBadgeParams struct {
	Editable BadgeEditable
}

func NewBadge(p NewBadgeParams) (*Badge, error) {
	if err := p.Editable.Validate(); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	return NewBadgeFromState(BadgeState{
		ID:       id,
		Editable: p.Editable,
	}), nil
}

func NewBadgeFromState(st BadgeState) *Badge {
	return &Badge{state: st}
}

