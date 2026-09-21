package account

import (
	"context"

	"github.com/google/uuid"

	authErr "nfxidentity/errors/src/auth"
	profileQuery "nfxidentity/modules/auth/query/profile"
)

type GetCommunityPublicProfileCardInput struct {
	ProfileID uuid.UUID
}

func (s *Service) GetCommunityPublicProfileCard(ctx context.Context, in GetCommunityPublicProfileCardInput) (*profileQuery.ForgerProfileItemVO, error) {
	profiles, err := s.profiles.ForgerList.BatchGetByProfileIDs(ctx, []uuid.UUID{in.ProfileID})
	if err != nil {
		return nil, err
	}
	if len(profiles) == 0 {
		return nil, authErr.ErrForgerProfileNotFound
	}
	return &profiles[0], nil
}

type GetAuthorityPublicProfileCardInput struct {
	ProfileID uuid.UUID
}

func (s *Service) GetAuthorityPublicProfileCard(ctx context.Context, in GetAuthorityPublicProfileCardInput) (*profileQuery.AuthorityProfileItemVO, error) {
	profiles, err := s.profiles.AuthorityList.BatchGetByProfileIDs(ctx, []uuid.UUID{in.ProfileID})
	if err != nil {
		return nil, err
	}
	if len(profiles) == 0 {
		return nil, authErr.ErrAuthorityProfileNotFound
	}
	return &profiles[0], nil
}

type BatchGetCommunityPublicProfileCardsInput struct {
	ProfileIDs []uuid.UUID
}

type BatchGetCommunityPublicProfileCardsOutput struct {
	Profiles []profileQuery.ForgerProfileItemVO `json:"profiles"`
}

func (s *Service) BatchGetCommunityPublicProfileCards(ctx context.Context, in BatchGetCommunityPublicProfileCardsInput) (*BatchGetCommunityPublicProfileCardsOutput, error) {
	if len(in.ProfileIDs) == 0 {
		return &BatchGetCommunityPublicProfileCardsOutput{}, nil
	}
	profiles, err := s.profiles.ForgerList.BatchGetByProfileIDs(ctx, in.ProfileIDs)
	if err != nil {
		return nil, err
	}
	return &BatchGetCommunityPublicProfileCardsOutput{Profiles: profiles}, nil
}

type BatchGetAuthorityPublicProfileCardsInput struct {
	ProfileIDs []uuid.UUID
}

type BatchGetAuthorityPublicProfileCardsOutput struct {
	Profiles []profileQuery.AuthorityProfileItemVO `json:"profiles"`
}

func (s *Service) BatchGetAuthorityPublicProfileCards(ctx context.Context, in BatchGetAuthorityPublicProfileCardsInput) (*BatchGetAuthorityPublicProfileCardsOutput, error) {
	if len(in.ProfileIDs) == 0 {
		return &BatchGetAuthorityPublicProfileCardsOutput{}, nil
	}
	profiles, err := s.profiles.AuthorityList.BatchGetByProfileIDs(ctx, in.ProfileIDs)
	if err != nil {
		return nil, err
	}
	return &BatchGetAuthorityPublicProfileCardsOutput{Profiles: profiles}, nil
}
