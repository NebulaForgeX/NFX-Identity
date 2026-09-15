package platform

import (
	"context"
	"strings"
	"time"

	"nfxidentity/modules/auth/infrastructure/rdb"
	"nfxidentity/pkgs/errx"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func (s *Service) SendSignupCode(ctx context.Context, email string) error {
	email = strings.ToLower(strings.TrimSpace(email))
	if email == "" {
		return errx.InvalidArg("INVALID_EMAIL", "email required")
	}
	code := RandomCode()
	s.StoreVerificationCode(ctx, email, code)
	return nil
}

func (s *Service) ListEmails(ctx context.Context, accountID uuid.UUID) ([]map[string]any, int, error) {
	var rows []rdb.Email
	if err := s.db.WithContext(ctx).Where("account_id = ? AND deleted_at IS NULL", accountID).Order("created_at").Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return emailsJSON(rows), len(rows), nil
}

func (s *Service) CreateEmail(ctx context.Context, accountID uuid.UUID, email string) (string, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	if email == "" {
		return "", errx.InvalidArg("INVALID_EMAIL", "email required")
	}
	now := time.Now()
	id := uuid.New()
	row := rdb.Email{ID: id, AccountID: accountID, Email: email, CreatedAt: now, UpdatedAt: now}
	if err := s.db.WithContext(ctx).Create(&row).Error; err != nil {
		return "", errx.Conflict("EMAIL_TAKEN", "email already registered")
	}
	return id.String(), nil
}

func (s *Service) SendEmailVerificationCode(ctx context.Context, accountID, emailID uuid.UUID) error {
	var row rdb.Email
	if err := s.db.WithContext(ctx).Where("id = ? AND account_id = ? AND deleted_at IS NULL", emailID, accountID).First(&row).Error; err != nil {
		return ErrNotFound
	}
	s.StoreVerificationCode(ctx, row.Email, RandomCode())
	return nil
}

func (s *Service) VerifyEmail(ctx context.Context, accountID, emailID uuid.UUID, code string) error {
	var row rdb.Email
	if err := s.db.WithContext(ctx).Where("id = ? AND account_id = ? AND deleted_at IS NULL", emailID, accountID).First(&row).Error; err != nil {
		return ErrNotFound
	}
	if !s.checkVerificationCode(ctx, row.Email, code) {
		return errx.InvalidArg("INVALID_VERIFICATION_CODE", "invalid verification code")
	}
	now := time.Now()
	return s.db.WithContext(ctx).Model(&row).Updates(map[string]any{"verified_at": now, "updated_at": now}).Error
}

func (s *Service) UpdateEmail(ctx context.Context, accountID, emailID uuid.UUID, email string) error {
	email = strings.ToLower(strings.TrimSpace(email))
	res := s.db.WithContext(ctx).Model(&rdb.Email{}).
		Where("id = ? AND account_id = ? AND deleted_at IS NULL", emailID, accountID).
		Updates(map[string]any{"email": email, "verified_at": nil, "updated_at": time.Now()})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

func (s *Service) SetPrimaryEmail(ctx context.Context, accountID, emailID uuid.UUID) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&rdb.Email{}).Where("account_id = ? AND deleted_at IS NULL", accountID).Update("is_primary", false).Error; err != nil {
			return err
		}
		res := tx.Model(&rdb.Email{}).Where("id = ? AND account_id = ? AND deleted_at IS NULL", emailID, accountID).Update("is_primary", true)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return ErrNotFound
		}
		return nil
	})
}

func (s *Service) DeleteEmail(ctx context.Context, accountID, emailID uuid.UUID) error {
	now := time.Now()
	res := s.db.WithContext(ctx).Model(&rdb.Email{}).
		Where("id = ? AND account_id = ? AND deleted_at IS NULL AND is_primary = false", emailID, accountID).
		Update("deleted_at", now)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errx.FailedPrecond("EMAIL_NOT_DELETABLE", "cannot delete primary or missing email")
	}
	return nil
}

