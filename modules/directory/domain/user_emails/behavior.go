package user_emails

import (
	dirErr "nfxid/errors/src/directory"
	"time"
)

func (ue *UserEmail) SetPrimary() error {
	if ue.DeletedAt() != nil {
		return dirErr.ErrUserEmailNotFound
	}
	ue.state.IsPrimary = true
	ue.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (ue *UserEmail) SetNotPrimary() error {
	if ue.DeletedAt() != nil {
		return dirErr.ErrUserEmailNotFound
	}
	ue.state.IsPrimary = false
	ue.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (ue *UserEmail) Verify() error {
	if ue.DeletedAt() != nil {
		return dirErr.ErrUserEmailNotFound
	}

	now := time.Now().UTC()
	ue.state.IsVerified = true
	ue.state.VerifiedAt = &now
	ue.state.VerificationToken = nil
	ue.state.UpdatedAt = now
	return nil
}

func (ue *UserEmail) UpdateVerificationToken(token *string) error {
	if ue.DeletedAt() != nil {
		return dirErr.ErrUserEmailNotFound
	}
	ue.state.VerificationToken = token
	ue.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (ue *UserEmail) Delete() error {
	if ue.DeletedAt() != nil {
		return nil // idempotent
	}

	now := time.Now().UTC()
	ue.state.DeletedAt = &now
	ue.state.UpdatedAt = now
	return nil
}
