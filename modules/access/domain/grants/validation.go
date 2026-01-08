package grants

import "github.com/google/uuid"

func (g *Grant) Validate() error {
	if g.SubjectType() == "" {
		return ErrSubjectTypeRequired
	}
	validSubjectTypes := map[SubjectType]struct{}{
		SubjectTypeUser:   {},
		SubjectTypeClient: {},
	}
	if _, ok := validSubjectTypes[g.SubjectType()]; !ok {
		return ErrInvalidSubjectType
	}
	if g.SubjectID() == uuid.Nil {
		return ErrSubjectIDRequired
	}
	if g.GrantType() == "" {
		return ErrGrantTypeRequired
	}
	validGrantTypes := map[GrantType]struct{}{
		GrantTypeRole:       {},
		GrantTypePermission: {},
	}
	if _, ok := validGrantTypes[g.GrantType()]; !ok {
		return ErrInvalidGrantType
	}
	if g.GrantRefID() == uuid.Nil {
		return ErrGrantRefIDRequired
	}
	validEffects := map[GrantEffect]struct{}{
		GrantEffectAllow: {},
		GrantEffectDeny:  {},
	}
	if _, ok := validEffects[g.Effect()]; !ok {
		return ErrInvalidGrantEffect
	}
	return nil
}
