package badge

import "time"

func (b *Badge) EnsureEditable(e BadgeEditable) error {
	if err := e.Validate(); err != nil {
		return err
	}
	return nil
}

func (b *Badge) Update(e BadgeEditable) error {
	if err := b.EnsureEditable(e); err != nil {
		return err
	}

	b.state.Editable = e
	return nil
}

func (b *Badge) Delete() error {
	if b.DeletedAt() != nil {
		return nil // idempotent
	}

	now := time.Now().UTC()
	b.state.DeletedAt = &now
	return nil
}
