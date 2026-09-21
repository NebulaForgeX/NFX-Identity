package account

import (
	"context"
	"time"

	"github.com/google/uuid"

	"nfxidentity/constants"
	"nfxidentity/enums"
	accountinternal "nfxidentity/modules/auth/application/account/internal"
	"nfxidentity/modules/auth/domain/profile"
	"nfxidentity/pkgs/slicex"
	"nfxidentity/pkgs/transaction"
)

type ConfirmAuthorityProfileBackgroundItemInput struct {
	ImageID   uuid.UUID
	SortOrder int
}

type ConfirmAuthorityProfileBackgroundsInput struct {
	AccountID uuid.UUID
	ProfileID uuid.UUID
	Images    []ConfirmAuthorityProfileBackgroundItemInput
}

type ConfirmAuthorityProfileBackgroundsOutput struct {
	Backgrounds []*profile.AuthorityProfileBackground
}

type ConfirmCommunityProfileBackgroundItemInput struct {
	ImageID   uuid.UUID
	SortOrder int
}

type ConfirmCommunityProfileBackgroundsInput struct {
	AccountID uuid.UUID
	ProfileID uuid.UUID
	Images    []ConfirmCommunityProfileBackgroundItemInput
}

type ConfirmCommunityProfileBackgroundsOutput struct {
	Backgrounds []*profile.ForgerProfileBackground
}

func (s *Service) ConfirmAuthorityProfileBackgrounds(
	ctx context.Context,
	in ConfirmAuthorityProfileBackgroundsInput,
) (*ConfirmAuthorityProfileBackgroundsOutput, error) {
	items := slicex.TrimByMax(
		in.Images,
		constants.AuthProfileMaxBackgrounds,
		func(a, b ConfirmAuthorityProfileBackgroundItemInput) bool {
			if a.SortOrder != b.SortOrder {
				return a.SortOrder < b.SortOrder
			}
			return a.ImageID.String() < b.ImageID.String()
		},
	)
	if err := s.requireOwnedProfile(ctx, in.AccountID, in.ProfileID, enums.AuthProfileScopeAuthority); err != nil {
		return nil, err
	}
	now := time.Now()
	backgrounds := make([]*profile.AuthorityProfileBackground, 0, len(items))
	for _, img := range items {
		backgrounds = append(backgrounds, accountinternal.NewAuthorityBackground(in.ProfileID, img.ImageID, img.SortOrder, now))
	}
	err := s.tx.WithUoW(ctx, func(ctx context.Context, uow transaction.UoW) error {
		profileRepo := s.repoFactory.Profile(uow)
		imageRepo := s.repoFactory.Image(uow)
		if err := profileRepo.Delete.DeleteAuthorityBackgroundsByProfileID(ctx, in.ProfileID); err != nil {
			return err
		}
		for _, background := range backgrounds {
			if _, err := imageRepo.Get.ByID(ctx, background.ImageID()); err != nil {
				return err
			}
			if err := profileRepo.Create.NewAuthorityBackground(ctx, background); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &ConfirmAuthorityProfileBackgroundsOutput{Backgrounds: backgrounds}, nil
}

func (s *Service) ConfirmCommunityProfileBackgrounds(
	ctx context.Context,
	in ConfirmCommunityProfileBackgroundsInput,
) (*ConfirmCommunityProfileBackgroundsOutput, error) {
	items := slicex.TrimByMax(
		in.Images,
		constants.AuthProfileMaxBackgrounds,
		func(a, b ConfirmCommunityProfileBackgroundItemInput) bool {
			if a.SortOrder != b.SortOrder {
				return a.SortOrder < b.SortOrder
			}
			return a.ImageID.String() < b.ImageID.String()
		},
	)
	if err := s.requireOwnedProfile(ctx, in.AccountID, in.ProfileID, enums.AuthProfileScopeCommunity); err != nil {
		return nil, err
	}
	now := time.Now()
	backgrounds := make([]*profile.ForgerProfileBackground, 0, len(items))
	for _, img := range items {
		backgrounds = append(backgrounds, accountinternal.NewForgerBackground(in.ProfileID, img.ImageID, img.SortOrder, now))
	}
	err := s.tx.WithUoW(ctx, func(ctx context.Context, uow transaction.UoW) error {
		profileRepo := s.repoFactory.Profile(uow)
		imageRepo := s.repoFactory.Image(uow)
		if err := profileRepo.Delete.DeleteForgerBackgroundsByProfileID(ctx, in.ProfileID); err != nil {
			return err
		}
		for _, background := range backgrounds {
			if _, err := imageRepo.Get.ByID(ctx, background.ImageID()); err != nil {
				return err
			}
			if err := profileRepo.Create.NewForgerBackground(ctx, background); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &ConfirmCommunityProfileBackgroundsOutput{Backgrounds: backgrounds}, nil
}
