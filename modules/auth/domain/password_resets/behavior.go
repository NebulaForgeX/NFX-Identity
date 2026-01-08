package password_resets

import (
	"time"
)

func (pr *PasswordReset) MarkAsUsed() error {
	if pr.Status() == ResetStatusUsed {
		return ErrResetAlreadyUsed
	}
	if pr.IsExpired() {
		return ErrResetExpired
	}

	now := time.Now().UTC()
	pr.state.UsedAt = &now
	pr.state.Status = ResetStatusUsed
	pr.state.UpdatedAt = now
	return nil
}

func (pr *PasswordReset) IncrementAttemptCount() error {
	pr.state.AttemptCount++
	pr.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (pr *PasswordReset) UpdateStatus(status ResetStatus) error {
	validStatuses := map[ResetStatus]struct{}{
		ResetStatusIssued:  {},
		ResetStatusUsed:    {},
		ResetStatusExpired: {},
		ResetStatusRevoked: {},
	}
	if _, ok := validStatuses[status]; !ok {
		return ErrInvalidResetStatus
	}

	pr.state.Status = status
	pr.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (pr *PasswordReset) IsExpired() bool {
	return time.Now().UTC().After(pr.ExpiresAt())
}

func (pr *PasswordReset) IsValid() bool {
	return pr.Status() == ResetStatusIssued && !pr.IsExpired()
}
