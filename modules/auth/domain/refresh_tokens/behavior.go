package refresh_tokens

import (
	authErr "nfxid/errors/src/auth"
	"time"
)

func (rt *RefreshToken) Revoke(reason RevokeReason) error {
	if rt.RevokedAt() != nil {
		return authErr.ErrTokenAlreadyRevoked
	}
	validReasons := map[RevokeReason]struct{}{
		RevokeReasonUserLogout:         {},
		RevokeReasonAdminRevoke:        {},
		RevokeReasonPasswordChanged:    {},
		RevokeReasonRotation:           {},
		RevokeReasonAccountLocked:      {},
		RevokeReasonDeviceChanged:      {},
		RevokeReasonSuspiciousActivity: {},
		RevokeReasonOther:              {},
	}
	if _, ok := validReasons[reason]; !ok {
		return authErr.ErrInvalidRevokeReason
	}

	now := time.Now().UTC()
	rt.state.RevokedAt = &now
	rt.state.RevokeReason = &reason
	rt.state.UpdatedAt = now
	return nil
}

func (rt *RefreshToken) IsExpired() bool {
	return time.Now().UTC().After(rt.ExpiresAt())
}

func (rt *RefreshToken) IsRevoked() bool {
	return rt.RevokedAt() != nil
}

func (rt *RefreshToken) IsValid() bool {
	return !rt.IsExpired() && !rt.IsRevoked()
}
