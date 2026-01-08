package user_phones

import (
	"time"
)

func (up *UserPhone) SetPrimary() error {
	if up.DeletedAt() != nil {
		return ErrUserPhoneNotFound
	}
	up.state.IsPrimary = true
	up.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (up *UserPhone) SetNotPrimary() error {
	if up.DeletedAt() != nil {
		return ErrUserPhoneNotFound
	}
	up.state.IsPrimary = false
	up.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (up *UserPhone) Verify() error {
	if up.DeletedAt() != nil {
		return ErrUserPhoneNotFound
	}

	now := time.Now().UTC()
	up.state.IsVerified = true
	up.state.VerifiedAt = &now
	up.state.VerificationCode = nil
	up.state.VerificationExpiresAt = nil
	up.state.UpdatedAt = now
	return nil
}

func (up *UserPhone) UpdateVerificationCode(code string, expiresAt time.Time) error {
	if up.DeletedAt() != nil {
		return ErrUserPhoneNotFound
	}
	up.state.VerificationCode = &code
	up.state.VerificationExpiresAt = &expiresAt
	up.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (up *UserPhone) IsVerificationCodeExpired() bool {
	if up.VerificationExpiresAt() == nil {
		return true
	}
	return time.Now().UTC().After(*up.VerificationExpiresAt())
}

func (up *UserPhone) Delete() error {
	if up.DeletedAt() != nil {
		return nil // idempotent
	}

	now := time.Now().UTC()
	up.state.DeletedAt = &now
	up.state.UpdatedAt = now
	return nil
}
