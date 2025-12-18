package authorization_code

import (
	"time"

	"github.com/google/uuid"
)

type NewAuthorizationCodeParams struct {
	Editable  AuthorizationCodeEditable
	CreatedBy *uuid.UUID
	ExpiresAt *time.Time
	IsActive  bool
}

func NewAuthorizationCode(p NewAuthorizationCodeParams) (*AuthorizationCode, error) {
	if err := p.Editable.Validate(); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewAuthorizationCodeFromState(AuthorizationCodeState{
		ID:        id,
		Editable:  p.Editable,
		CreatedBy: p.CreatedBy,
		ExpiresAt: p.ExpiresAt,
		IsActive:  p.IsActive,
		CreatedAt: now,
		UpdatedAt: now,
	}), nil
}

func NewAuthorizationCodeFromState(st AuthorizationCodeState) *AuthorizationCode {
	return &AuthorizationCode{state: st}
}
