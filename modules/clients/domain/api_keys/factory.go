package api_keys

import (
	clientsErr "nfxid/errors/src/clients"
	"time"

	"github.com/google/uuid"
)

type NewAPIKeyParams struct {
	KeyID     string
	AppID     uuid.UUID
	KeyHash   string
	HashAlg   string
	Name      string
	Status    APIKeyStatus
	ExpiresAt *time.Time
	CreatedBy *uuid.UUID
	Metadata  map[string]interface{}
}

func NewAPIKey(p NewAPIKeyParams) (*APIKey, error) {
	if err := validateAPIKeyParams(p); err != nil {
		return nil, err
	}

	status := p.Status
	if status == "" {
		status = APIKeyStatusActive
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewAPIKeyFromState(APIKeyState{
		ID:        id,
		KeyID:     p.KeyID,
		AppID:     p.AppID,
		KeyHash:   p.KeyHash,
		HashAlg:   p.HashAlg,
		Name:      p.Name,
		Status:    status,
		ExpiresAt: p.ExpiresAt,
		CreatedBy: p.CreatedBy,
		Metadata:  p.Metadata,
		CreatedAt: now,
	}), nil
}

func NewAPIKeyFromState(st APIKeyState) *APIKey {
	return &APIKey{state: st}
}

func validateAPIKeyParams(p NewAPIKeyParams) error {
	if p.KeyID == "" {
		return clientsErr.ErrKeyIDRequired
	}
	if p.AppID == uuid.Nil {
		return clientsErr.ErrAppIDRequired
	}
	if p.KeyHash == "" {
		return clientsErr.ErrKeyHashRequired
	}
	if p.HashAlg == "" {
		return clientsErr.ErrHashAlgRequired
	}
	if p.Name == "" {
		return clientsErr.ErrNameRequired
	}
	if p.Status != "" {
		validStatuses := map[APIKeyStatus]struct{}{
			APIKeyStatusActive:  {},
			APIKeyStatusRevoked: {},
			APIKeyStatusExpired: {},
		}
		if _, ok := validStatuses[p.Status]; !ok {
			return clientsErr.ErrInvalidAPIKeyStatus
		}
	}
	return nil
}
