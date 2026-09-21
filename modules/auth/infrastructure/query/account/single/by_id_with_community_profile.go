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

// ByIDWithCommunityProfile 返回「当前登录态」单 profile 读模型（FullAccountInformationWithCommunityProfileVO）：
// 按 accountID 取账号/邮箱/电话，按 token 的 profileID 取单个 profile 及其 avatar/background
// （profileID 匹配不到时回退到该账号第一个 profile）。读 FullAccountInformationWithCommunityProfile* view。
func (h *Handler) ByIDWithCommunityProfile(
	ctx context.Context,
	accountID, profileID uuid.UUID,
) (*accountQuery.FullAccountInformationWithCommunityProfileVO, error) {
	var account rdbviews.FullAccountInformationWithForgerProfileAccountView
	if err := h.db.WithContext(ctx).
		Table(rdbviews.FullAccountInformationWithForgerProfileAccountView{}.TableName()).
		Where(rdbviews.FullAccountInformationWithForgerProfileAccountViewCols.AccountID+" = ?", accountID.String()).
		First(&account).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, authErr.ErrAccountNotFound
		}
		return nil, authErr.ErrAccountBaseGetFailed
	}

	g, gctx := errgroup.WithContext(ctx)

	var emails []rdbviews.FullAccountInformationWithForgerProfileEmailView
	var phones []rdbviews.FullAccountInformationWithForgerProfilePhoneView
	var profileRows []rdbviews.FullAccountInformationWithForgerProfileView

	g.Go(func() error {
		if err := h.db.WithContext(gctx).
			Table(rdbviews.FullAccountInformationWithForgerProfileEmailView{}.TableName()).
			Where(rdbviews.FullAccountInformationWithForgerProfileEmailViewCols.AccountID+" = ?", accountID.String()).
			Order(strings.Join([]string{
				rdbviews.FullAccountInformationWithForgerProfileEmailViewCols.IsPrimary + " DESC",
				rdbviews.FullAccountInformationWithForgerProfileEmailViewCols.CreatedAt + " ASC",
			}, ", ")).
			Find(&emails).Error; err != nil {
			return authErr.ErrAccountEmailsGetFailed
		}
		return nil
	})

	g.Go(func() error {
		if err := h.db.WithContext(gctx).
			Table(rdbviews.FullAccountInformationWithForgerProfilePhoneView{}.TableName()).
			Where(rdbviews.FullAccountInformationWithForgerProfilePhoneViewCols.AccountID+" = ?", accountID.String()).
			Order(strings.Join([]string{
				rdbviews.FullAccountInformationWithForgerProfilePhoneViewCols.IsPrimary + " DESC",
				rdbviews.FullAccountInformationWithForgerProfilePhoneViewCols.CreatedAt + " ASC",
			}, ", ")).
			Find(&phones).Error; err != nil {
			return authErr.ErrAccountPhonesGetFailed
		}
		return nil
	})

	g.Go(func() error {
		if err := h.db.WithContext(gctx).
			Table(rdbviews.FullAccountInformationWithForgerProfileView{}.TableName()).
			Where(rdbviews.FullAccountInformationWithForgerProfileViewCols.AccountID+" = ?", accountID.String()).
			Order(rdbviews.FullAccountInformationWithForgerProfileViewCols.CreatedAt + " ASC").
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
		emailVOs = append(emailVOs, mapper.FullAccountInformationWithForgerProfileEmailViewToVO(&emails[i]))
	}

	phoneVOs := make([]accountQuery.PhoneVO, 0, len(phones))
	for i := range phones {
		phoneVOs = append(phoneVOs, mapper.FullAccountInformationWithForgerProfilePhoneViewToVO(&phones[i]))
	}

	out := &accountQuery.FullAccountInformationWithCommunityProfileVO{
		Account: mapper.FullAccountInformationWithForgerProfileAccountViewToVO(&account),
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

	profile := mapper.FullAccountInformationWithForgerProfileViewToVO(&profileRows[selected])
	profile.Avatars = []accountQuery.CommunityProfileAvatarVO{}
	profile.Backgrounds = []accountQuery.CommunityProfileBackgroundVO{}

	if profileRows[selected].ID != nil {
		selectedProfileID := *profileRows[selected].ID

		var avatars []rdbviews.FullAccountInformationWithForgerProfileAvatarView
		var backgrounds []rdbviews.FullAccountInformationWithForgerProfileBackgroundView
		var settingsRow rdbviews.FullAccountInformationWithForgerProfileSettingsView

		subG, subCtx := errgroup.WithContext(ctx)
		subG.Go(func() error {
			if err := h.db.WithContext(subCtx).
				Table(rdbviews.FullAccountInformationWithForgerProfileAvatarView{}.TableName()).
				Where(rdbviews.FullAccountInformationWithForgerProfileAvatarViewCols.ProfileID+" = ?", selectedProfileID.String()).
				Where(rdbviews.FullAccountInformationWithForgerProfileAvatarViewCols.IsActive+" = ?", true).
				Order(strings.Join([]string{
					rdbviews.FullAccountInformationWithForgerProfileAvatarViewCols.IsActive + " DESC",
					rdbviews.FullAccountInformationWithForgerProfileAvatarViewCols.CreatedAt + " ASC",
				}, ", ")).
				Find(&avatars).Error; err != nil {
				return authErr.ErrAccountProfileAvatarsGetFailed
			}
			return nil
		})
		subG.Go(func() error {
			if err := h.db.WithContext(subCtx).
				Table(rdbviews.FullAccountInformationWithForgerProfileBackgroundView{}.TableName()).
				Where(rdbviews.FullAccountInformationWithForgerProfileBackgroundViewCols.ProfileID+" = ?", selectedProfileID.String()).
				Order(strings.Join([]string{
					rdbviews.FullAccountInformationWithForgerProfileBackgroundViewCols.SortOrder + " ASC",
					rdbviews.FullAccountInformationWithForgerProfileBackgroundViewCols.CreatedAt + " ASC",
				}, ", ")).
				Find(&backgrounds).Error; err != nil {
				return authErr.ErrAccountProfileBackgroundsGetFailed
			}
			return nil
		})
		subG.Go(func() error {
			if err := h.db.WithContext(subCtx).
				Table(rdbviews.FullAccountInformationWithForgerProfileSettingsView{}.TableName()).
				Where(rdbviews.FullAccountInformationWithForgerProfileSettingsViewCols.ID+" = ?", selectedProfileID.String()).
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
			profile.Avatars = append(profile.Avatars, mapper.FullAccountInformationWithForgerProfileAvatarViewToVO(&avatars[i]))
		}
		for i := range backgrounds {
			profile.Backgrounds = append(profile.Backgrounds, mapper.FullAccountInformationWithForgerProfileBackgroundViewToVO(&backgrounds[i]))
		}
		profile.Settings = mapper.FullAccountInformationWithForgerProfileSettingsViewToVO(&settingsRow)
	}

	out.CommunityProfile = &profile
	return out, nil
}
