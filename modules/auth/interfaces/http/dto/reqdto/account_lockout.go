package reqdto

import (
	accountLockoutAppCommands "nfxid/modules/auth/application/account_lockouts/commands"
	accountLockoutDomain "nfxid/modules/auth/domain/account_lockouts"

	"github.com/google/uuid"
)

type AccountLockoutCreateRequestDTO struct {
	UserID      string  `json:"user_id" validate:"required"`
	LockedUntil *string `json:"locked_until,omitempty"`
	LockReason  string  `json:"lock_reason" validate:"required"`
	LockedBy    *string `json:"locked_by,omitempty"`
	ActorID     *string `json:"actor_id,omitempty"`
}

type AccountLockoutUnlockRequestDTO struct {
	UserID        string  `json:"user_id" validate:"required"`
	UnlockedBy    string  `json:"unlocked_by" validate:"required"`
	UnlockActorID *string `json:"unlock_actor_id,omitempty"`
}

type AccountLockoutByUserIDRequestDTO struct {
	UserID string `uri:"user_id" validate:"required"`
}

type AccountLockoutDeleteRequestDTO struct {
	UserID string `uri:"user_id" validate:"required"`
}

func (r *AccountLockoutCreateRequestDTO) ToCreateCmd() (accountLockoutAppCommands.CreateAccountLockoutCmd, error) {
	uid, err := uuid.Parse(r.UserID)
	if err != nil {
		return accountLockoutAppCommands.CreateAccountLockoutCmd{}, err
	}

	cmd := accountLockoutAppCommands.CreateAccountLockoutCmd{
		UserID:      uid,
		LockedUntil: r.LockedUntil,
		LockReason:  accountLockoutDomain.LockReason(r.LockReason),
		LockedBy:    r.LockedBy,
	}

	if r.ActorID != nil {
		aid, err := uuid.Parse(*r.ActorID)
		if err == nil {
			cmd.ActorID = &aid
		}
	}

	return cmd, nil
}

func (r *AccountLockoutUnlockRequestDTO) ToUnlockCmd() (accountLockoutAppCommands.UnlockAccountCmd, error) {
	uid, err := uuid.Parse(r.UserID)
	if err != nil {
		return accountLockoutAppCommands.UnlockAccountCmd{}, err
	}

	cmd := accountLockoutAppCommands.UnlockAccountCmd{
		UserID:     uid,
		UnlockedBy: r.UnlockedBy,
	}

	if r.UnlockActorID != nil {
		aid, err := uuid.Parse(*r.UnlockActorID)
		if err == nil {
			cmd.UnlockActorID = &aid
		}
	}

	return cmd, nil
}

func (r *AccountLockoutDeleteRequestDTO) ToDeleteCmd() (accountLockoutAppCommands.DeleteAccountLockoutCmd, error) {
	uid, err := uuid.Parse(r.UserID)
	if err != nil {
		return accountLockoutAppCommands.DeleteAccountLockoutCmd{}, err
	}

	return accountLockoutAppCommands.DeleteAccountLockoutCmd{
		UserID: uid,
	}, nil
}
