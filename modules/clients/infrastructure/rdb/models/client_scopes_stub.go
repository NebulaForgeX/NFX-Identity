// Placeholder models for client_scopes until schema/gen provides _dbgen. DO NOT EDIT by hand when gen is added.
package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClientScope struct {
	ID           uuid.UUID  `gorm:"type:uuid;primaryKey"`
	ApplicationID uuid.UUID  `gorm:"type:uuid"`
	Scope        string     `gorm:"type:varchar(255)"`
	GrantedBy    *uuid.UUID `gorm:"type:uuid"`
	GrantedAt    time.Time  `gorm:"type:timestamp"`
	ExpiresAt    *time.Time `gorm:"type:timestamp"`
	CreatedAt    time.Time  `gorm:"autoCreateTime"`
	RevokedAt    *time.Time `gorm:"type:timestamp"`
	RevokedBy    *uuid.UUID `gorm:"type:uuid"`
	RevokeReason *string    `gorm:"type:text"`
}

func (ClientScope) TableName() string { return "clients.client_scopes" }

func (m *ClientScope) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == uuid.Nil {
		m.ID, err = uuid.NewV7()
	}
	return
}

var ClientScopeCols = struct {
	ID, ApplicationID, Scope, GrantedBy, GrantedAt, ExpiresAt, CreatedAt, RevokedAt, RevokedBy, RevokeReason string
}{
	ID: "id", ApplicationID: "application_id", Scope: "scope", GrantedBy: "granted_by", GrantedAt: "granted_at",
	ExpiresAt: "expires_at", CreatedAt: "created_at", RevokedAt: "revoked_at", RevokedBy: "revoked_by", RevokeReason: "revoke_reason",
}
