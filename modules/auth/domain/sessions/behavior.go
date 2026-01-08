package sessions

import (
	"time"
)

func (s *Session) UpdateLastSeen() error {
	if s.RevokedAt() != nil {
		return ErrSessionAlreadyRevoked
	}
	if s.IsExpired() {
		return ErrSessionExpired
	}

	now := time.Now().UTC()
	s.state.LastSeenAt = now
	s.state.UpdatedAt = now
	return nil
}

func (s *Session) Revoke(reason SessionRevokeReason, revokedBy string) error {
	if s.RevokedAt() != nil {
		return ErrSessionAlreadyRevoked
	}
	validReasons := map[SessionRevokeReason]struct{}{
		SessionRevokeReasonUserLogout:      {},
		SessionRevokeReasonAdminRevoke:     {},
		SessionRevokeReasonPasswordChanged: {},
		SessionRevokeReasonDeviceChanged:   {},
		SessionRevokeReasonAccountLocked:   {},
		SessionRevokeReasonSuspiciousActivity: {},
		SessionRevokeReasonSessionExpired:  {},
		SessionRevokeReasonOther:           {},
	}
	if _, ok := validReasons[reason]; !ok {
		return ErrInvalidRevokeReason
	}

	now := time.Now().UTC()
	s.state.RevokedAt = &now
	s.state.RevokeReason = &reason
	s.state.RevokedBy = &revokedBy
	s.state.UpdatedAt = now
	return nil
}

func (s *Session) IsExpired() bool {
	return time.Now().UTC().After(s.ExpiresAt())
}

func (s *Session) IsRevoked() bool {
	return s.RevokedAt() != nil
}

func (s *Session) IsValid() bool {
	return !s.IsExpired() && !s.IsRevoked()
}
