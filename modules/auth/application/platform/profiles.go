package platform

import (
	"context"
	"encoding/json"
	"nfxidentity/errors/src/auth"
	"nfxidentity/errors/src/sys"
	"time"

	"nfxidentity/modules/auth/domain/authorityprofile"
	"nfxidentity/modules/auth/domain/forgerprofile"
	"nfxidentity/modules/auth/domain/settings"
	profileQuery "nfxidentity/modules/auth/query/profile"
	"nfxidentity/pkgs/transaction"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

func (s *Service) forgerVO(p *forgerprofile.Profile) map[string]any {
	return map[string]any{
		"profile_id": p.ID().String(), "account_id": p.AccountID().String(), "display_name": p.DisplayName(),
		"profile_language": p.ProfileLanguage(), "preference": p.Preference(),
		"first_name": p.FirstName(), "last_name": p.LastName(), "city": p.City(), "country": p.Country(),
		"gender": p.Gender(), "birthday": p.Birthday(), "website": p.Website(), "timezone": p.Timezone(),
		"bio": p.Bio(), "forger_roles": p.Roles(), "created_at": p.CreatedAt(), "updated_at": p.UpdatedAt(),
	}
}

func (s *Service) authorityVO(p *authorityprofile.Profile) map[string]any {
	return map[string]any{
		"profile_id": p.ID().String(), "account_id": p.AccountID().String(), "display_name": p.DisplayName(),
		"profile_language": p.ProfileLanguage(), "preference": p.Preference(),
		"first_name": p.FirstName(), "last_name": p.LastName(), "city": p.City(), "country": p.Country(),
		"gender": p.Gender(), "birthday": p.Birthday(), "website": p.Website(), "timezone": p.Timezone(),
		"bio": p.Bio(), "authority_roles": p.Roles(), "created_at": p.CreatedAt(), "updated_at": p.UpdatedAt(),
	}
}

func avatarID(id *uuid.UUID) any {
	if id == nil {
		return nil
	}
	return id.String()
}

func forgerItemMap(r profileQuery.ForgerItemVO) map[string]any {
	return map[string]any{
		"profile_id": r.ProfileID.String(), "account_id": r.AccountID.String(), "display_name": r.DisplayName,
		"profile_language": r.ProfileLanguage, "city": r.City, "country": r.Country, "website": r.Website,
		"timezone": r.Timezone, "forger_roles": []string(r.ForgerRoles), "avatar_image_id": avatarID(r.AvatarImageID),
		"created_at": r.CreatedAt,
	}
}

func authorityItemMap(r profileQuery.AuthorityItemVO) map[string]any {
	return map[string]any{
		"profile_id": r.ProfileID.String(), "account_id": r.AccountID.String(), "display_name": r.DisplayName,
		"profile_language": r.ProfileLanguage, "city": r.City, "country": r.Country,
		"authority_roles": []string(r.AuthorityRoles), "avatar_image_id": avatarID(r.AvatarImageID),
		"created_at": r.CreatedAt,
	}
}

func (s *Service) ForgerRoles(ctx context.Context, profileID uuid.UUID) ([]string, error) {
	p, err := s.repoFactory.Forger(none()).Get.ByID(ctx, profileID)
	if err != nil {
		return nil, sys.ErrNotFound
	}
	return p.Roles(), nil
}

func (s *Service) AuthorityRoles(ctx context.Context, profileID uuid.UUID) ([]string, error) {
	p, err := s.repoFactory.Authority(none()).Get.ByID(ctx, profileID)
	if err != nil {
		return nil, sys.ErrNotFound
	}
	return p.Roles(), nil
}

func (s *Service) HasAuthorityRole(ctx context.Context, profileID, role string) (bool, error) {
	pid, err := uuid.Parse(profileID)
	if err != nil {
		return false, err
	}
	p, err := s.repoFactory.Authority(none()).Get.ByID(ctx, pid)
	if err != nil {
		return false, nil
	}
	return p.HasRole(role), nil
}

func (s *Service) HasForgerRole(ctx context.Context, profileID, role string) (bool, error) {
	pid, err := uuid.Parse(profileID)
	if err != nil {
		return false, err
	}
	p, err := s.repoFactory.Forger(none()).Get.ByID(ctx, pid)
	if err != nil {
		return false, nil
	}
	return p.HasRole(role), nil
}

func (s *Service) PatchProfile(ctx context.Context, accountID, profileID uuid.UUID, kind string, patch map[string]any) error {
	switch kind {
	case "forger":
		p, err := s.repoFactory.Forger(none()).Get.ByAccountAndID(ctx, accountID, profileID)
		if err != nil {
			return auth.ErrProfileNotOwned
		}
		p.ApplyPatch(patch)
		return s.repoFactory.Forger(none()).Update.Generic(ctx, p)
	case "authority":
		p, err := s.repoFactory.Authority(none()).Get.ByAccountAndID(ctx, accountID, profileID)
		if err != nil {
			return auth.ErrProfileNotOwned
		}
		p.ApplyPatch(patch)
		return s.repoFactory.Authority(none()).Update.Generic(ctx, p)
	default:
		return auth.ErrInvalidProfileKind
	}
}

func (s *Service) PatchProfileSettings(ctx context.Context, profileID uuid.UUID, kind string, loginNotification *bool) error {
	if loginNotification == nil {
		return nil
	}
	row, err := s.repoFactory.Settings(none()).Get.ByID(ctx, kind, profileID)
	if err != nil {
		return err
	}
	row.SetLoginNotification(*loginNotification)
	return s.repoFactory.Settings(none()).Update.Generic(ctx, row)
}

func (s *Service) UpdatePreference(ctx context.Context, accountID, profileID uuid.UUID, kind, preference string) error {
	raw := json.RawMessage([]byte(preference))
	if preference == "" {
		raw = json.RawMessage([]byte("{}"))
	}
	switch kind {
	case "forger":
		p, err := s.repoFactory.Forger(none()).Get.ByAccountAndID(ctx, accountID, profileID)
		if err != nil {
			return auth.ErrProfileNotOwned
		}
		p.SetPreference(raw)
		return s.repoFactory.Forger(none()).Update.Generic(ctx, p)
	case "authority":
		p, err := s.repoFactory.Authority(none()).Get.ByAccountAndID(ctx, accountID, profileID)
		if err != nil {
			return auth.ErrProfileNotOwned
		}
		p.SetPreference(raw)
		return s.repoFactory.Authority(none()).Update.Generic(ctx, p)
	default:
		return auth.ErrInvalidProfileKind
	}
}

func (s *Service) CreateForgerProfile(ctx context.Context, accountID uuid.UUID, displayName, lang string) (string, error) {
	now := time.Now()
	id := uuid.New()
	err := s.tx.WithUoW(ctx, func(ctx context.Context, uow transaction.UoW) error {
		forgerRepo := s.repoFactory.Forger(uow)
		settingsRepo := s.repoFactory.Settings(uow)
		if err := forgerRepo.Create.New(ctx, forgerprofile.NewFromState(forgerprofile.State{
			ID: id, AccountID: accountID, Roles: pq.StringArray{"forger"},
			ProfileLanguage: langOrDefault(lang), DisplayName: &displayName, CreatedAt: now, UpdatedAt: now,
		})); err != nil {
			return err
		}
		return settingsRepo.Create.New(ctx, settings.NewFromState(settings.State{
			ID: id, Kind: "forger", LoginNotification: true, CreatedAt: now, UpdatedAt: now,
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
		authorityRepo := s.repoFactory.Authority(uow)
		settingsRepo := s.repoFactory.Settings(uow)
		if err := authorityRepo.Create.New(ctx, authorityprofile.NewFromState(authorityprofile.State{
			ID: id, AccountID: accountID, Roles: pq.StringArray{"administrator"},
			ProfileLanguage: langOrDefault(lang), DisplayName: &displayName, CreatedAt: now, UpdatedAt: now,
		})); err != nil {
			return err
		}
		return settingsRepo.Create.New(ctx, settings.NewFromState(settings.State{
			ID: id, Kind: "authority", LoginNotification: true, CreatedAt: now, UpdatedAt: now,
		}))
	})
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func (s *Service) DeleteProfile(ctx context.Context, accountID, profileID uuid.UUID, kind string) error {
	now := time.Now()
	switch kind {
	case "forger":
		p, err := s.repoFactory.Forger(none()).Get.ByAccountAndID(ctx, accountID, profileID)
		if err != nil {
			return auth.ErrProfileNotOwned
		}
		p.SoftDelete(now)
		return s.repoFactory.Forger(none()).Update.Generic(ctx, p)
	case "authority":
		p, err := s.repoFactory.Authority(none()).Get.ByAccountAndID(ctx, accountID, profileID)
		if err != nil {
			return auth.ErrProfileNotOwned
		}
		p.SoftDelete(now)
		return s.repoFactory.Authority(none()).Update.Generic(ctx, p)
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
		backgroundRepo := s.repoFactory.Background(uow)
		imageRepo := s.repoFactory.Image(uow)
		if err := backgroundRepo.Delete.AllByProfileID(ctx, kind, profileID); err != nil {
			return err
		}
		for _, img := range images {
			iid, err := uuid.Parse(img.ImageID)
			if err != nil {
				return auth.ErrInvalidImageID
			}
			if _, err := imageRepo.Get.ByID(ctx, iid); err != nil {
				return err
			}
			if err := backgroundRepo.Create.New(ctx, newBackground(profileID, iid, kind, img.SortOrder, now)); err != nil {
				return err
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
	return s.tx.WithUoW(ctx, func(ctx context.Context, uow transaction.UoW) error {
		imageRepo := s.repoFactory.Image(uow)
		if _, err := imageRepo.Get.ByID(ctx, iid); err != nil {
			return err
		}
		avatarRepo := s.repoFactory.Avatar(uow)
		if err := avatarRepo.Update.DeactivateActive(ctx, kind, profileID); err != nil {
			return err
		}
		return avatarRepo.Create.New(ctx, newAvatar(profileID, iid, kind, now))
	})
}

func (s *Service) ClearAvatar(ctx context.Context, accountID, profileID uuid.UUID, kind string) error {
	if err := s.EnsureOwnedProfile(ctx, accountID.String(), profileID.String(), kind); err != nil {
		return err
	}
	return s.tx.WithUoW(ctx, func(ctx context.Context, uow transaction.UoW) error {
		return s.repoFactory.Avatar(uow).Update.DeactivateActive(ctx, kind, profileID)
	})
}

func (s *Service) ListAccountProfiles(ctx context.Context, accountID uuid.UUID, kind string) ([]map[string]any, error) {
	if kind == "authority" {
		rows, err := s.profiles.Authority.ByAccountID(ctx, accountID)
		if err != nil {
			return nil, err
		}
		out := make([]map[string]any, 0, len(rows))
		for _, r := range rows {
			out = append(out, authorityItemMap(r))
		}
		return out, nil
	}
	rows, err := s.profiles.Forger.ByAccountID(ctx, accountID)
	if err != nil {
		return nil, err
	}
	out := make([]map[string]any, 0, len(rows))
	for _, r := range rows {
		out = append(out, forgerItemMap(r))
	}
	return out, nil
}

func (s *Service) SearchProfiles(ctx context.Context, kind, query string, limit, offset int) ([]map[string]any, int64, error) {
	limit = limitOr(limit, 20)
	if kind == "authority" {
		rows, total, err := s.profiles.Authority.Search(ctx, query, limit, offset)
		if err != nil {
			return nil, 0, err
		}
		items := make([]map[string]any, 0, len(rows))
		for _, r := range rows {
			items = append(items, authorityItemMap(r))
		}
		return items, total, nil
	}
	rows, total, err := s.profiles.Forger.Search(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	items := make([]map[string]any, 0, len(rows))
	for _, r := range rows {
		items = append(items, forgerItemMap(r))
	}
	return items, total, nil
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
	if p, err := s.repoFactory.Forger(none()).Get.ByID(ctx, profileID); err == nil {
		return s.forgerVO(p), nil
	}
	p, err := s.repoFactory.Authority(none()).Get.ByID(ctx, profileID)
	if err != nil {
		return nil, sys.ErrNotFound
	}
	return s.authorityVO(p), nil
}

func (s *Service) UpdateAuthorityRoles(ctx context.Context, actorAccountID, profileID uuid.UUID, roles []string) error {
	if err := s.requireOwner(ctx, actorAccountID); err != nil {
		return err
	}
	for _, r := range roles {
		if r == "owner" {
			return auth.ErrOwnerRoleImmutable
		}
	}
	p, err := s.repoFactory.Authority(none()).Get.ByID(ctx, profileID)
	if err != nil {
		return err
	}
	p.SetRoles(roles)
	return s.repoFactory.Authority(none()).Update.Generic(ctx, p)
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
