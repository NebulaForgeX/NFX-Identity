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

func (h *Handler) PrimaryEmailByCommunityProfileID(ctx context.Context, profileID uuid.UUID) (*accountQuery.PrimaryEmailVO, error) {
	var profile rdbviews.FullAccountInformationWithForgerProfileView
	if err := h.db.WithContext(ctx).
		Table(rdbviews.FullAccountInformationWithForgerProfileView{}.TableName()).
		Where(rdbviews.FullAccountInformationWithForgerProfileViewCols.ID+" = ?", profileID.String()).
		First(&profile).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, authErr.ErrForgerProfileNotFound
		}
		return nil, authErr.ErrAccountProfileGetFailed
	}
	if profile.AccountID == nil {
		return nil, authErr.ErrForgerProfileNotFound
	}

	var email rdbviews.FullAccountInformationWithForgerProfileEmailView
	if err := h.db.WithContext(ctx).
		Table(rdbviews.FullAccountInformationWithForgerProfileEmailView{}.TableName()).
		Where(rdbviews.FullAccountInformationWithForgerProfileEmailViewCols.AccountID+" = ?", profile.AccountID.String()).
		Order(strings.Join([]string{
			rdbviews.FullAccountInformationWithForgerProfileEmailViewCols.IsPrimary + " DESC",
			rdbviews.FullAccountInformationWithForgerProfileEmailViewCols.CreatedAt + " ASC",
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
