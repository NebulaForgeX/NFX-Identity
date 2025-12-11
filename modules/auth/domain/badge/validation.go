package badge

import badgeErrors "nfxid/modules/auth/domain/badge/errors"

func (e *BadgeEditable) Validate() error {
	if e.Name == "" {
		return badgeErrors.ErrInvalidName
	}
	if len(e.Name) > 50 {
		return badgeErrors.ErrInvalidName
	}
	if e.Description != nil && *e.Description == "" {
		return badgeErrors.ErrInvalidDescription
	}
	if e.IconURL != nil && *e.IconURL == "" {
		return badgeErrors.ErrInvalidIconURL
	}
	if e.Color != nil && *e.Color == "" {
		return badgeErrors.ErrInvalidColor
	}
	if e.Category != nil {
		validCategories := map[string]struct{}{
			"achievement": {},
			"skill":       {},
			"community":   {},
			"special":     {},
		}
		if _, ok := validCategories[*e.Category]; !ok {
			return badgeErrors.ErrInvalidCategory
		}
	}
	return nil
}
