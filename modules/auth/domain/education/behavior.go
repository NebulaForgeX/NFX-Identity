package education

import (
	"time"

	educationErrors "nebulaid/modules/auth/domain/education/errors"
)

func (e *Education) EnsureEditable(ed EducationEditable) error {
	if err := ed.Validate(); err != nil {
		return err
	}
	if e.DeletedAt() != nil {
		return educationErrors.ErrEducationNotFound
	}
	return nil
}

func (e *Education) Update(ed EducationEditable) error {
	if err := e.EnsureEditable(ed); err != nil {
		return err
	}

	e.state.Editable = ed
	e.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (e *Education) Delete() error {
	if e.DeletedAt() != nil {
		return nil // idempotent
	}

	now := time.Now().UTC()
	e.state.DeletedAt = &now
	e.state.UpdatedAt = now
	return nil
}
