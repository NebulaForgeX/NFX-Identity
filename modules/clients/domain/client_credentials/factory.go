package client_credentials

import (
	"time"

	"github.com/google/uuid"
)

type NewClientCredentialParams struct {
	AppID      uuid.UUID
	ClientID   string
	SecretHash string
	HashAlg    string
	Status     CredentialStatus
	ExpiresAt  *time.Time
	CreatedBy  *uuid.UUID
}

func NewClientCredential(p NewClientCredentialParams) (*ClientCredential, error) {
	if err := validateClientCredentialParams(p); err != nil {
		return nil, err
	}

	status := p.Status
	if status == "" {
		status = CredentialStatusActive
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewClientCredentialFromState(ClientCredentialState{
		ID:         id,
		AppID:      p.AppID,
		ClientID:   p.ClientID,
		SecretHash: p.SecretHash,
		HashAlg:    p.HashAlg,
		Status:     status,
		ExpiresAt:  p.ExpiresAt,
		CreatedBy:  p.CreatedBy,
		CreatedAt:  now,
	}), nil
}

func NewClientCredentialFromState(st ClientCredentialState) *ClientCredential {
	return &ClientCredential{state: st}
}

func validateClientCredentialParams(p NewClientCredentialParams) error {
	if p.ClientID == "" {
		return ErrClientIDRequired
	}
	if p.AppID == uuid.Nil {
		return ErrAppIDRequired
	}
	if p.SecretHash == "" {
		return ErrSecretHashRequired
	}
	if p.HashAlg == "" {
		return ErrHashAlgRequired
	}
	if p.Status != "" {
		validStatuses := map[CredentialStatus]struct{}{
			CredentialStatusActive:   {},
			CredentialStatusExpired:  {},
			CredentialStatusRevoked:  {},
			CredentialStatusRotating: {},
		}
		if _, ok := validStatuses[p.Status]; !ok {
			return ErrInvalidCredentialStatus
		}
	}
	return nil
}
