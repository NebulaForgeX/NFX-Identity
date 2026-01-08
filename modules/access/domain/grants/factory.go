package grants

import (
	"time"

	"github.com/google/uuid"
)

type NewGrantParams struct {
	SubjectType  SubjectType
	SubjectID    uuid.UUID
	GrantType    GrantType
	GrantRefID   uuid.UUID
	TenantID     *uuid.UUID
	AppID        *uuid.UUID
	ResourceType *string
	ResourceID   *uuid.UUID
	Effect       GrantEffect
	ExpiresAt    *time.Time
	CreatedBy    *uuid.UUID
}

func NewGrant(p NewGrantParams) (*Grant, error) {
	if err := validateGrantParams(p); err != nil {
		return nil, err
	}

	effect := p.Effect
	if effect == "" {
		effect = GrantEffectAllow
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewGrantFromState(GrantState{
		ID:           id,
		SubjectType:  p.SubjectType,
		SubjectID:    p.SubjectID,
		GrantType:    p.GrantType,
		GrantRefID:   p.GrantRefID,
		TenantID:     p.TenantID,
		AppID:        p.AppID,
		ResourceType: p.ResourceType,
		ResourceID:   p.ResourceID,
		Effect:       effect,
		ExpiresAt:    p.ExpiresAt,
		CreatedAt:    now,
		CreatedBy:    p.CreatedBy,
	}), nil
}

func NewGrantFromState(st GrantState) *Grant {
	return &Grant{state: st}
}

func validateGrantParams(p NewGrantParams) error {
	if p.SubjectType == "" {
		return ErrSubjectTypeRequired
	}
	validSubjectTypes := map[SubjectType]struct{}{
		SubjectTypeUser:   {},
		SubjectTypeClient: {},
	}
	if _, ok := validSubjectTypes[p.SubjectType]; !ok {
		return ErrInvalidSubjectType
	}
	if p.SubjectID == uuid.Nil {
		return ErrSubjectIDRequired
	}
	if p.GrantType == "" {
		return ErrGrantTypeRequired
	}
	validGrantTypes := map[GrantType]struct{}{
		GrantTypeRole:       {},
		GrantTypePermission: {},
	}
	if _, ok := validGrantTypes[p.GrantType]; !ok {
		return ErrInvalidGrantType
	}
	if p.GrantRefID == uuid.Nil {
		return ErrGrantRefIDRequired
	}
	if p.Effect != "" {
		validEffects := map[GrantEffect]struct{}{
			GrantEffectAllow: {},
			GrantEffectDeny:  {},
		}
		if _, ok := validEffects[p.Effect]; !ok {
			return ErrInvalidGrantEffect
		}
	}
	return nil
}
