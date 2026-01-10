package respdto

import (
	"time"

	groupAppResult "nfxid/modules/tenants/application/groups/results"

	"github.com/google/uuid"
)

type GroupDTO struct {
	ID            uuid.UUID              `json:"id"`
	GroupID       string                 `json:"group_id"`
	TenantID      uuid.UUID              `json:"tenant_id"`
	Name          string                 `json:"name"`
	Type          string                 `json:"type"`
	ParentGroupID *uuid.UUID             `json:"parent_group_id,omitempty"`
	Description   *string                `json:"description,omitempty"`
	CreatedBy     *uuid.UUID             `json:"created_by,omitempty"`
	Metadata      map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt     time.Time              `json:"created_at"`
	UpdatedAt     time.Time              `json:"updated_at"`
	DeletedAt     *time.Time             `json:"deleted_at,omitempty"`
}

// GroupROToDTO converts application GroupRO to response DTO
func GroupROToDTO(v *groupAppResult.GroupRO) *GroupDTO {
	if v == nil {
		return nil
	}

	return &GroupDTO{
		ID:            v.ID,
		GroupID:       v.GroupID,
		TenantID:      v.TenantID,
		Name:          v.Name,
		Type:          string(v.Type),
		ParentGroupID: v.ParentGroupID,
		Description:   v.Description,
		CreatedBy:     v.CreatedBy,
		Metadata:      v.Metadata,
		CreatedAt:     v.CreatedAt,
		UpdatedAt:     v.UpdatedAt,
		DeletedAt:     v.DeletedAt,
	}
}

// GroupListROToDTO converts list of GroupRO to DTOs
func GroupListROToDTO(results []groupAppResult.GroupRO) []GroupDTO {
	dtos := make([]GroupDTO, len(results))
	for i, v := range results {
		if dto := GroupROToDTO(&v); dto != nil {
			dtos[i] = *dto
		}
	}
	return dtos
}
