package occupation

import (
	"time"

	occupationErrors "nfxid/modules/auth/domain/profile_occupation/errors"

	"github.com/google/uuid"
)

type NewOccupationParams struct {
	ProfileID uuid.UUID
	Editable  OccupationEditable
}

func NewOccupation(p NewOccupationParams) (*Occupation, error) {
	if p.ProfileID == uuid.Nil {
		return nil, occupationErrors.ErrOccupationProfileIDRequired
	}
	if err := p.Editable.Validate(); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewOccupationFromState(OccupationState{
		ID:        id,
		ProfileID: p.ProfileID,
		Editable:  p.Editable,
		CreatedAt: now,
		UpdatedAt: now,
	}), nil
}

func NewOccupationFromState(st OccupationState) *Occupation {
	return &Occupation{state: st}
}
