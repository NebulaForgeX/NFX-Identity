package reqdto

import (
	groupAppCommands "nfxid/modules/tenants/application/groups/commands"
	groupDomain "nfxid/modules/tenants/domain/groups"

	"github.com/google/uuid"
)

type GroupCreateRequestDTO struct {
	GroupID       string                 `json:"group_id" validate:"required"`
	TenantID      uuid.UUID              `json:"tenant_id" validate:"required,uuid"`
	Name          string                 `json:"name" validate:"required"`
	Type          string                 `json:"type" validate:"required"`
	ParentGroupID *uuid.UUID             `json:"parent_group_id,omitempty"`
	Description   *string                `json:"description,omitempty"`
	CreatedBy     *uuid.UUID             `json:"created_by,omitempty"`
	Metadata      map[string]interface{} `json:"metadata,omitempty"`
}

type GroupUpdateRequestDTO struct {
	ID            uuid.UUID              `params:"id" validate:"required,uuid"`
	Name          string                 `json:"name" validate:"required"`
	Type          string                 `json:"type" validate:"required"`
	ParentGroupID *uuid.UUID             `json:"parent_group_id,omitempty"`
	Description   *string                `json:"description,omitempty"`
	Metadata      map[string]interface{} `json:"metadata,omitempty"`
}

type GroupByIDRequestDTO struct {
	ID uuid.UUID `params:"id" validate:"required,uuid"`
}

func (r *GroupCreateRequestDTO) ToCreateCmd() groupAppCommands.CreateGroupCmd {
	return groupAppCommands.CreateGroupCmd{
		GroupID:       r.GroupID,
		TenantID:      r.TenantID,
		Name:          r.Name,
		Type:          groupDomain.GroupType(r.Type),
		ParentGroupID: r.ParentGroupID,
		Description:   r.Description,
		CreatedBy:     r.CreatedBy,
		Metadata:      r.Metadata,
	}
}

func (r *GroupUpdateRequestDTO) ToUpdateCmd() groupAppCommands.UpdateGroupCmd {
	return groupAppCommands.UpdateGroupCmd{
		GroupID:       r.ID,
		Name:          r.Name,
		Type:          groupDomain.GroupType(r.Type),
		ParentGroupID: r.ParentGroupID,
		Description:   r.Description,
		Metadata:      r.Metadata,
	}
}
