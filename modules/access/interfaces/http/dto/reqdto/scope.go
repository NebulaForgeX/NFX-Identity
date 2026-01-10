package reqdto

import (
	scopeAppCommands "nfxid/modules/access/application/scopes/commands"
)

type ScopeCreateRequestDTO struct {
	Scope       string  `json:"scope" validate:"required"`
	Description *string `json:"description,omitempty"`
	IsSystem    bool    `json:"is_system,omitempty"`
}

type ScopeUpdateRequestDTO struct {
	Scope       string  `params:"scope" validate:"required"`
	Description *string `json:"description,omitempty"`
}

type ScopeByScopeRequestDTO struct {
	Scope string `params:"scope" validate:"required"`
}

func (r *ScopeCreateRequestDTO) ToCreateCmd() scopeAppCommands.CreateScopeCmd {
	return scopeAppCommands.CreateScopeCmd{
		Scope:       r.Scope,
		Description: r.Description,
		IsSystem:    r.IsSystem,
	}
}

func (r *ScopeUpdateRequestDTO) ToUpdateCmd() scopeAppCommands.UpdateScopeCmd {
	return scopeAppCommands.UpdateScopeCmd{
		Scope:       r.Scope,
		Description: r.Description,
	}
}
