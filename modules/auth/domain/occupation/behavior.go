package occupation

import (
	"time"

	occupationErrors "nebulaid/modules/auth/domain/occupation/errors"
)

func (o *Occupation) EnsureEditable(ed OccupationEditable) error {
	if err := ed.Validate(); err != nil {
		return err
	}
	if o.DeletedAt() != nil {
		return occupationErrors.ErrOccupationNotFound
	}
	return nil
}

func (o *Occupation) Update(ed OccupationEditable) error {
	if err := o.EnsureEditable(ed); err != nil {
		return err
	}

	o.state.Editable = ed
	o.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (o *Occupation) Delete() error {
	if o.DeletedAt() != nil {
		return nil // idempotent
	}

	now := time.Now().UTC()
	o.state.DeletedAt = &now
	o.state.UpdatedAt = now
	return nil
}