func (s *Service) ListPhones(ctx context.Context, accountID uuid.UUID) ([]map[string]any, int, error) {
	var rows []rdb.Phone
	if err := s.db.WithContext(ctx).Where("account_id = ? AND deleted_at IS NULL", accountID).Order("created_at").Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return phonesJSON(rows), len(rows), nil
}

func (s *Service) CreatePhone(ctx context.Context, accountID uuid.UUID, phone string) (string, error) {
	phone = strings.TrimSpace(phone)
	if phone == "" {
		return "", errx.InvalidArg("INVALID_PHONE", "phone required")
	}
	now := time.Now()
	id := uuid.New()
	row := rdb.Phone{ID: id, AccountID: accountID, Phone: phone, CreatedAt: now, UpdatedAt: now}
	if err := s.db.WithContext(ctx).Create(&row).Error; err != nil {
		return "", errx.Conflict("PHONE_TAKEN", "phone already registered")
	}
	return id.String(), nil
}

func (s *Service) SendPhoneVerificationCode(ctx context.Context, accountID, phoneID uuid.UUID) error {
	var row rdb.Phone
	if err := s.db.WithContext(ctx).Where("id = ? AND account_id = ? AND deleted_at IS NULL", phoneID, accountID).First(&row).Error; err != nil {
		return ErrNotFound
	}
	s.StoreVerificationCode(ctx, "phone:"+row.Phone, RandomCode())
	return nil
}

func (s *Service) VerifyPhone(ctx context.Context, accountID, phoneID uuid.UUID, code string) error {
	var row rdb.Phone
	if err := s.db.WithContext(ctx).Where("id = ? AND account_id = ? AND deleted_at IS NULL", phoneID, accountID).First(&row).Error; err != nil {
		return ErrNotFound
	}
	if !s.checkVerificationCode(ctx, "phone:"+row.Phone, code) {
		return errx.InvalidArg("INVALID_VERIFICATION_CODE", "invalid verification code")
	}
	now := time.Now()
	return s.db.WithContext(ctx).Model(&row).Updates(map[string]any{"verified_at": now, "updated_at": now}).Error
}

func (s *Service) UpdatePhone(ctx context.Context, accountID, phoneID uuid.UUID, phone string) error {
	phone = strings.TrimSpace(phone)
	res := s.db.WithContext(ctx).Model(&rdb.Phone{}).
		Where("id = ? AND account_id = ? AND deleted_at IS NULL", phoneID, accountID).
		Updates(map[string]any{"phone": phone, "verified_at": nil, "updated_at": time.Now()})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

func (s *Service) SetPrimaryPhone(ctx context.Context, accountID, phoneID uuid.UUID) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&rdb.Phone{}).Where("account_id = ? AND deleted_at IS NULL", accountID).Update("is_primary", false).Error; err != nil {
			return err
		}
		res := tx.Model(&rdb.Phone{}).Where("id = ? AND account_id = ? AND deleted_at IS NULL", phoneID, accountID).Update("is_primary", true)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return ErrNotFound
		}
		return nil
	})
}

func (s *Service) DeletePhone(ctx context.Context, accountID, phoneID uuid.UUID) error {
	now := time.Now()
	res := s.db.WithContext(ctx).Model(&rdb.Phone{}).
		Where("id = ? AND account_id = ? AND deleted_at IS NULL AND is_primary = false", phoneID, accountID).
		Update("deleted_at", now)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errx.FailedPrecond("PHONE_NOT_DELETABLE", "cannot delete primary or missing phone")
	}
	return nil
}

func (s *Service) AccountIDByEmail(ctx context.Context, email string) (uuid.UUID, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	var row rdb.Email
	if err := s.db.WithContext(ctx).Where("email = ? AND deleted_at IS NULL", email).First(&row).Error; err != nil {
		return uuid.Nil, ErrNotFound
	}
	return row.AccountID, nil
}

func (s *Service) AccountIDByPhone(ctx context.Context, phone string) (uuid.UUID, error) {
	phone = strings.TrimSpace(phone)
	var row rdb.Phone
	if err := s.db.WithContext(ctx).Where("phone = ? AND deleted_at IS NULL", phone).First(&row).Error; err != nil {
		return uuid.Nil, ErrNotFound
	}
	return row.AccountID, nil
}

