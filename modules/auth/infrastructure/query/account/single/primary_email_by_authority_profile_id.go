package single

import (
	"context"
	"errors"
	"strings"

	authErr "nfxidentity/errors/src/auth"
	rdbviews "nfxidentity/modules/auth/infrastructure/rdb/views"
	accountQuery "nfxidentity/modules/auth/query/account"
	"nfxidentity/pkgs/ptrx"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *Handler) PrimaryEmailByAuthorityProfileID(ctx context.Context, profileID uuid.UUID) (*accountQuery.PrimaryEmailVO, error) {
	var profile rdbviews.FullAccountInformationWithAuthorityProfileView
	if err := h.db.WithContext(ctx).
		Table(rdbviews.FullAccountInformationWithAuthorityProfileView{}.TableName()).
		Where(rdbviews.FullAccountInformationWithAuthorityProfileViewCols.ID+" = ?", profileID.String()).
		First(&profile).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, authErr.ErrAuthorityProfileNotFound
		}
		return nil, authErr.ErrAccountProfileGetFailed
	}
	if profile.AccountID == nil {
		return nil, authErr.ErrAuthorityProfileNotFound
	}

	var email rdbviews.FullAccountInformationWithAuthorityProfileEmailView
	if err := h.db.WithContext(ctx).
		Table(rdbviews.FullAccountInformationWithAuthorityProfileEmailView{}.TableName()).
		Where(rdbviews.FullAccountInformationWithAuthorityProfileEmailViewCols.AccountID+" = ?", profile.AccountID.String()).
		Order(strings.Join([]string{
			rdbviews.FullAccountInformationWithAuthorityProfileEmailViewCols.IsPrimary + " DESC",
			rdbviews.FullAccountInformationWithAuthorityProfileEmailViewCols.CreatedAt + " ASC",
		}, ", ")).
		First(&email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, authErr.ErrEmailBindingNotFound
		}
		return nil, authErr.ErrAccountEmailsGetFailed
	}

	out := &accountQuery.PrimaryEmailVO{Email: ptrx.Deref(email.Email)}
	if profile.ProfileLanguage != nil {
		out.ProfileLanguage = string(*profile.ProfileLanguage)
	}
	if profile.DisplayName != nil {
		out.DisplayName = profile.DisplayName
	}
	if out.Email == "" {
		return nil, authErr.ErrEmailBindingNotFound
	}
	return out, nil
}
