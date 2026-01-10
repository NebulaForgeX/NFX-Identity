package respdto

import (
	"time"

	grantAppResult "nfxid/modules/access/application/grants/results"

	"github.com/google/uuid"
)

type GrantDTO struct {
	ID           uuid.UUID  `json:"id"`
	SubjectType  string     `json:"subject_type"`
	SubjectID    uuid.UUID  `json:"subject_id"`
	GrantType    string     `json:"grant_type"`
	GrantRefID   uuid.UUID  `json:"grant_ref_id"`
	TenantID     *uuid.UUID `json:"tenant_id,omitempty"`
	AppID        *uuid.UUID `json:"app_id,omitempty"`
	ResourceType *string    `json:"resource_type,omitempty"`
	ResourceID   *uuid.UUID `json:"resource_id,omitempty"`
	Effect       string     `json:"effect"`
	ExpiresAt    *time.Time `json:"expires_at,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	CreatedBy    *uuid.UUID `json:"created_by,omitempty"`
	RevokedAt    *time.Time `json:"revoked_at,omitempty"`
	RevokedBy    *uuid.UUID `json:"revoked_by,omitempty"`
	RevokeReason *string    `json:"revoke_reason,omitempty"`
}

// GrantROToDTO converts application GrantRO to response DTO
func GrantROToDTO(v *grantAppResult.GrantRO) *GrantDTO {
	if v == nil {
		return nil
	}

	return &GrantDTO{
		ID:           v.ID,
		SubjectType:  string(v.SubjectType),
		SubjectID:    v.SubjectID,
		GrantType:    string(v.GrantType),
		GrantRefID:   v.GrantRefID,
		TenantID:     v.TenantID,
		AppID:        v.AppID,
		ResourceType: v.ResourceType,
		ResourceID:   v.ResourceID,
		Effect:       string(v.Effect),
		ExpiresAt:    v.ExpiresAt,
		CreatedAt:    v.CreatedAt,
		CreatedBy:    v.CreatedBy,
		RevokedAt:    v.RevokedAt,
		RevokedBy:    v.RevokedBy,
		RevokeReason: v.RevokeReason,
	}
}
