package platform

import (
	"context"
	"encoding/json"
	"time"

	"nfxidentity/modules/auth/domain/authorityprofile"
	"nfxidentity/modules/auth/domain/forgerprofile"
	"nfxidentity/modules/auth/domain/settings"
	"nfxidentity/pkgs/errx"
	"nfxidentity/pkgs/transaction"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

func (s *Service) forgerVO(p *forgerprofile.Profile) map[string]any {
	return map[string]any{
		"profile_id": p.ID().String(), "account_id": p.AccountID().String(), "display_name": p.DisplayName(),
		"profile_language": p.ProfileLanguage(), "city": p.City(), "country": p.Country(), "website": p.Website(),
		"timezone": p.Timezone(), "forger_roles": p.Roles(), "created_at": p.CreatedAt(),
	}
}

func (s *Service) authorityVO(p *authorityprofile.Profile) map[string]any {
	return map[string]any{
		"profile_id": p.ID().String(), "account_id": p.AccountID().String(), "display_name": p.DisplayName(),
		"profile_language": p.ProfileLanguage(), "city": p.City(), "country": p.Country(),
		"authority_roles": p.Roles(), "created_at": p.CreatedAt(),
	}
}

func (s *Service) ForgerRoles(ctx context.Context, profileID uuid.UUID) ([]string, error) {
	p, err := s.repos.Forger(none()).Get.ByID(ctx, profileID)
	if err != nil {
		return nil, ErrNotFound
	}
	return p.Roles(), nil
}

func (s *Service) AuthorityRoles(ctx context.Context, profileID uuid.UUID) ([]string, error) {
	p, err := s.repos.Authority(none()).Get.ByID(ctx, profileID)
	if err != nil {
		return nil, ErrNotFound
	}
	return p.Roles(), nil
}

func (s *Service) HasAuthorityRole(ctx context.Context, profileID, role string) (bool, error) {
	pid, err := uuid.Parse(profileID)
	if err != nil {
		return false, err
	}
	p, err := s.repos.Authority(none()).Get.ByID(ctx, pid)
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
	p, err := s.repos.Forger(none()).Get.ByID(ctx, pid)
	if err != nil {
		return false, nil
	}
	return p.HasRole(role), nil
}

func (s *Service) PatchProfile(ctx context.Context, accountID, profileID uuid.UUID, kind string, patch map[string]any) error {
	switch kind {
	case "forger":
		p, err := s.repos.Forger(none()).Get.ByAccountAndID(ctx, accountID, profileID)
		if err != nil {
			return ErrProfileNotOwned
		}
		p.ApplyPatch(patch)
		return s.repos.Forger(none()).Update.Generic(ctx, p)
	case "authority":
		p, err := s.repos.Authority(none()).Get.ByAccountAndID(ctx, accountID, profileID)
		if err != nil {
			return ErrProfileNotOwned
		}
		p.ApplyPatch(patch)
		return s.repos.Authority(none()).Update.Generic(ctx, p)
	default:
		return errx.InvalidArg("INVALID_PROFILE_KIND", "kind must be forger or authority")
	}
}

func (s *Service) PatchProfileSettings(ctx context.Context, profileID uuid.UUID, kind string, loginNotification *bool) error {
	if loginNotification == nil {
		return nil
	}
	row, err := s.repos.Settings(none()).Get.ByID(ctx, kind, profileID)
	if err != nil {
		return err
	}
	row.SetLoginNotification(*loginNotification)
	return s.repos.Settings(none()).Update.Generic(ctx, row)
}

func (s *Service) UpdatePreference(ctx context.Context, accountID, profileID uuid.UUID, kind, preference string) error {
	raw := json.RawMessage([]byte(preference))
	if preference == "" {
		raw = json.RawMessage([]byte("{}"))
	}
	switch kind {
	case "forger":
		p, err := s.repos.Forger(none()).Get.ByAccountAndID(ctx, accountID, profileID)
		if err != nil {
			return ErrProfileNotOwned
		}
		p.SetPreference(raw)
		return s.repos.Forger(none()).Update.Generic(ctx, p)
	case "authority":
		p, err := s.repos.Authority(none()).Get.ByAccountAndID(ctx, accountID, profileID)
		if err != nil {
			return ErrProfileNotOwned
		}
		p.SetPreference(raw)
		return s.repos.Authority(none()).Update.Generic(ctx, p)
	default:
		return errx.InvalidArg("INVALID_PROFILE_KIND", "kind must be forger or authority")
	}
}

func (s *Service) CreateForgerProfile(ctx context.Context, accountID uuid.UUID, displayName, lang string) (string, error) {
	now := time.Now()
	id := uuid.New()
	err := s.tx.WithUoW(ctx, func(ctx context.Context, uow transaction.UoW) error {
		if err := s.repos.Forger(uow).Create.New(ctx, forgerprofile.NewFromState(forgerprofile.State{
			ID: id, AccountID: accountID, Roles: pq.StringArray{"forger"},
			ProfileLanguage: langOrDefault(lang), DisplayName: &displayName, CreatedAt: now, UpdatedAt: now,
		})); err != nil {
			return err
		}
		return s.repos.Settings(uow).Create.New(ctx, settings.NewFromState(settings.State{
			ID: id, Kind: "forger", LoginNotification: true, CreatedAt: now, UpdatedAt: now,
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
		p, err := s.repos.Forger(none()).Get.ByAccountAndID(ctx, accountID, profileID)
		if err != nil {
			return ErrProfileNotOwned
		}
		p.SoftDelete(now)
		return s.repos.Forger(none()).Update.Generic(ctx, p)
	case "authority":
		p, err := s.repos.Authority(none()).Get.ByAccountAndID(ctx, accountID, profileID)
		if err != nil {
			return ErrProfileNotOwned
		}
		p.SoftDelete(now)
		return s.repos.Authority(none()).Update.Generic(ctx, p)
	default:
		return errx.InvalidArg("INVALID_PROFILE_KIND", "kind must be forger or authority")
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
		bg := s.repos.Background(uow)
		imgRepo := s.repos.Image(uow)
		if err := bg.Delete.AllByProfileID(ctx, kind, profileID); err != nil {
			return err
		}
		for _, img := range images {
			iid, err := uuid.Parse(img.ImageID)
			if err != nil {
				return errx.InvalidArg("INVALID_IMAGE_ID", "invalid image id")
			}
			if _, err := imgRepo.Get.ByID(ctx, iid); err != nil {
				return err
			}
			if err := bg.Create.New(ctx, newBackground(profileID, iid, kind, img.SortOrder, now)); err != nil {
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
		return errx.InvalidArg("INVALID_IMAGE_ID", "invalid image id")
	}
	now := time.Now()
	return s.tx.WithUoW(ctx, func(ctx context.Context, uow transaction.UoW) error {
		if _, err := s.repos.Image(uow).Get.ByID(ctx, iid); err != nil {
			return err
		}
		av := s.repos.Avatar(uow)
		if err := av.Update.DeactivateActive(ctx, kind, profileID); err != nil {
			return err
		}
		return av.Create.New(ctx, newAvatar(profileID, iid, kind, now))
	})
}

func (s *Service) ClearAvatar(ctx context.Context, accountID, profileID uuid.UUID, kind string) error {
	if err := s.EnsureOwnedProfile(ctx, accountID.String(), profileID.String(), kind); err != nil {
		return err
	}
	return s.tx.WithUoW(ctx, func(ctx context.Context, uow transaction.UoW) error {
		return s.repos.Avatar(uow).Update.DeactivateActive(ctx, kind, profileID)
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
			out = append(out, map[string]any{
				"profile_id": r.ProfileID.String(), "account_id": r.AccountID.String(), "display_name": r.DisplayName,
				"profile_language": r.ProfileLanguage, "city": r.City, "country": r.Country,
				"authority_roles": []string(r.AuthorityRoles), "created_at": r.CreatedAt,
			})
		}
		return out, nil
	}
	rows, err := s.profiles.Forger.ByAccountID(ctx, accountID)
	if err != nil {
		return nil, err
	}
	out := make([]map[string]any, 0, len(rows))
	for _, r := range rows {
		out = append(out, map[string]any{
			"profile_id": r.ProfileID.String(), "account_id": r.AccountID.String(), "display_name": r.DisplayName,
			"profile_language": r.ProfileLanguage, "city": r.City, "country": r.Country, "website": r.Website,
			"timezone": r.Timezone, "forger_roles": []string(r.ForgerRoles), "created_at": r.CreatedAt,
		})
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
			items = append(items, map[string]any{
				"profile_id": r.ProfileID.String(), "account_id": r.AccountID.String(), "display_name": r.DisplayName,
				"profile_language": r.ProfileLanguage, "city": r.City, "country": r.Country,
				"authority_roles": []string(r.AuthorityRoles), "created_at": r.CreatedAt,
			})
		}
		return items, total, nil
	}
	rows, total, err := s.profiles.Forger.Search(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	items := make([]map[string]any, 0, len(rows))
	for _, r := range rows {
		items = append(items, map[string]any{
			"profile_id": r.ProfileID.String(), "account_id": r.AccountID.String(), "display_name": r.DisplayName,
			"profile_language": r.ProfileLanguage, "city": r.City, "country": r.Country, "website": r.Website,
			"timezone": r.Timezone, "forger_roles": []string(r.ForgerRoles), "created_at": r.CreatedAt,
		})
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
	if p, err := s.repos.Forger(none()).Get.ByID(ctx, profileID); err == nil {
		return s.forgerVO(p), nil
	}
	p, err := s.repos.Authority(none()).Get.ByID(ctx, profileID)
	if err != nil {
		return nil, ErrNotFound
	}
	return s.authorityVO(p), nil
}

func (s *Service) UpdateAuthorityRoles(ctx context.Context, actorAccountID, profileID uuid.UUID, roles []string) error {
	if err := s.requireOwner(ctx, actorAccountID); err != nil {
		return err
	}
	for _, r := range roles {
		if r == "owner" {
			return errx.Forbidden("OWNER_ROLE_IMMUTABLE", "owner role cannot be assigned via API")
		}
	}
	p, err := s.repos.Authority(none()).Get.ByID(ctx, profileID)
	if err != nil {
		return err
	}
	p.SetRoles(roles)
	return s.repos.Authority(none()).Update.Generic(ctx, p)
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
