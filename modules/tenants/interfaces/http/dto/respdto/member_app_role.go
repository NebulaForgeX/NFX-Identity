package respdto

import (
	"time"

	memberAppRoleAppResult "nfxid/modules/tenants/application/member_app_roles/results"

	"github.com/google/uuid"
)

type MemberAppRoleDTO struct {
	ID           uuid.UUID  `json:"id"`
	MemberID     uuid.UUID  `json:"member_id"`
	AppID        uuid.UUID  `json:"app_id"`
	RoleID       uuid.UUID  `json:"role_id"`
	AssignedAt   time.Time  `json:"assigned_at"`
	AssignedBy   *uuid.UUID `json:"assigned_by,omitempty"`
	ExpiresAt    *time.Time `json:"expires_at,omitempty"`
	RevokedAt    *time.Time `json:"revoked_at,omitempty"`
	RevokedBy    *uuid.UUID `json:"revoked_by,omitempty"`
	RevokeReason *string    `json:"revoke_reason,omitempty"`
}

// MemberAppRoleROToDTO converts application MemberAppRoleRO to response DTO
func MemberAppRoleROToDTO(v *memberAppRoleAppResult.MemberAppRoleRO) *MemberAppRoleDTO {
	if v == nil {
		return nil
	}

	return &MemberAppRoleDTO{
		ID:           v.ID,
		MemberID:     v.MemberID,
		AppID:        v.AppID,
		RoleID:       v.RoleID,
		AssignedAt:   v.AssignedAt,
		AssignedBy:   v.AssignedBy,
		ExpiresAt:    v.ExpiresAt,
		RevokedAt:    v.RevokedAt,
		RevokedBy:    v.RevokedBy,
		RevokeReason: v.RevokeReason,
	}
}

// MemberAppRoleListROToDTO converts list of MemberAppRoleRO to DTOs
func MemberAppRoleListROToDTO(results []memberAppRoleAppResult.MemberAppRoleRO) []MemberAppRoleDTO {
	dtos := make([]MemberAppRoleDTO, len(results))
	for i, v := range results {
		if dto := MemberAppRoleROToDTO(&v); dto != nil {
			dtos[i] = *dto
		}
	}
	return dtos
}
