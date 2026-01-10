package results

import (
	"time"

	"nfxid/modules/clients/domain/client_credentials"

	"github.com/google/uuid"
)

type ClientCredentialRO struct {
	ID           uuid.UUID
	AppID        uuid.UUID
	ClientID     string
	SecretHash   string
	HashAlg      string
	Status       client_credentials.CredentialStatus
	CreatedAt    time.Time
	RotatedAt    *time.Time
	ExpiresAt    *time.Time
	LastUsedAt   *time.Time
	CreatedBy    *uuid.UUID
	RevokedAt    *time.Time
	RevokedBy    *uuid.UUID
	RevokeReason *string
}

// ClientCredentialMapper 将 Domain ClientCredential 转换为 Application ClientCredentialRO
func ClientCredentialMapper(cc *client_credentials.ClientCredential) ClientCredentialRO {
	if cc == nil {
		return ClientCredentialRO{}
	}

	return ClientCredentialRO{
		ID:           cc.ID(),
		AppID:        cc.AppID(),
		ClientID:     cc.ClientID(),
		SecretHash:   cc.SecretHash(),
		HashAlg:      cc.HashAlg(),
		Status:       cc.Status(),
		CreatedAt:    cc.CreatedAt(),
		RotatedAt:    cc.RotatedAt(),
		ExpiresAt:    cc.ExpiresAt(),
		LastUsedAt:   cc.LastUsedAt(),
		CreatedBy:    cc.CreatedBy(),
		RevokedAt:    cc.RevokedAt(),
		RevokedBy:    cc.RevokedBy(),
		RevokeReason: cc.RevokeReason(),
	}
}
