package profile_badge

import profileBadgeErrors "nfxid/modules/auth/domain/profile_badge/errors"

func (e *ProfileBadgeEditable) Validate() error {
	if e.Level != nil && *e.Level < 0 {
		return profileBadgeErrors.ErrProfileBadgeNotFound
	}
	return nil
}
