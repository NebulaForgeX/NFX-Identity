package reqdto

import (
	actionAppCommands "nfxid/modules/access/application/actions/commands"

	"github.com/google/uuid"
)

type ActionByIDRequestDTO struct {
	ID uuid.UUID `uri:"id" validate:"required,uuid"`
}

type ActionByKeyRequestDTO struct {
	Key string `uri:"key" validate:"required"`
}

type ActionCreateRequestDTO struct {
	Key         string  `json:"key" validate:"required"`
	Service     string  `json:"service" validate:"required"`
	Status      string  `json:"status" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	IsSystem    bool    `json:"is_system,omitempty"`
}

func (r *ActionCreateRequestDTO) ToCreateCmd() actionAppCommands.CreateActionCmd {
	return actionAppCommands.CreateActionCmd{
		Key:         r.Key,
		Service:     r.Service,
		Status:      r.Status,
		Name:        r.Name,
		Description: r.Description,
		IsSystem:    r.IsSystem,
	}
}
