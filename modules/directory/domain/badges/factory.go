package badges

import (
	"time"

	"github.com/google/uuid"
)

type NewBadgeParams struct {
	Name        string
	Description *string
	IconURL     *string
	Color       *string
	Category    *string
	IsSystem    bool
}

func NewBadge(p NewBadgeParams) (*Badge, error) {
	if err := validateBadgeParams(p); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewBadgeFromState(BadgeState{
		ID:          id,
		Name:        p.Name,
		Description: p.Description,
		IconURL:     p.IconURL,
		Color:       p.Color,
		Category:    p.Category,
		IsSystem:    p.IsSystem,
		CreatedAt:   now,
		UpdatedAt:   now,
	}), nil
}

func NewBadgeFromState(st BadgeState) *Badge {
	return &Badge{state: st}
}

func validateBadgeParams(p NewBadgeParams) error {
	if p.Name == "" {
		return ErrNameRequired
	}
	return nil
}
