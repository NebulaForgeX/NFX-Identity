package profile

import (
	"time"

	profileErrors "nebulaid/modules/auth/domain/profile/errors"

	"github.com/google/uuid"
)

type NewProfileParams struct {
	UserID   uuid.UUID
	Editable ProfileEditable
}

func NewProfile(p NewProfileParams) (*Profile, error) {
	if p.UserID == uuid.Nil {
		return nil, profileErrors.ErrProfileUserIDRequired
	}
	if err := p.Editable.Validate(); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewProfileFromState(ProfileState{
		ID:        id,
		UserID:    p.UserID,
		Editable:  p.Editable,
		CreatedAt: now,
		UpdatedAt: now,
	}), nil
}

func NewProfileFromState(st ProfileState) *Profile {
	return &Profile{state: st}
}
