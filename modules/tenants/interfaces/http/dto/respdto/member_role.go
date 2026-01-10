package respdto

import (
	"time"

	memberRoleAppResult "nfxid/modules/tenants/application/member_roles/results"

	"github.com/google/uuid"
)

type MemberRoleDTO struct {
	ID          uuid.UUID  `json:"id"`
	TenantID    uuid.UUID  `json:"tenant_id"`
	MemberID    uuid.UUID  `json:"member_id"`
	RoleID      uuid.UUID  `json:"role_id"`
	AssignedAt  time.Time  `json:"assigned_at"`
	AssignedBy  *uuid.UUID `json:"assigned_by,omitempty"`
	ExpiresAt   *time.Time `json:"expires_at,omitempty"`
	Scope       *string    `json:"scope,omitempty"`
	RevokedAt   *time.Time `json:"revoked_at,omitempty"`
	RevokedBy   *uuid.UUID `json:"revoked_by,omitempty"`
	RevokeReason *string    `json:"revoke_reason,omitempty"`
}

// MemberRoleROToDTO converts application MemberRoleRO to response DTO
func MemberRoleROToDTO(v *memberRoleAppResult.MemberRoleRO) *MemberRoleDTO {
	if v == nil {
		return nil
	}

	return &MemberRoleDTO{
		ID:          v.ID,
		TenantID:    v.TenantID,
		MemberID:    v.MemberID,
		RoleID:      v.RoleID,
		AssignedAt:  v.AssignedAt,
		AssignedBy:  v.AssignedBy,
		ExpiresAt:   v.ExpiresAt,
		Scope:       v.Scope,
		RevokedAt:   v.RevokedAt,
		RevokedBy:   v.RevokedBy,
		RevokeReason: v.RevokeReason,
	}
}

// MemberRoleListROToDTO converts list of MemberRoleRO to DTOs
func MemberRoleListROToDTO(results []memberRoleAppResult.MemberRoleRO) []MemberRoleDTO {
	dtos := make([]MemberRoleDTO, len(results))
	for i, v := range results {
		if dto := MemberRoleROToDTO(&v); dto != nil {
			dtos[i] = *dto
		}
	}
	return dtos
}
