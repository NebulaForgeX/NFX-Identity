package respdto

import (
	"time"

	memberAppResult "nfxid/modules/tenants/application/members/results"

	"github.com/google/uuid"
)

type MemberDTO struct {
	ID          uuid.UUID              `json:"id"`
	MemberID    uuid.UUID              `json:"member_id"`
	TenantID    uuid.UUID              `json:"tenant_id"`
	UserID      uuid.UUID              `json:"user_id"`
	Status      string                 `json:"status"`
	Source      string                 `json:"source"`
	JoinedAt    *time.Time             `json:"joined_at,omitempty"`
	LeftAt      *time.Time             `json:"left_at,omitempty"`
	CreatedAt   time.Time              `json:"created_at"`
	CreatedBy   *uuid.UUID             `json:"created_by,omitempty"`
	UpdatedAt   time.Time              `json:"updated_at"`
	ExternalRef *string                `json:"external_ref,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

// MemberROToDTO converts application MemberRO to response DTO
func MemberROToDTO(v *memberAppResult.MemberRO) *MemberDTO {
	if v == nil {
		return nil
	}

	return &MemberDTO{
		ID:          v.ID,
		MemberID:    v.MemberID,
		TenantID:    v.TenantID,
		UserID:      v.UserID,
		Status:      string(v.Status),
		Source:      string(v.Source),
		JoinedAt:    v.JoinedAt,
		LeftAt:      v.LeftAt,
		CreatedAt:   v.CreatedAt,
		CreatedBy:   v.CreatedBy,
		UpdatedAt:   v.UpdatedAt,
		ExternalRef: v.ExternalRef,
		Metadata:    v.Metadata,
	}
}

// MemberListROToDTO converts list of MemberRO to DTOs
func MemberListROToDTO(results []memberAppResult.MemberRO) []MemberDTO {
	dtos := make([]MemberDTO, len(results))
	for i, v := range results {
		if dto := MemberROToDTO(&v); dto != nil {
			dtos[i] = *dto
		}
	}
	return dtos
}