func (s *Service) ForgerRoles(ctx context.Context, profileID uuid.UUID) ([]string, error) {
	var p rdb.ForgerProfile
	if err := s.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", profileID).First(&p).Error; err != nil {
		return nil, ErrNotFound
	}
	return []string(p.ForgerRoles), nil
}

func (s *Service) AuthorityRoles(ctx context.Context, profileID uuid.UUID) ([]string, error) {
	var p rdb.AuthorityProfile
	if err := s.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", profileID).First(&p).Error; err != nil {
		return nil, ErrNotFound
	}
	return []string(p.AuthorityRoles), nil
}

func (s *Service) InvalidateFull(ctx context.Context, accountID, profileID uuid.UUID) error {
	if s.redis == nil {
		return nil
	}
	return s.redis.Del(ctx, fullCacheKey(accountID, profileID)).Err()
}

func fullCacheKey(accountID, profileID uuid.UUID) string {
	return "nfxidentity:full:" + accountID.String() + ":" + profileID.String()
}

func (s *Service) ChangePassword(ctx context.Context, accountID uuid.UUID, current, next, code string) error {
	email, _ := s.primaryContacts(ctx, accountID)
	if email == "" || !s.checkVerificationCode(ctx, email, code) {
		return errx.InvalidArg("INVALID_VERIFICATION_CODE", "invalid verification code")
	}
	var ident rdb.Identity
	if err := s.db.WithContext(ctx).Where("account_id = ? AND identity_provider = ? AND deleted_at IS NULL", accountID, "password").First(&ident).Error; err != nil {
		return ErrNotFound
	}
	if ident.PasswordHash == nil || bcrypt.CompareHashAndPassword([]byte(*ident.PasswordHash), []byte(current)) != nil {
		return ErrInvalidCredentials
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(next), bcrypt.DefaultCost)
	if err != nil {
		return errx.Internal("HASH_FAILED", "failed to hash password")
	}
	h := string(hash)
	return s.db.WithContext(ctx).Model(&ident).Updates(map[string]any{"password_hash": h, "updated_at": time.Now()}).Error
}

func (s *Service) SendPasswordCode(ctx context.Context, accountID uuid.UUID) error {
	email, _ := s.primaryContacts(ctx, accountID)
	if email == "" {
		return ErrNotFound
	}
	s.StoreVerificationCode(ctx, email, RandomCode())
	return nil
}

func (s *Service) PatchProfile(ctx context.Context, accountID, profileID uuid.UUID, kind string, patch map[string]any) error {
	now := time.Now()
	patch["updated_at"] = now
	switch kind {
	case "forger":
		res := s.db.WithContext(ctx).Model(&rdb.ForgerProfile{}).
			Where("id = ? AND account_id = ? AND deleted_at IS NULL", profileID, accountID).
			Updates(patch)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return ErrProfileNotOwned
		}
	case "authority":
		res := s.db.WithContext(ctx).Model(&rdb.AuthorityProfile{}).
			Where("id = ? AND account_id = ? AND deleted_at IS NULL", profileID, accountID).
			Updates(patch)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return ErrProfileNotOwned
		}
	default:
		return errx.InvalidArg("INVALID_PROFILE_KIND", "kind must be forger or authority")
	}
	return nil
}

func (s *Service) PatchProfileSettings(ctx context.Context, profileID uuid.UUID, kind string, loginNotification *bool) error {
	if loginNotification == nil {
		return nil
	}
	now := time.Now()
	switch kind {
	case "forger":
		return s.db.WithContext(ctx).Model(&rdb.ForgerProfileSettings{}).
			Where("id = ?", profileID).
			Updates(map[string]any{"login_notification": *loginNotification, "updated_at": now}).Error
	case "authority":
		return s.db.WithContext(ctx).Model(&rdb.AuthorityProfileSettings{}).
			Where("id = ?", profileID).
			Updates(map[string]any{"login_notification": *loginNotification, "updated_at": now}).Error
	default:
		return errx.InvalidArg("INVALID_PROFILE_KIND", "kind must be forger or authority")
	}
}

