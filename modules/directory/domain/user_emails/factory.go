package user_emails

import (
	dirErr "nfxidentity/errors/src/directory"
	"time"

	"github.com/google/uuid"
)

type NewUserEmailParams struct {
	UserID            uuid.UUID
	Email             string
	IsPrimary         bool
	IsVerified        bool
	VerificationToken *string
}

func NewUserEmail(p NewUserEmailParams) (*UserEmail, error) {
	if err := validateUserEmailParams(p); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewUserEmailFromState(UserEmailState{
		ID:                id,
		UserID:            p.UserID,
		Email:             p.Email,
		IsPrimary:         p.IsPrimary,
		IsVerified:        p.IsVerified,
		VerifiedAt:        nil,
		VerificationToken: p.VerificationToken,
		CreatedAt:         now,
		UpdatedAt:         now,
	}), nil
}

func NewUserEmailFromState(st UserEmailState) *UserEmail {
	return &UserEmail{state: st}
}

func validateUserEmailParams(p NewUserEmailParams) error {
	if p.UserID == uuid.Nil {
		return dirErr.ErrUserIDRequired
	}
	if p.Email == "" {
		return dirErr.ErrEmailRequired
	}
	return nil
}
