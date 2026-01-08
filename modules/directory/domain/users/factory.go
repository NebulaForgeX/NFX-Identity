package users

import (
	"time"

	"github.com/google/uuid"
)

type NewUserParams struct {
	TenantID   uuid.UUID
	Username   string
	Status     UserStatus
	IsVerified bool
}

func NewUser(p NewUserParams) (*User, error) {
	if err := validateUserParams(p); err != nil {
		return nil, err
	}

	status := p.Status
	if status == "" {
		status = UserStatusPending
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewUserFromState(UserState{
		ID:         id,
		TenantID:   p.TenantID,
		Username:   p.Username,
		Status:     status,
		IsVerified: p.IsVerified,
		CreatedAt:  now,
		UpdatedAt:  now,
	}), nil
}

func NewUserFromState(st UserState) *User {
	return &User{state: st}
}

func validateUserParams(p NewUserParams) error {
	if p.Username == "" {
		return ErrUsernameRequired
	}
	if p.TenantID == uuid.Nil {
		return ErrTenantIDRequired
	}
	if p.Status != "" {
		validStatuses := map[UserStatus]struct{}{
			UserStatusPending:  {},
			UserStatusActive:   {},
			UserStatusDeactive: {},
		}
		if _, ok := validStatuses[p.Status]; !ok {
			return ErrInvalidUserStatus
		}
	}
	return nil
}