func (s *Service) UpdatePreference(ctx context.Context, accountID, profileID uuid.UUID, kind, preference string) error {
	raw := datatypes.JSON([]byte(preference))
	if preference == "" {
		raw = datatypes.JSON([]byte("{}"))
	}
	switch kind {
	case "forger":
		res := s.db.WithContext(ctx).Model(&rdb.ForgerProfile{}).
			Where("id = ? AND account_id = ? AND deleted_at IS NULL", profileID, accountID).
			Update("preference", raw)
		if res.RowsAffected == 0 {
			return ErrProfileNotOwned
		}
		return res.Error
	case "authority":
		res := s.db.WithContext(ctx).Model(&rdb.AuthorityProfile{}).
			Where("id = ? AND account_id = ? AND deleted_at IS NULL", profileID, accountID).
			Update("preference", raw)
		if res.RowsAffected == 0 {
			return ErrProfileNotOwned
		}
		return res.Error
	default:
		return errx.InvalidArg("INVALID_PROFILE_KIND", "kind must be forger or authority")
	}
}

func (s *Service) CreateForgerProfile(ctx context.Context, accountID uuid.UUID, displayName, lang string) (string, error) {
	now := time.Now()
	id := uuid.New()
	p := rdb.ForgerProfile{
		ID: id, AccountID: accountID, ForgerRoles: pq.StringArray{"forger"},
		ProfileLanguage: langOrDefault(lang), DisplayName: &displayName, CreatedAt: now, UpdatedAt: now,
	}
	if err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&p).Error; err != nil {
			return err
		}
		return tx.Create(&rdb.ForgerProfileSettings{ProfileSettings: rdb.ProfileSettings{
			ID: id, LoginNotification: true, CreatedAt: now, UpdatedAt: now,
		}}).Error
	}); err != nil {
		return "", err
	}
	return id.String(), nil
}

func (s *Service) DeleteProfile(ctx context.Context, accountID, profileID uuid.UUID, kind string) error {
	now := time.Now()
	switch kind {
	case "forger":
		res := s.db.WithContext(ctx).Model(&rdb.ForgerProfile{}).
			Where("id = ? AND account_id = ? AND deleted_at IS NULL", profileID, accountID).
			Update("deleted_at", now)
		if res.RowsAffected == 0 {
			return ErrProfileNotOwned
		}
		return res.Error
	case "authority":
		res := s.db.WithContext(ctx).Model(&rdb.AuthorityProfile{}).
			Where("id = ? AND account_id = ? AND deleted_at IS NULL", profileID, accountID).
			Update("deleted_at", now)
		if res.RowsAffected == 0 {
			return ErrProfileNotOwned
		}
		return res.Error
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
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if kind == "authority" {
			if err := tx.Where("profile_id = ?", profileID).Delete(&rdb.AuthorityProfileBackground{}).Error; err != nil {
				return err
			}
			for _, img := range images {
				iid, err := uuid.Parse(img.ImageID)
				if err != nil {
					return errx.InvalidArg("INVALID_IMAGE_ID", "invalid image id")
				}
				row := rdb.AuthorityProfileBackground{ProfileBackground: rdb.ProfileBackground{
					ID: uuid.New(), ProfileID: profileID, ImageID: iid, SortOrder: img.SortOrder, CreatedAt: now, UpdatedAt: now,
				}}
				if err := tx.Create(&row).Error; err != nil {
					return err
				}
			}
			return nil
		}
		if err := tx.Where("profile_id = ?", profileID).Delete(&rdb.ForgerProfileBackground{}).Error; err != nil {
			return err
		}
		for _, img := range images {
			iid, err := uuid.Parse(img.ImageID)
			if err != nil {
				return errx.InvalidArg("INVALID_IMAGE_ID", "invalid image id")
			}
			row := rdb.ForgerProfileBackground{ProfileBackground: rdb.ProfileBackground{
				ID: uuid.New(), ProfileID: profileID, ImageID: iid, SortOrder: img.SortOrder, CreatedAt: now, UpdatedAt: now,
			}}
			if err := tx.Create(&row).Error; err != nil {
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
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if kind == "authority" {
			_ = tx.Model(&rdb.AuthorityProfileAvatar{}).Where("profile_id = ? AND is_active = true", profileID).Update("is_active", false).Error
			return tx.Create(&rdb.AuthorityProfileAvatar{ProfileAvatar: rdb.ProfileAvatar{
				ID: uuid.New(), ProfileID: profileID, ImageID: iid, IsActive: true, CreatedAt: now, UpdatedAt: now,
			}}).Error
		}
		_ = tx.Model(&rdb.ForgerProfileAvatar{}).Where("profile_id = ? AND is_active = true", profileID).Update("is_active", false).Error
		return tx.Create(&rdb.ForgerProfileAvatar{ProfileAvatar: rdb.ProfileAvatar{
			ID: uuid.New(), ProfileID: profileID, ImageID: iid, IsActive: true, CreatedAt: now, UpdatedAt: now,
		}}).Error
	})
}

func (s *Service) ListAccountProfiles(ctx context.Context, accountID uuid.UUID, kind string) ([]map[string]any, error) {
	if kind == "authority" {
		var rows []rdb.AuthorityProfile
		if err := s.db.WithContext(ctx).Where("account_id = ? AND deleted_at IS NULL", accountID).Find(&rows).Error; err != nil {
			return nil, err
		}
		out := make([]map[string]any, 0, len(rows))
		for _, r := range rows {
			out = append(out, s.authorityItem(r))
		}
		return out, nil
	}
	var rows []rdb.ForgerProfile
	if err := s.db.WithContext(ctx).Where("account_id = ? AND deleted_at IS NULL", accountID).Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]map[string]any, 0, len(rows))
	for _, r := range rows {
		out = append(out, s.forgerItem(r))
	}
	return out, nil
}

