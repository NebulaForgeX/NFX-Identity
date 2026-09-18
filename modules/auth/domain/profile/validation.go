package profile

import (
	"nfxidentity/constants"
	"nfxidentity/enums"
	authErr "nfxidentity/errors/src/auth"
)

// * =============================== Authority profile validation =============================== !//
func validateAuthorityProfileLanguage(lang enums.AuthProfileLanguage) error {
	switch lang {
	case enums.AuthProfileLanguageEn,
		enums.AuthProfileLanguageZh,
		enums.AuthProfileLanguageFr:
		return nil
	default:
		return authErr.ErrAuthorityProfileLanguageInvalid
	}
}

func validateAuthorityRole(role enums.AuthAuthorityRole) error {
	if constants.AuthAuthorityRole.Valid(role) {
		return nil
	}
	return authErr.ErrAuthorityProfileAuthorityRoleInvalid
}

func normalizeAuthorityRoles(roles []enums.AuthAuthorityRole) ([]enums.AuthAuthorityRole, error) {
	if len(roles) == 0 {
		return nil, authErr.ErrAuthorityProfileAuthorityRoleInvalid
	}
	seen := make(map[enums.AuthAuthorityRole]struct{}, len(roles))
	out := make([]enums.AuthAuthorityRole, 0, len(roles))
	for _, role := range roles {
		if err := validateAuthorityRole(role); err != nil {
			return nil, err
		}
		if _, ok := seen[role]; ok {
			continue
		}
		seen[role] = struct{}{}
		out = append(out, role)
	}
	if len(out) == 0 {
		return nil, authErr.ErrAuthorityProfileAuthorityRoleInvalid
	}
	return out, nil
}

func (e *AuthorityProfileEditable) Validate() error {
	return validateAuthorityProfileLanguage(e.ProfileLanguage)
}

// * =============================== Forger profile validation =============================== !//
func validateForgerProfileLanguage(lang enums.AuthProfileLanguage) error {
	switch lang {
	case enums.AuthProfileLanguageEn,
		enums.AuthProfileLanguageZh,
		enums.AuthProfileLanguageFr:
		return nil
	default:
		return authErr.ErrForgerProfileLanguageInvalid
	}
}

func validateForgerRole(role enums.AuthForgerRole) error {
	if constants.AuthForgerRole.Valid(role) {
		return nil
	}
	return authErr.ErrForgerProfileForgerRoleInvalid
}

func normalizeForgerRoles(roles []enums.AuthForgerRole) ([]enums.AuthForgerRole, error) {
	if len(roles) == 0 {
		return nil, authErr.ErrForgerProfileForgerRoleInvalid
	}
	seen := make(map[enums.AuthForgerRole]struct{}, len(roles))
	out := make([]enums.AuthForgerRole, 0, len(roles))
	for _, role := range roles {
		if err := validateForgerRole(role); err != nil {
			return nil, err
		}
		if _, ok := seen[role]; ok {
			continue
		}
		seen[role] = struct{}{}
		out = append(out, role)
	}
	if len(out) == 0 {
		return nil, authErr.ErrForgerProfileForgerRoleInvalid
	}
	return out, nil
}

func (e *ForgerProfileEditable) Validate() error {
	return validateForgerProfileLanguage(e.ProfileLanguage)
}
