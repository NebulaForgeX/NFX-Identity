package scopes

import (
	"time"
)

type Scope struct {
	state ScopeState
}

type ScopeState struct {
	Scope       string
	Description *string
	IsSystem    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

func (s *Scope) ScopeKey() string      { return s.state.Scope }
func (s *Scope) Description() *string  { return s.state.Description }
func (s *Scope) IsSystem() bool       { return s.state.IsSystem }
func (s *Scope) CreatedAt() time.Time { return s.state.CreatedAt }
func (s *Scope) UpdatedAt() time.Time { return s.state.UpdatedAt }
func (s *Scope) DeletedAt() *time.Time { return s.state.DeletedAt }
