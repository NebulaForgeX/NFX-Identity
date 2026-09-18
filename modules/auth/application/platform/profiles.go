package platform

import (
	"context"
	"encoding/json"
	"time"

	"nfxidentity/enums"
	"nfxidentity/errors/src/auth"
	"nfxidentity/errors/src/sys"
	"nfxidentity/modules/auth/domain/profile"
	profileQuery "nfxidentity/modules/auth/query/profile"
	"nfxidentity/pkgs/patch"
	"nfxidentity/pkgs/query"
	"nfxidentity/pkgs/transaction"

	"github.com/google/uuid"
)

func (s *Service) forgerVO(p *profile.ForgerProfile) map[string]any {
	return map[string]any{
		"profile_id": p.ID().String(), "account_id": p.AccountID().String(), "display_name": p.DisplayName(),
		"profile_language": p.ProfileLanguage(), "preference": p.Preference(),
		"first_name": p.FirstName(), "last_name": p.LastName(), "city": p.City(), "country": p.Country(),
		"gender": p.Gender(), "birthday": p.Birthday(), "website": p.Website(), "timezone": p.Timezone(),
		"bio": p.Bio(), "forger_roles": p.ForgerRoles(), "created_at": p.CreatedAt(), "updated_at": p.UpdatedAt(),
	}
}

func (s *Service) authorityVO(p *profile.AuthorityProfile) map[string]any {
	return map[string]any{
		"profile_id": p.ID().String(), "account_id": p.AccountID().String(), "display_name": p.DisplayName(),
		"profile_language": p.ProfileLanguage(), "preference": p.Preference(),
		"first_name": p.FirstName(), "last_name": p.LastName(), "city": p.City(), "country": p.Country(),
		"gender": p.Gender(), "birthday": p.Birthday(), "website": p.Website(), "timezone": p.Timezone(),
		"bio": p.Bio(), "authority_roles": p.AuthorityRoles(), "created_at": p.CreatedAt(), "updated_at": p.UpdatedAt(),
	}
}

func avatarID(id *uuid.UUID) any {
	if id == nil {
		return nil
	}
	return id.String()
}

func forgerItemMap(r profileQuery.ForgerProfileItemVO) map[string]any {
	return map[string]any{
		"profile_id": r.ProfileID.String(), "account_id": r.AccountID.String(), "display_name": r.DisplayName,
		"profile_language": r.ProfileLanguage, "city": r.City, "country": r.Country, "website": r.Website,
		"timezone": r.Timezone, "forger_roles": roleStrings(r.ForgerRoles), "avatar_image_id": avatarID(r.AvatarImageID),
		"created_at": r.CreatedAt,
	}
}

func authorityItemMap(r profileQuery.AuthorityProfileItemVO) map[string]any {
	return map[string]any{
		"profile_id": r.ProfileID.String(), "account_id": r.AccountID.String(), "display_name": r.DisplayName,
		"profile_language": r.ProfileLanguage, "city": r.City, "country": r.Country,
		"authority_roles": roleStrings(r.AuthorityRoles), "avatar_image_id": avatarID(r.AvatarImageID),
		"created_at": r.CreatedAt,
	}
}

func (s *Service) ForgerRoles(ctx context.Context, profileID uuid.UUID) ([]string, error) {
	p, err := s.repoFactory.Profile(none()).Get.ForgerByProfileID(ctx, profileID)
	if err != nil {
		return nil, sys.ErrNotFound
	}
	return roleStrings(p.ForgerRoles()), nil
}

func (s *Service) AuthorityRoles(ctx context.Context, profileID uuid.UUID) ([]string, error) {
	p, err := s.repoFactory.Profile(none()).Get.AuthorityByProfileID(ctx, profileID)
	if err != nil {
		return nil, sys.ErrNotFound
	}
	return roleStrings(p.AuthorityRoles()), nil
}

func (s *Service) HasAuthorityRole(ctx context.Context, profileID, role string) (bool, error) {
	pid, err := uuid.Parse(profileID)
	if err != nil {
		return false, err
	}
	p, err := s.repoFactory.Profile(none()).Get.AuthorityByProfileID(ctx, pid)
	if err != nil {
		return false, nil
	}
	return p.HasRole(enums.AuthAuthorityRole(role)), nil
}

