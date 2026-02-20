// Placeholder models for api_keys until schema/gen provides _dbgen.
package models

import (
	"nfxid/enums"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type ApiKey struct {
	ID             uuid.UUID                   `gorm:"type:uuid;primaryKey"`
	KeyID          string                      `gorm:"type:varchar(255)"`
	ApplicationID  uuid.UUID                   `gorm:"type:uuid"`
	KeyHash        string                      `gorm:"type:varchar(255)"`
	HashAlg        string                      `gorm:"type:varchar(50)"`
	Name           string                      `gorm:"type:varchar(255)"`
	Status         enums.ClientsApiKeyStatus    `gorm:"type:varchar(50)"`
	ExpiresAt      *time.Time                  `gorm:"type:timestamp"`
	CreatedAt      time.Time                   `gorm:"autoCreateTime"`
	RevokedAt      *time.Time                  `gorm:"type:timestamp"`
	RevokedBy      *uuid.UUID                  `gorm:"type:uuid"`
	RevokeReason   *string                     `gorm:"type:text"`
	LastUsedAt     *time.Time                  `gorm:"type:timestamp"`
	CreatedBy      *uuid.UUID                  `gorm:"type:uuid"`
	Metadata       *datatypes.JSON             `gorm:"type:jsonb"`
}

func (ApiKey) TableName() string { return "clients.api_keys" }

func (m *ApiKey) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == uuid.Nil {
		m.ID, err = uuid.NewV7()
	}
	return
}

var ApiKeyCols = struct {
	ID, KeyID, ApplicationID, KeyHash, HashAlg, Name, Status, ExpiresAt,
	CreatedAt, RevokedAt, RevokedBy, RevokeReason, LastUsedAt, CreatedBy, Metadata string
}{
	ID: "id", KeyID: "key_id", ApplicationID: "application_id", KeyHash: "key_hash", HashAlg: "hash_alg",
	Name: "name", Status: "status", ExpiresAt: "expires_at", CreatedAt: "created_at",
	RevokedAt: "revoked_at", RevokedBy: "revoked_by", RevokeReason: "revoke_reason",
	LastUsedAt: "last_used_at", CreatedBy: "created_by", Metadata: "metadata",
}
