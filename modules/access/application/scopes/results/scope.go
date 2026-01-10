package results

import (
	"time"

	"nfxid/modules/access/domain/scopes"
)

type ScopeRO struct {
	Scope       string
	Description *string
	IsSystem    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

// ScopeMapper 将 Domain Scope 转换为 Application ScopeRO
func ScopeMapper(s *scopes.Scope) ScopeRO {
	if s == nil {
		return ScopeRO{}
	}

	return ScopeRO{
		Scope:       s.ScopeKey(),
		Description: s.Description(),
		IsSystem:    s.IsSystem(),
		CreatedAt:   s.CreatedAt(),
		UpdatedAt:   s.UpdatedAt(),
		DeletedAt:   s.DeletedAt(),
	}
}
