// Placeholder models for member_app_roles until schema/gen provides _dbgen.
package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemberAppRole struct {
	ID            uuid.UUID  `gorm:"type:uuid;primaryKey"`
	MemberID      uuid.UUID  `gorm:"type:uuid"`
	ApplicationID uuid.UUID  `gorm:"type:uuid"`
	RoleID        uuid.UUID  `gorm:"type:uuid"`
	AssignedAt    time.Time  `gorm:"type:timestamp"`
	AssignedBy    *uuid.UUID `gorm:"type:uuid"`
	ExpiresAt     *time.Time `gorm:"type:timestamp"`
	RevokedAt     *time.Time `gorm:"type:timestamp"`
	RevokedBy     *uuid.UUID `gorm:"type:uuid"`
	RevokeReason  *string    `gorm:"type:text"`
}

func (MemberAppRole) TableName() string { return "tenants.member_app_roles" }

func (m *MemberAppRole) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == uuid.Nil {
		m.ID, err = uuid.NewV7()
	}
	return
}

var MemberAppRoleCols = struct {
	ID, MemberID, ApplicationID, RoleID, AssignedAt, AssignedBy, ExpiresAt, RevokedAt, RevokedBy, RevokeReason string
}{
	ID: "id", MemberID: "member_id", ApplicationID: "application_id", RoleID: "role_id",
	AssignedAt: "assigned_at", AssignedBy: "assigned_by", ExpiresAt: "expires_at",
	RevokedAt: "revoked_at", RevokedBy: "revoked_by", RevokeReason: "revoke_reason",
}
