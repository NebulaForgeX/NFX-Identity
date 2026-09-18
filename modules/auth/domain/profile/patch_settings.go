package profile

import (
	authErr "nfxidentity/errors/src/auth"
	"nfxidentity/pkgs/patch"
	"time"
)

// * =============================== Authority profile settings patch =============================== !//
type AuthorityProfileSettingsPatch struct {
	LoginNotification patch.PatchField[bool]
}

func (p AuthorityProfileSettingsPatch) Validate() error {
	if patch.IsPatchEmpty(p) {
		return authErr.ErrAuthorityProfileSettingsPatchEmpty
	}
	return nil
}

func (p AuthorityProfileSettingsPatch) DiffAgainst(s *AuthorityProfileSettings) AuthorityProfileSettingsPatch {
	return AuthorityProfileSettingsPatch{
		LoginNotification: p.LoginNotification.DiffVal(s.LoginNotification()),
	}
}

func (s *AuthorityProfileSettings) ApplyPatch(p AuthorityProfileSettingsPatch) error {
	if err := s.EnsureNotDeleted(); err != nil {
		return err
	}
	patch.ApplyVal(&s.state.LoginNotification, p.LoginNotification)
	s.state.UpdatedAt = time.Now().UTC()
	return nil
}

// * =============================== Forger profile settings patch =============================== !//
type ForgerProfileSettingsPatch struct {
	LoginNotification patch.PatchField[bool]
}

func (p ForgerProfileSettingsPatch) Validate() error {
	if patch.IsPatchEmpty(p) {
		return authErr.ErrForgerProfileSettingsPatchEmpty
	}
	return nil
}

func (p ForgerProfileSettingsPatch) DiffAgainst(s *ForgerProfileSettings) ForgerProfileSettingsPatch {
	return ForgerProfileSettingsPatch{
		LoginNotification: p.LoginNotification.DiffVal(s.LoginNotification()),
	}
}

func (s *ForgerProfileSettings) ApplyPatch(p ForgerProfileSettingsPatch) error {
	if err := s.EnsureNotDeleted(); err != nil {
		return err
	}
	patch.ApplyVal(&s.state.LoginNotification, p.LoginNotification)
	s.state.UpdatedAt = time.Now().UTC()
	return nil
}
