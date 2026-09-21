package single

import (
	"context"
	"errors"
	"strings"

	authErr "nfxidentity/errors/src/auth"
	"nfxidentity/modules/auth/infrastructure/query/account/mapper"
	rdbviews "nfxidentity/modules/auth/infrastructure/rdb/views"
	accountQuery "nfxidentity/modules/auth/query/account"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

// ByIDWithAuthorityProfile 返回「当前登录态」单 profile 读模型（FullAccountInformationWithAuthorityProfileVO）：
// 按 accountID 取账号/邮箱/电话，按 token 的 profileID 取单个 profile 及其 avatar/background
// （profileID 匹配不到时回退到该账号第一个 profile）。读 FullAccountInformationWithAuthorityProfile* view。
func (h *Handler) ByIDWithAuthorityProfile(
	ctx context.Context,
	accountID, profileID uuid.UUID,
) (*accountQuery.FullAccountInformationWithAuthorityProfileVO, error) {
	var account rdbviews.FullAccountInformationWithAuthorityProfileAccountView
	if err := h.db.WithContext(ctx).
		Table(rdbviews.FullAccountInformationWithAuthorityProfileAccountView{}.TableName()).
		Where(rdbviews.FullAccountInformationWithAuthorityProfileAccountViewCols.AccountID+" = ?", accountID.String()).
		First(&account).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, authErr.ErrAccountNotFound
		}
		return nil, authErr.ErrAccountBaseGetFailed
	}

	g, gctx := errgroup.WithContext(ctx)

	var emails []rdbviews.FullAccountInformationWithAuthorityProfileEmailView
	var phones []rdbviews.FullAccountInformationWithAuthorityProfilePhoneView
	var profileRows []rdbviews.FullAccountInformationWithAuthorityProfileView

	g.Go(func() error {
		if err := h.db.WithContext(gctx).
			Table(rdbviews.FullAccountInformationWithAuthorityProfileEmailView{}.TableName()).
			Where(rdbviews.FullAccountInformationWithAuthorityProfileEmailViewCols.AccountID+" = ?", accountID.String()).
			Order(strings.Join([]string{
				rdbviews.FullAccountInformationWithAuthorityProfileEmailViewCols.IsPrimary + " DESC",
				rdbviews.FullAccountInformationWithAuthorityProfileEmailViewCols.CreatedAt + " ASC",
			}, ", ")).
			Find(&emails).Error; err != nil {
			return authErr.ErrAccountEmailsGetFailed
		}
		return nil
	})

	g.Go(func() error {
		if err := h.db.WithContext(gctx).
			Table(rdbviews.FullAccountInformationWithAuthorityProfilePhoneView{}.TableName()).
			Where(rdbviews.FullAccountInformationWithAuthorityProfilePhoneViewCols.AccountID+" = ?", accountID.String()).
			Order(strings.Join([]string{
				rdbviews.FullAccountInformationWithAuthorityProfilePhoneViewCols.IsPrimary + " DESC",
				rdbviews.FullAccountInformationWithAuthorityProfilePhoneViewCols.CreatedAt + " ASC",
			}, ", ")).
			Find(&phones).Error; err != nil {
			return authErr.ErrAccountPhonesGetFailed
		}
		return nil
	})

	g.Go(func() error {
		if err := h.db.WithContext(gctx).
			Table(rdbviews.FullAccountInformationWithAuthorityProfileView{}.TableName()).
			Where(rdbviews.FullAccountInformationWithAuthorityProfileViewCols.AccountID+" = ?", accountID.String()).
			Order(rdbviews.FullAccountInformationWithAuthorityProfileViewCols.CreatedAt + " ASC").
			Find(&profileRows).Error; err != nil {
			return authErr.ErrAccountProfileGetFailed
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	emailVOs := make([]accountQuery.EmailVO, 0, len(emails))
	for i := range emails {
		emailVOs = append(emailVOs, mapper.FullAccountInformationWithAuthorityProfileEmailViewToVO(&emails[i]))
	}

	phoneVOs := make([]accountQuery.PhoneVO, 0, len(phones))
	for i := range phones {
		phoneVOs = append(phoneVOs, mapper.FullAccountInformationWithAuthorityProfilePhoneViewToVO(&phones[i]))
	}

	out := &accountQuery.FullAccountInformationWithAuthorityProfileVO{
		Account: mapper.FullAccountInformationWithAuthorityProfileAccountViewToVO(&account),
		Emails:  emailVOs,
		Phones:  phoneVOs,
	}

	if len(profileRows) == 0 {
		return out, nil
	}

	// 选中当前 profile：优先匹配 token 的 profileID，匹配不到回退第一个（按 created_at ASC）。
	selected := 0
	for i := range profileRows {
		if profileRows[i].ID != nil && *profileRows[i].ID == profileID {
			selected = i
			break
		}
	}

	profile := mapper.FullAccountInformationWithAuthorityProfileViewToVO(&profileRows[selected])
	profile.Avatars = []accountQuery.AuthorityProfileAvatarVO{}
	profile.Backgrounds = []accountQuery.AuthorityProfileBackgroundVO{}

	if profileRows[selected].ID != nil {
		selectedProfileID := *profileRows[selected].ID

		var avatars []rdbviews.FullAccountInformationWithAuthorityProfileAvatarView
		var backgrounds []rdbviews.FullAccountInformationWithAuthorityProfileBackgroundView
		var settingsRow rdbviews.FullAccountInformationWithAuthorityProfileSettingsView

		subG, subCtx := errgroup.WithContext(ctx)
		subG.Go(func() error {
			if err := h.db.WithContext(subCtx).
				Table(rdbviews.FullAccountInformationWithAuthorityProfileAvatarView{}.TableName()).
				Where(rdbviews.FullAccountInformationWithAuthorityProfileAvatarViewCols.ProfileID+" = ?", selectedProfileID.String()).
				Where(rdbviews.FullAccountInformationWithAuthorityProfileAvatarViewCols.IsActive+" = ?", true).
				Order(strings.Join([]string{
					rdbviews.FullAccountInformationWithAuthorityProfileAvatarViewCols.IsActive + " DESC",
					rdbviews.FullAccountInformationWithAuthorityProfileAvatarViewCols.CreatedAt + " ASC",
				}, ", ")).
				Find(&avatars).Error; err != nil {
				return authErr.ErrAccountProfileAvatarsGetFailed
			}
			return nil
		})
		subG.Go(func() error {
			if err := h.db.WithContext(subCtx).
				Table(rdbviews.FullAccountInformationWithAuthorityProfileBackgroundView{}.TableName()).
				Where(rdbviews.FullAccountInformationWithAuthorityProfileBackgroundViewCols.ProfileID+" = ?", selectedProfileID.String()).
				Order(strings.Join([]string{
					rdbviews.FullAccountInformationWithAuthorityProfileBackgroundViewCols.SortOrder + " ASC",
					rdbviews.FullAccountInformationWithAuthorityProfileBackgroundViewCols.CreatedAt + " ASC",
				}, ", ")).
				Find(&backgrounds).Error; err != nil {
				return authErr.ErrAccountProfileBackgroundsGetFailed
			}
			return nil
		})
		subG.Go(func() error {
			if err := h.db.WithContext(subCtx).
				Table(rdbviews.FullAccountInformationWithAuthorityProfileSettingsView{}.TableName()).
				Where(rdbviews.FullAccountInformationWithAuthorityProfileSettingsViewCols.ID+" = ?", selectedProfileID.String()).
				First(&settingsRow).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return nil
				}
				return authErr.ErrAccountProfileSettingsGetFailed
			}
			return nil
		})
		if err := subG.Wait(); err != nil {
			return nil, err
		}

		for i := range avatars {
			profile.Avatars = append(profile.Avatars, mapper.FullAccountInformationWithAuthorityProfileAvatarViewToVO(&avatars[i]))
		}
		for i := range backgrounds {
			profile.Backgrounds = append(profile.Backgrounds, mapper.FullAccountInformationWithAuthorityProfileBackgroundViewToVO(&backgrounds[i]))
		}
		profile.Settings = mapper.FullAccountInformationWithAuthorityProfileSettingsViewToVO(&settingsRow)
	}

	out.AuthorityProfile = &profile
	return out, nil
}
