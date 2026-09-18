package profile

import (
	"encoding/json"
	"time"

	"nfxidentity/enums"

	"gorm.io/datatypes"
)

func (u *AuthorityProfile) HasRole(role enums.AuthAuthorityRole) bool {
	for _, existing := range u.state.AuthorityRoles {
		if existing == role {
			return true
		}
	}
	return false
}

func (u *ForgerProfile) HasRole(role enums.AuthForgerRole) bool {
	for _, existing := range u.state.ForgerRoles {
		if existing == role {
			return true
		}
	}
	return false
}

func (u *AuthorityProfile) SetPreference(raw json.RawMessage) {
	if len(raw) == 0 {
		u.state.Preference = nil
		return
	}
	j := datatypes.JSON(raw)
	u.state.Preference = &j
	u.state.UpdatedAt = time.Now().UTC()
}

func (u *ForgerProfile) SetPreference(raw json.RawMessage) {
	if len(raw) == 0 {
		u.state.Preference = nil
		return
	}
	j := datatypes.JSON(raw)
	u.state.Preference = &j
	u.state.UpdatedAt = time.Now().UTC()
}
