package education

import (
	"time"

	educationErrors "nebulaid/modules/auth/domain/education/errors"

	"github.com/google/uuid"
)

type NewEducationParams struct {
	ProfileID uuid.UUID
	Editable  EducationEditable
}

func NewEducation(p NewEducationParams) (*Education, error) {
	if p.ProfileID == uuid.Nil {
		return nil, educationErrors.ErrEducationProfileIDRequired
	}
	if err := p.Editable.Validate(); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewEducationFromState(EducationState{
		ID:        id,
		ProfileID: p.ProfileID,
		Editable:  p.Editable,
		CreatedAt: now,
		UpdatedAt: now,
	}), nil
}

func NewEducationFromState(st EducationState) *Education {
	return &Education{state: st}
}