func (s *Service) SearchProfiles(ctx context.Context, kind, query string, limit, offset int) ([]map[string]any, int64, error) {
	like := "%" + query + "%"
	if kind == "authority" {
		q := s.db.WithContext(ctx).Model(&rdb.AuthorityProfile{}).Where("deleted_at IS NULL")
		if query != "" {
			q = q.Where("display_name ILIKE ?", like)
		}
		var total int64
		if err := q.Count(&total).Error; err != nil {
			return nil, 0, err
		}
		var rows []rdb.AuthorityProfile
		if err := q.Order("created_at desc").Limit(limitOr(limit, 20)).Offset(offset).Find(&rows).Error; err != nil {
			return nil, 0, err
		}
		items := make([]map[string]any, 0, len(rows))
		for _, r := range rows {
			items = append(items, s.authorityItem(r))
		}
		return items, total, nil
	}
	q := s.db.WithContext(ctx).Model(&rdb.ForgerProfile{}).Where("deleted_at IS NULL")
	if query != "" {
		q = q.Where("display_name ILIKE ? OR city ILIKE ?", like, like)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []rdb.ForgerProfile
	if err := q.Order("created_at desc").Limit(limitOr(limit, 20)).Offset(offset).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	items := make([]map[string]any, 0, len(rows))
	for _, r := range rows {
		items = append(items, s.forgerItem(r))
	}
	return items, total, nil
}

func (s *Service) PublicProfileCard(ctx context.Context, profileID uuid.UUID) (map[string]any, error) {
	var p rdb.ForgerProfile
	if err := s.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", profileID).First(&p).Error; err != nil {
		var a rdb.AuthorityProfile
		if err := s.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", profileID).First(&a).Error; err != nil {
			return nil, ErrNotFound
		}
		return s.authorityItem(a), nil
	}
	return s.forgerItem(p), nil
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
	return s.db.WithContext(ctx).Model(&rdb.AuthorityProfile{}).
		Where("id = ? AND deleted_at IS NULL", profileID).
		Update("authority_roles", pq.StringArray(roles)).Error
}

func (s *Service) GetAccountByID(ctx context.Context, accountID uuid.UUID) (map[string]any, error) {
	return s.FullAccountWithProfile(ctx, accountID, uuid.Nil, "forger")
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