func (s *Service) HasForgerRole(ctx context.Context, profileID, role string) (bool, error) {
	pid, err := uuid.Parse(profileID)
	if err != nil {
		return false, err
	}
	p, err := s.repoFactory.Profile(none()).Get.ForgerByProfileID(ctx, pid)
	if err != nil {
		return false, nil
	}
	return p.HasRole(enums.AuthForgerRole(role)), nil
}

func (s *Service) PatchProfile(ctx context.Context, accountID, profileID uuid.UUID, kind string, raw map[string]any) error {
	switch kind {
	case "forger":
		p, err := s.repoFactory.Profile(none()).Get.ForgerByAccountAndID(ctx, accountID, profileID)
		if err != nil {
			return auth.ErrProfileNotOwned
		}
		if err := p.ApplyPatch(forgerPatchFromMap(raw)); err != nil {
			return err
		}
		return s.repoFactory.Profile(none()).Update.ForgerGeneric(ctx, p)
	case "authority":
		p, err := s.repoFactory.Profile(none()).Get.AuthorityByAccountAndID(ctx, accountID, profileID)
		if err != nil {
			return auth.ErrProfileNotOwned
		}
		if err := p.ApplyPatch(authorityPatchFromMap(raw)); err != nil {
			return err
		}
		return s.repoFactory.Profile(none()).Update.AuthorityGeneric(ctx, p)
	default:
		return auth.ErrInvalidProfileKind
	}
}

func (s *Service) PatchProfileSettings(ctx context.Context, profileID uuid.UUID, kind string, loginNotification *bool) error {
	if loginNotification == nil {
		return nil
	}
	switch kind {
	case "forger":
		return s.repoFactory.Profile(none()).Update.ForgerPartialSettings(ctx, profileID, profile.ForgerProfileSettingsPatch{
			LoginNotification: patch.Set(*loginNotification),
		})
	case "authority":
		return s.repoFactory.Profile(none()).Update.AuthorityPartialSettings(ctx, profileID, profile.AuthorityProfileSettingsPatch{
			LoginNotification: patch.Set(*loginNotification),
		})
	default:
		return auth.ErrInvalidProfileKind
	}
}

func (s *Service) UpdatePreference(ctx context.Context, accountID, profileID uuid.UUID, kind, preference string) error {
	raw := json.RawMessage([]byte(preference))
	if preference == "" {
		raw = json.RawMessage([]byte("{}"))
	}
	switch kind {
	case "forger":
		p, err := s.repoFactory.Profile(none()).Get.ForgerByAccountAndID(ctx, accountID, profileID)
		if err != nil {
			return auth.ErrProfileNotOwned
		}
		p.SetPreference(raw)
		return s.repoFactory.Profile(none()).Update.ForgerGeneric(ctx, p)
	case "authority":
		p, err := s.repoFactory.Profile(none()).Get.AuthorityByAccountAndID(ctx, accountID, profileID)
		if err != nil {
			return auth.ErrProfileNotOwned
		}
		p.SetPreference(raw)
		return s.repoFactory.Profile(none()).Update.AuthorityGeneric(ctx, p)
	default:
		return auth.ErrInvalidProfileKind
	}
}

