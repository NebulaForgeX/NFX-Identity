package client_credentials

import (
	"time"

	"github.com/google/uuid"
)

func (cc *ClientCredential) UpdateLastUsed() error {
	if cc.RevokedAt() != nil {
		return ErrCredentialAlreadyRevoked
	}
	if cc.IsExpired() {
		return ErrCredentialExpired
	}

	now := time.Now().UTC()
	cc.state.LastUsedAt = &now
	return nil
}

func (cc *ClientCredential) Revoke(revokedBy uuid.UUID, reason string) error {
	if cc.RevokedAt() != nil {
		return ErrCredentialAlreadyRevoked
	}

	now := time.Now().UTC()
	cc.state.RevokedAt = &now
	cc.state.RevokedBy = &revokedBy
	cc.state.RevokeReason = &reason
	cc.state.Status = CredentialStatusRevoked
	return nil
}

func (cc *ClientCredential) Rotate(newSecretHash, newHashAlg string) error {
	if newSecretHash == "" {
		return ErrSecretHashRequired
	}
	if newHashAlg == "" {
		return ErrHashAlgRequired
	}

	now := time.Now().UTC()
	cc.state.RotatedAt = &now
	cc.state.SecretHash = newSecretHash
	cc.state.HashAlg = newHashAlg
	cc.state.Status = CredentialStatusRotating
	return nil
}

func (cc *ClientCredential) IsExpired() bool {
	if cc.ExpiresAt() == nil {
		return false
	}
	return time.Now().UTC().After(*cc.ExpiresAt())
}

func (cc *ClientCredential) IsRevoked() bool {
	return cc.RevokedAt() != nil
}

func (cc *ClientCredential) IsValid() bool {
	return cc.Status() == CredentialStatusActive && !cc.IsExpired() && !cc.IsRevoked()
}
