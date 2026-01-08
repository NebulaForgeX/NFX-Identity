package badges

import (
	"time"
)

func (b *Badge) Update(name string, description, iconURL, color, category *string) error {
	if b.DeletedAt() != nil {
		return ErrBadgeNotFound
	}
	if name == "" {
		return ErrNameRequired
	}

	b.state.Name = name
	if description != nil {
		b.state.Description = description
	}
	if iconURL != nil {
		b.state.IconURL = iconURL
	}
	if color != nil {
		b.state.Color = color
	}
	if category != nil {
		b.state.Category = category
	}

	b.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (b *Badge) Delete() error {
	if b.IsSystem() {
		return ErrBadgeNotFound // system badges cannot be deleted
	}
	if b.DeletedAt() != nil {
		return nil // idempotent
	}

	now := time.Now().UTC()
	b.state.DeletedAt = &now
	b.state.UpdatedAt = now
	return nil
}
