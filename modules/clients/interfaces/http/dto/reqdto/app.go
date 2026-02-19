package reqdto

import (
	appAppCommands "nfxid/modules/clients/application/apps/commands"
	appDomain "nfxid/modules/clients/domain/apps"

	"github.com/google/uuid"
)

type AppCreateRequestDTO struct {
	AppID       string                 `json:"app_id" validate:"required"`
	TenantID    uuid.UUID              `json:"tenant_id" validate:"required"`
	Name        string                 `json:"name" validate:"required"`
	Description *string                `json:"description,omitempty"`
	Type        string                 `json:"type" validate:"required"`
	Status      string                 `json:"status,omitempty"`
	Environment string                 `json:"environment" validate:"required"`
	CreatedBy   *uuid.UUID             `json:"created_by,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type AppUpdateRequestDTO struct {
	ID          uuid.UUID              `uri:"id" validate:"required,uuid"`
	Name        string                 `json:"name" validate:"required"`
	Description *string                `json:"description,omitempty"`
	Type        string                 `json:"type" validate:"required"`
	Environment string                 `json:"environment" validate:"required"`
	UpdatedBy   *uuid.UUID             `json:"updated_by,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type AppByIDRequestDTO struct {
	ID uuid.UUID `uri:"id" validate:"required,uuid"`
}

type AppByAppIDRequestDTO struct {
	AppID string `uri:"app_id" validate:"required"`
}

func (r *AppCreateRequestDTO) ToCreateCmd() appAppCommands.CreateAppCmd {
	cmd := appAppCommands.CreateAppCmd{
		AppID:       r.AppID,
		TenantID:    r.TenantID,
		Name:        r.Name,
		Description: r.Description,
		CreatedBy:   r.CreatedBy,
		Metadata:    r.Metadata,
	}

	if r.Type != "" {
		cmd.Type = appDomain.AppType(r.Type)
	}
	if r.Status != "" {
		cmd.Status = appDomain.AppStatus(r.Status)
	}
	if r.Environment != "" {
		cmd.Environment = appDomain.Environment(r.Environment)
	}

	return cmd
}

func (r *AppUpdateRequestDTO) ToUpdateCmd() appAppCommands.UpdateAppCmd {
	cmd := appAppCommands.UpdateAppCmd{
		AppID:       r.ID,
		Name:        r.Name,
		Description: r.Description,
		UpdatedBy:   r.UpdatedBy,
		Metadata:    r.Metadata,
	}

	if r.Type != "" {
		cmd.Type = appDomain.AppType(r.Type)
	}
	if r.Environment != "" {
		cmd.Environment = appDomain.Environment(r.Environment)
	}

	return cmd
}
