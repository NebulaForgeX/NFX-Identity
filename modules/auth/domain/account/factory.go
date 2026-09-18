package account

import (
	"time"

	"nfxidentity/enums"

	"github.com/google/uuid"
)

type NewAccountParams struct {
	AccountStatus  enums.AuthAccountStatus
	SignupPlatform enums.AuthSignupPlatform
}

func NewAccount(p NewAccountParams) (*Account, error) {
	status := p.AccountStatus
	if status == "" {
		status = enums.AuthAccountStatusActive
	}
	if err := validateAccountStatus(status); err != nil {
		return nil, err
	}
	if err := validateSignupPlatform(p.SignupPlatform); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	now := time.Now().UTC()
	return NewAccountFromState(AccountState{
		ID:             id,
		AccountStatus:  status,
		SignupPlatform: p.SignupPlatform,
		CreatedAt:      now,
		UpdatedAt:      now,
	}), nil
}
