package refreshtoken

import (
	"strings"
	"time"

	"nfxidentity/enums"
	authErr "nfxidentity/errors/src/auth"

	"github.com/google/uuid"
)

type NewRefreshTokenParams struct {
	AccountID    uuid.UUID
	IdentityID   *uuid.UUID
	ProfileID    *uuid.UUID
	ProfileScope *enums.AuthProfileScope
	DeviceID     *string
	TokenHash    string
	ExpiresAt    time.Time
}

func NewRefreshToken(p NewRefreshTokenParams) (*RefreshToken, error) {
	if p.AccountID == uuid.Nil {
		return nil, authErr.ErrRefreshTokenAccountIDInvalid
	}
	now := time.Now().UTC()
	exp := p.ExpiresAt.UTC()
	if err := validateRefreshTokenIssue(now, exp, p.TokenHash); err != nil {
		return nil, err
	}
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	return NewRefreshTokenFromState(RefreshTokenState{
		ID:           id,
		AccountID:    p.AccountID,
		IdentityID:   p.IdentityID,
		ProfileID:    p.ProfileID,
		ProfileScope: p.ProfileScope,
		DeviceID:     p.DeviceID,
		TokenHash:    strings.TrimSpace(p.TokenHash),
		ExpiresAt:    exp,
		CreatedAt:    now,
	}), nil
}