func (s *Service) CreateForgerProfile(ctx context.Context, accountID uuid.UUID, displayName, lang string) (string, error) {
	now := time.Now()
	id := uuid.New()
	err := s.tx.WithUoW(ctx, func(ctx context.Context, uow transaction.UoW) error {
		profileRepo := s.repoFactory.Profile(uow)

		lang := enums.AuthProfileLanguage(langOrDefault(lang))
		if err := profileRepo.Create.NewForger(ctx, profile.NewForgerProfileFromState(profile.ForgerProfileState{
			ID: id, AccountID: accountID, ForgerRoles: []enums.AuthForgerRole{enums.AuthForgerRoleForger},
			ProfileLanguage: lang, Preference: profile.Default(lang), DisplayName: &displayName, CreatedAt: now, UpdatedAt: now,
		})); err != nil {
			return err
		}
		return profileRepo.Create.NewForgerSettings(ctx, profile.NewForgerProfileSettingsFromState(profile.ForgerProfileSettingsState{
			ID: id, LoginNotification: true, CreatedAt: now, UpdatedAt: now,
		}))
	})
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func (s *Service) CreateAuthorityProfile(ctx context.Context, accountID uuid.UUID, displayName, lang string) (string, error) {
	if err := s.requireOwner(ctx, accountID); err != nil {
		return "", err
	}
	now := time.Now()
	id := uuid.New()
	err := s.tx.WithUoW(ctx, func(ctx context.Context, uow transaction.UoW) error {
		profileRepo := s.repoFactory.Profile(uow)

		lang := enums.AuthProfileLanguage(langOrDefault(lang))
		if err := profileRepo.Create.NewAuthority(ctx, profile.NewAuthorityProfileFromState(profile.AuthorityProfileState{
			ID: id, AccountID: accountID, AuthorityRoles: []enums.AuthAuthorityRole{enums.AuthAuthorityRoleAdministrator},
			ProfileLanguage: lang, Preference: profile.Default(lang), DisplayName: &displayName, CreatedAt: now, UpdatedAt: now,
		})); err != nil {
			return err
		}
		return profileRepo.Create.NewAuthoritySettings(ctx, profile.NewAuthorityProfileSettingsFromState(profile.AuthorityProfileSettingsState{
			ID: id, LoginNotification: true, CreatedAt: now, UpdatedAt: now,
		}))
	})
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func (s *Service) DeleteProfile(ctx context.Context, accountID, profileID uuid.UUID, kind string) error {
	switch kind {
	case "forger":
		p, err := s.repoFactory.Profile(none()).Get.ForgerByAccountAndID(ctx, accountID, profileID)
		if err != nil {
			return auth.ErrProfileNotOwned
		}
		if err := p.Delete(); err != nil {
			return err
		}
		return s.repoFactory.Profile(none()).Update.ForgerGeneric(ctx, p)
	case "authority":
		p, err := s.repoFactory.Profile(none()).Get.AuthorityByAccountAndID(ctx, accountID, profileID)
		if err != nil {
			return auth.ErrProfileNotOwned
		}
		if err := p.Delete(); err != nil {
			return err
		}
		return s.repoFactory.Profile(none()).Update.AuthorityGeneric(ctx, p)
	default:
		return auth.ErrInvalidProfileKind
	}
}

func (s *Service) ConfirmBackgrounds(ctx context.Context, accountID, profileID uuid.UUID, kind string, images []struct {
	ImageID   string
	SortOrder int
}) error {
	if err := s.EnsureOwnedProfile(ctx, accountID.String(), profileID.String(), kind); err != nil {
		return err
	}
	now := time.Now()
	return s.tx.WithUoW(ctx, func(ctx context.Context, uow transaction.UoW) error {
		profileRepo := s.repoFactory.Profile(uow)
		imageRepo := s.repoFactory.Image(uow)
		switch kind {
		case "forger":
			if err := profileRepo.Delete.DeleteForgerBackgroundsByProfileID(ctx, profileID); err != nil {
				return err
			}
		case "authority":
			if err := profileRepo.Delete.DeleteAuthorityBackgroundsByProfileID(ctx, profileID); err != nil {
				return err
			}
		default:
			return auth.ErrInvalidProfileKind
		}
		for _, img := range images {
			iid, err := uuid.Parse(img.ImageID)
			if err != nil {
				return auth.ErrInvalidImageID
			}
			if _, err := imageRepo.Get.ByID(ctx, iid); err != nil {
				return err
			}
			switch kind {
			case "forger":
				if err := profileRepo.Create.NewForgerBackground(ctx, newForgerBackground(profileID, iid, img.SortOrder, now)); err != nil {
					return err
				}
			case "authority":
				if err := profileRepo.Create.NewAuthorityBackground(ctx, newAuthorityBackground(profileID, iid, img.SortOrder, now)); err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func (s *Service) ConfirmAvatar(ctx context.Context, accountID, profileID uuid.UUID, kind, imageID string) error {
	if err := s.EnsureOwnedProfile(ctx, accountID.String(), profileID.String(), kind); err != nil {
		return err
	}
	iid, err := uuid.Parse(imageID)
	if err != nil {
		return auth.ErrInvalidImageID
	}
	now := time.Now()
	inactive := false
	return s.tx.WithUoW(ctx, func(ctx context.Context, uow transaction.UoW) error {
		imageRepo := s.repoFactory.Image(uow)
		if _, err := imageRepo.Get.ByID(ctx, iid); err != nil {
			return err
		}
		profileRepo := s.repoFactory.Profile(uow)
		switch kind {
		case "forger":
			avatars, err := profileRepo.Get.ListForgerAvatarsByProfileID(ctx, profileID)
			if err != nil {
				return err
			}
			for _, a := range avatars {
				if !a.IsActive() {
					continue
				}
				if err := a.Update(profile.ForgerProfileAvatarEditable{IsActive: &inactive}); err != nil {
					return err
				}
				if err := profileRepo.Update.ForgerAvatarGeneric(ctx, a); err != nil {
					return err
				}
			}
			return profileRepo.Create.NewForgerAvatar(ctx, newForgerAvatar(profileID, iid, now))
		case "authority":
			avatars, err := profileRepo.Get.ListAuthorityAvatarsByProfileID(ctx, profileID)
			if err != nil {
				return err
			}
			for _, a := range avatars {
				if !a.IsActive() {
					continue
				}
				if err := a.Update(profile.AuthorityProfileAvatarEditable{IsActive: &inactive}); err != nil {
					return err
				}
				if err := profileRepo.Update.AuthorityAvatarGeneric(ctx, a); err != nil {
					return err
				}
			}
			return profileRepo.Create.NewAuthorityAvatar(ctx, newAuthorityAvatar(profileID, iid, now))
		default:
			return auth.ErrInvalidProfileKind
		}
	})
}

func (s *Service) ClearAvatar(ctx context.Context, accountID, profileID uuid.UUID, kind string) error {
	if err := s.EnsureOwnedProfile(ctx, accountID.String(), profileID.String(), kind); err != nil {
		return err
	}
	return s.tx.WithUoW(ctx, func(ctx context.Context, uow transaction.UoW) error {
		profileRepo := s.repoFactory.Profile(uow)
		switch kind {
		case "forger":
			return profileRepo.Delete.DeleteForgerAvatarsByProfileID(ctx, profileID)
		case "authority":
			return profileRepo.Delete.DeleteAuthorityAvatarsByProfileID(ctx, profileID)
		default:
			return auth.ErrInvalidProfileKind
		}
	})
}

func (s *Service) ListAccountProfiles(ctx context.Context, accountID uuid.UUID, kind string) ([]map[string]any, error) {
	if kind == "authority" {
		rows, err := s.profiles.AuthorityList.ByAccountID(ctx, accountID)
		if err != nil {
			return nil, err
		}
		out := make([]map[string]any, 0, len(rows))
		for _, r := range rows {
			out = append(out, authorityItemMap(r))
		}
		return out, nil
	}
	rows, err := s.profiles.ForgerList.ByAccountID(ctx, accountID)
	if err != nil {
		return nil, err
	}
	out := make([]map[string]any, 0, len(rows))
	for _, r := range rows {
		out = append(out, forgerItemMap(r))
	}
	return out, nil
}

func (s *Service) SearchProfiles(ctx context.Context, kind, qstr string, limit, offset int) ([]map[string]any, int64, error) {
	q := profileQuery.ListQuery{DomainPagination: query.DomainPagination{Limit: limitOr(limit, 20), Offset: offset}}
	if qstr != "" {
		q.Search = &qstr
	}
	if kind == "authority" {
		page, err := s.profiles.AuthorityList.Search(ctx, q)
		if err != nil {
			return nil, 0, err
		}
		items := make([]map[string]any, 0, len(page.Items))
		for _, r := range page.Items {
			items = append(items, authorityItemMap(r))
		}
		return items, page.Total, nil
	}
	page, err := s.profiles.ForgerList.Search(ctx, q)
	if err != nil {
		return nil, 0, err
	}
	items := make([]map[string]any, 0, len(page.Items))
	for _, r := range page.Items {
		items = append(items, forgerItemMap(r))
	}
	return items, page.Total, nil
}

func (s *Service) ListOwnerForgerProfiles(ctx context.Context, accountID uuid.UUID, query string, limit, offset int) ([]map[string]any, int64, error) {
	if err := s.requireOwner(ctx, accountID); err != nil {
		return nil, 0, err
	}
	return s.SearchProfiles(ctx, "forger", query, limit, offset)
}

func (s *Service) ListOwnerAuthorityProfiles(ctx context.Context, accountID uuid.UUID, query string, limit, offset int) ([]map[string]any, int64, error) {
	if err := s.requireOwner(ctx, accountID); err != nil {
		return nil, 0, err
	}
	return s.SearchProfiles(ctx, "authority", query, limit, offset)
}

func (s *Service) PublicProfileCard(ctx context.Context, profileID uuid.UUID) (map[string]any, error) {
	if p, err := s.repoFactory.Profile(none()).Get.ForgerByProfileID(ctx, profileID); err == nil {
		return s.forgerVO(p), nil
	}
	p, err := s.repoFactory.Profile(none()).Get.AuthorityByProfileID(ctx, profileID)
	if err != nil {
		return nil, sys.ErrNotFound
	}
	return s.authorityVO(p), nil
}

func (s *Service) UpdateAuthorityRoles(ctx context.Context, actorAccountID, profileID uuid.UUID, roles []string) error {
	if err := s.requireOwner(ctx, actorAccountID); err != nil {
		return err
	}
	converted := make([]enums.AuthAuthorityRole, 0, len(roles))
	for _, r := range roles {
		if r == "owner" {
			return auth.ErrOwnerRoleImmutable
		}
		converted = append(converted, enums.AuthAuthorityRole(r))
	}
	p, err := s.repoFactory.Profile(none()).Get.AuthorityByProfileID(ctx, profileID)
	if err != nil {
		return err
	}
	if err := p.SetAuthorityRoles(converted); err != nil {
		return err
	}
	return s.repoFactory.Profile(none()).Update.AuthorityGeneric(ctx, p)
}

func (s *Service) BatchPublicCards(ctx context.Context, ids []string) []map[string]any {
	out := make([]map[string]any, 0, len(ids))
	for _, id := range ids {
		pid, err := uuid.Parse(id)
		if err != nil {
			continue
		}
		if card, err := s.PublicProfileCard(ctx, pid); err == nil {
			out = append(out, card)
		}
	}
	return out
}

func forgerPatchFromMap(raw map[string]any) profile.ForgerProfilePatch {
	p := profile.ForgerProfilePatch{
		DisplayName: strPatch(raw, "display_name"),
		FirstName:   strPatch(raw, "first_name"),
		LastName:    strPatch(raw, "last_name"),
		Country:     strPatch(raw, "country"),
		City:        strPatch(raw, "city"),
		Gender:      strPatch(raw, "gender"),
		Website:     strPatch(raw, "website"),
		Timezone:    strPatch(raw, "timezone"),
		Bio:         strPatch(raw, "bio"),
		Birthday:    timePatch(raw, "birthday"),
	}
	if v, ok := raw["profile_language"].(string); ok {
		p.ProfileLanguage = patch.Set(enums.AuthProfileLanguage(v))
	}
	return p
}

func authorityPatchFromMap(raw map[string]any) profile.AuthorityProfilePatch {
	p := profile.AuthorityProfilePatch{
		DisplayName: strPatch(raw, "display_name"),
		FirstName:   strPatch(raw, "first_name"),
		LastName:    strPatch(raw, "last_name"),
		Country:     strPatch(raw, "country"),
		City:        strPatch(raw, "city"),
		Gender:      strPatch(raw, "gender"),
		Website:     strPatch(raw, "website"),
		Timezone:    strPatch(raw, "timezone"),
		Bio:         strPatch(raw, "bio"),
		Birthday:    timePatch(raw, "birthday"),
	}
	if v, ok := raw["profile_language"].(string); ok {
		p.ProfileLanguage = patch.Set(enums.AuthProfileLanguage(v))
	}
	return p
}

func strPatch(raw map[string]any, key string) patch.PatchField[string] {
	v, ok := raw[key]
	if !ok {
		return patch.Unset[string]()
	}
	if v == nil {
		return patch.SetNull[string]()
	}
	s, ok := v.(string)
	if !ok {
		return patch.Unset[string]()
	}
	return patch.Set(s)
}

func timePatch(raw map[string]any, key string) patch.PatchField[time.Time] {
	v, ok := raw[key]
	if !ok {
		return patch.Unset[time.Time]()
	}
	if v == nil {
		return patch.SetNull[time.Time]()
	}
	s, ok := v.(string)
	if !ok {
		return patch.Unset[time.Time]()
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return patch.Unset[time.Time]()
	}
	return patch.Set(t)
}
