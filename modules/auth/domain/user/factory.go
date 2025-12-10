package user

import (
	"time"

	"github.com/google/uuid"
)

type NewUserParams struct {
	Editable UserEditable
	Status   string // pending, active, deactive
}

func NewUser(p NewUserParams) (*User, error) {
	if err := p.Editable.Validate(); err != nil {
		return nil, err
	}

	status := p.Status
	if status == "" {
		status = "active"
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewUserFromState(UserState{
		ID:         id,
		Editable:   p.Editable,
		Status:     status,
		IsVerified: false,
		CreatedAt:  now,
		UpdatedAt:  now,
	}), nil
}

func NewUserFromState(st UserState) *User {
	return &User{state: st}
}

