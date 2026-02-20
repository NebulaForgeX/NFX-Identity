package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemberRole struct {
	ID           uuid.UUID  `gorm:"type:uuid;primaryKey"`
	TenantID     uuid.UUID  `gorm:"type:uuid"`
	MemberID     uuid.UUID  `gorm:"type:uuid"`
	RoleID       uuid.UUID  `gorm:"type:uuid"`
	AssignedAt   time.Time  `gorm:"type:timestamp"`
	AssignedBy   *uuid.UUID `gorm:"type:uuid"`
	ExpiresAt    *time.Time `gorm:"type:timestamp"`
	Scope        *string    `gorm:"type:text"`
	RevokedAt    *time.Time `gorm:"type:timestamp"`
	RevokedBy    *uuid.UUID `gorm:"type:uuid"`
	RevokeReason *string    `gorm:"type:text"`
}

func (MemberRole) TableName() string { return "tenants.member_roles" }

func (m *MemberRole) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == uuid.Nil {
		m.ID, err = uuid.NewV7()
	}
	return
}

var MemberRoleCols = struct {
	ID, TenantID, MemberID, RoleID, AssignedAt, AssignedBy, ExpiresAt, Scope, RevokedAt, RevokedBy, RevokeReason string
}{
	ID: "id", TenantID: "tenant_id", MemberID: "member_id", RoleID: "role_id", AssignedAt: "assigned_at",
	AssignedBy: "assigned_by", ExpiresAt: "expires_at", Scope: "scope", RevokedAt: "revoked_at",
	RevokedBy: "revoked_by", RevokeReason: "revoke_reason",
}
