package account

import (
	"context"
	"time"

	"github.com/google/uuid"

	"nfxidentity/enums"
	authErr "nfxidentity/errors/src/auth"
	accountinternal "nfxidentity/modules/auth/application/account/internal"
	"nfxidentity/modules/auth/domain/profile"
	"nfxidentity/pkgs/transaction"
)

type ConfirmCommunityProfileAvatarInput struct {
	AccountID uuid.UUID
	ProfileID uuid.UUID
	ImageID   uuid.UUID
}

type ConfirmAuthorityProfileAvatarInput struct {
	AccountID uuid.UUID
	ProfileID uuid.UUID
	ImageID   uuid.UUID
}

type ClearCommunityProfileAvatarInput struct {
	AccountID uuid.UUID
	ProfileID uuid.UUID
}

type ClearAuthorityProfileAvatarInput struct {
	AccountID uuid.UUID
	ProfileID uuid.UUID
}

func (s *Service) ConfirmCommunityProfileAvatar(ctx context.Context, in ConfirmCommunityProfileAvatarInput) error {
	if err := s.requireOwnedProfile(ctx, in.AccountID, in.ProfileID, enums.AuthProfileScopeCommunity); err != nil {
		return err
	}
	if in.ImageID == uuid.Nil {
		return authErr.ErrInvalidImageID
	}
	now := time.Now()
	inactive := false
	return s.tx.WithUoW(ctx, func(ctx context.Context, uow transaction.UoW) error {
		imageRepo := s.repoFactory.Image(uow)
		if _, err := imageRepo.Get.ByID(ctx, in.ImageID); err != nil {
			return err
		}
		profileRepo := s.repoFactory.Profile(uow)
		avatars, err := profileRepo.Get.ListForgerAvatarsByProfileID(ctx, in.ProfileID)
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
		return profileRepo.Create.NewForgerAvatar(ctx, accountinternal.NewForgerAvatar(in.ProfileID, in.ImageID, now))
	})
}

func (s *Service) ConfirmAuthorityProfileAvatar(ctx context.Context, in ConfirmAuthorityProfileAvatarInput) error {
	if err := s.requireOwnedProfile(ctx, in.AccountID, in.ProfileID, enums.AuthProfileScopeAuthority); err != nil {
		return err
	}
	if in.ImageID == uuid.Nil {
		return authErr.ErrInvalidImageID
	}
	now := time.Now()
	inactive := false
	return s.tx.WithUoW(ctx, func(ctx context.Context, uow transaction.UoW) error {
		imageRepo := s.repoFactory.Image(uow)
		if _, err := imageRepo.Get.ByID(ctx, in.ImageID); err != nil {
			return err
		}
		profileRepo := s.repoFactory.Profile(uow)
		avatars, err := profileRepo.Get.ListAuthorityAvatarsByProfileID(ctx, in.ProfileID)
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
		return profileRepo.Create.NewAuthorityAvatar(ctx, accountinternal.NewAuthorityAvatar(in.ProfileID, in.ImageID, now))
	})
}

func (s *Service) ClearCommunityProfileAvatar(ctx context.Context, in ClearCommunityProfileAvatarInput) error {
	if err := s.requireOwnedProfile(ctx, in.AccountID, in.ProfileID, enums.AuthProfileScopeCommunity); err != nil {
		return err
	}
	return s.tx.WithUoW(ctx, func(ctx context.Context, uow transaction.UoW) error {
		return s.repoFactory.Profile(uow).Delete.DeleteForgerAvatarsByProfileID(ctx, in.ProfileID)
	})
}

func (s *Service) ClearAuthorityProfileAvatar(ctx context.Context, in ClearAuthorityProfileAvatarInput) error {
	if err := s.requireOwnedProfile(ctx, in.AccountID, in.ProfileID, enums.AuthProfileScopeAuthority); err != nil {
		return err
	}
	return s.tx.WithUoW(ctx, func(ctx context.Context, uow transaction.UoW) error {
		return s.repoFactory.Profile(uow).Delete.DeleteAuthorityAvatarsByProfileID(ctx, in.ProfileID)
	})
}
