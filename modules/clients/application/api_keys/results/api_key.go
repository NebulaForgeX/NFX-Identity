package results

import (
	"time"

	"nfxid/modules/clients/domain/api_keys"

	"github.com/google/uuid"
)

type APIKeyRO struct {
	ID          uuid.UUID
	KeyID       string
	AppID       uuid.UUID
	KeyHash     string
	HashAlg     string
	Name        string
	Status      api_keys.APIKeyStatus
	ExpiresAt   *time.Time
	CreatedAt   time.Time
	RevokedAt   *time.Time
	RevokedBy   *uuid.UUID
	RevokeReason *string
	LastUsedAt  *time.Time
	CreatedBy   *uuid.UUID
	Metadata    map[string]interface{}
}

// APIKeyMapper 将 Domain APIKey 转换为 Application APIKeyRO
func APIKeyMapper(ak *api_keys.APIKey) APIKeyRO {
	if ak == nil {
		return APIKeyRO{}
	}

	return APIKeyRO{
		ID:          ak.ID(),
		KeyID:       ak.KeyID(),
		AppID:       ak.AppID(),
		KeyHash:     ak.KeyHash(),
		HashAlg:     ak.HashAlg(),
		Name:        ak.Name(),
		Status:      ak.Status(),
		ExpiresAt:   ak.ExpiresAt(),
		CreatedAt:   ak.CreatedAt(),
		RevokedAt:   ak.RevokedAt(),
		RevokedBy:   ak.RevokedBy(),
		RevokeReason: ak.RevokeReason(),
		LastUsedAt:  ak.LastUsedAt(),
		CreatedBy:   ak.CreatedBy(),
		Metadata:    ak.Metadata(),
	}
}
