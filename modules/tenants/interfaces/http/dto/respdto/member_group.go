package respdto

import (
	"time"

	memberGroupAppResult "nfxid/modules/tenants/application/member_groups/results"

	"github.com/google/uuid"
)

type MemberGroupDTO struct {
	ID         uuid.UUID  `json:"id"`
	MemberID   uuid.UUID  `json:"member_id"`
	GroupID    uuid.UUID  `json:"group_id"`
	AssignedAt time.Time  `json:"assigned_at"`
	AssignedBy *uuid.UUID `json:"assigned_by,omitempty"`
	RevokedAt  *time.Time `json:"revoked_at,omitempty"`
	RevokedBy  *uuid.UUID `json:"revoked_by,omitempty"`
}

// MemberGroupROToDTO converts application MemberGroupRO to response DTO
func MemberGroupROToDTO(v *memberGroupAppResult.MemberGroupRO) *MemberGroupDTO {
	if v == nil {
		return nil
	}

	return &MemberGroupDTO{
		ID:         v.ID,
		MemberID:   v.MemberID,
		GroupID:    v.GroupID,
		AssignedAt: v.AssignedAt,
		AssignedBy: v.AssignedBy,
		RevokedAt:  v.RevokedAt,
		RevokedBy:  v.RevokedBy,
	}
}

// MemberGroupListROToDTO converts list of MemberGroupRO to DTOs
func MemberGroupListROToDTO(results []memberGroupAppResult.MemberGroupRO) []MemberGroupDTO {
	dtos := make([]MemberGroupDTO, len(results))
	for i, v := range results {
		if dto := MemberGroupROToDTO(&v); dto != nil {
			dtos[i] = *dto
		}
	}
	return dtos
}
