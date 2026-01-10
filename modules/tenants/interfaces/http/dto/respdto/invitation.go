package respdto

import (
	"time"

	invitationAppResult "nfxid/modules/tenants/application/invitations/results"

	"github.com/google/uuid"
)

type InvitationDTO struct {
	ID              uuid.UUID              `json:"id"`
	InviteID        string                 `json:"invite_id"`
	TenantID        uuid.UUID              `json:"tenant_id"`
	Email           string                 `json:"email"`
	TokenHash       string                 `json:"token_hash"`
	ExpiresAt       time.Time              `json:"expires_at"`
	Status          string                 `json:"status"`
	InvitedBy       uuid.UUID              `json:"invited_by"`
	InvitedAt       time.Time              `json:"invited_at"`
	AcceptedByUserID *uuid.UUID            `json:"accepted_by_user_id,omitempty"`
	AcceptedAt      *time.Time             `json:"accepted_at,omitempty"`
	RevokedBy       *uuid.UUID             `json:"revoked_by,omitempty"`
	RevokedAt       *time.Time             `json:"revoked_at,omitempty"`
	RevokeReason    *string                `json:"revoke_reason,omitempty"`
	RoleIDs         []uuid.UUID            `json:"role_ids,omitempty"`
	Metadata        map[string]interface{} `json:"metadata,omitempty"`
}

// InvitationROToDTO converts application InvitationRO to response DTO
func InvitationROToDTO(v *invitationAppResult.InvitationRO) *InvitationDTO {
	if v == nil {
		return nil
	}

	return &InvitationDTO{
		ID:              v.ID,
		InviteID:        v.InviteID,
		TenantID:        v.TenantID,
		Email:           v.Email,
		TokenHash:       v.TokenHash,
		ExpiresAt:       v.ExpiresAt,
		Status:          string(v.Status),
		InvitedBy:       v.InvitedBy,
		InvitedAt:       v.InvitedAt,
		AcceptedByUserID: v.AcceptedByUserID,
		AcceptedAt:      v.AcceptedAt,
		RevokedBy:       v.RevokedBy,
		RevokedAt:       v.RevokedAt,
		RevokeReason:    v.RevokeReason,
		RoleIDs:         v.RoleIDs,
		Metadata:        v.Metadata,
	}
}

// InvitationListROToDTO converts list of InvitationRO to DTOs
func InvitationListROToDTO(results []invitationAppResult.InvitationRO) []InvitationDTO {
	dtos := make([]InvitationDTO, len(results))
	for i, v := range results {
		if dto := InvitationROToDTO(&v); dto != nil {
			dtos[i] = *dto
		}
	}
	return dtos
}
