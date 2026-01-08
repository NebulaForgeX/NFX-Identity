package scopes

import (
	"time"
)

type NewScopeParams struct {
	Scope       string
	Description *string
	IsSystem    bool
}

func NewScope(p NewScopeParams) (*Scope, error) {
	if err := validateScopeParams(p); err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewScopeFromState(ScopeState{
		Scope:       p.Scope,
		Description: p.Description,
		IsSystem:    p.IsSystem,
		CreatedAt:   now,
		UpdatedAt:   now,
	}), nil
}

func NewScopeFromState(st ScopeState) *Scope {
	return &Scope{state: st}
}

func validateScopeParams(p NewScopeParams) error {
	if p.Scope == "" {
		return ErrScopeRequired
	}
	return nil
}
