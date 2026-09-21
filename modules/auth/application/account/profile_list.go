package account

import (
	"context"

	"github.com/google/uuid"

	profileQuery "nfxidentity/modules/auth/query/profile"
	"nfxidentity/pkgs/httpx"
)

type ListAuthorityProfilesInput struct {
	AccountID uuid.UUID
}

type SearchAuthorityProfilesInput struct {
	Query profileQuery.ListQuery
}

func (s *Service) ListAuthorityProfiles(ctx context.Context, in ListAuthorityProfilesInput) (httpx.Page[profileQuery.AuthorityProfileItemVO], error) {
	items, err := s.profiles.AuthorityList.ByAccountID(ctx, in.AccountID)
	if err != nil {
		return httpx.Page[profileQuery.AuthorityProfileItemVO]{}, err
	}
	return httpx.NewPage(items, int64(len(items))), nil
}

func (s *Service) SearchAuthorityProfiles(
	ctx context.Context,
	in SearchAuthorityProfilesInput,
) (httpx.Page[profileQuery.AuthorityProfileItemVO], error) {
	in.Query.Normalize()
	return s.profiles.AuthorityList.Search(ctx, in.Query)
}

type ListCommunityProfilesInput struct {
	AccountID uuid.UUID
}

type SearchCommunityProfilesInput struct {
	Query profileQuery.ListQuery
}

func (s *Service) ListCommunityProfiles(ctx context.Context, in ListCommunityProfilesInput) (httpx.Page[profileQuery.ForgerProfileItemVO], error) {
	items, err := s.profiles.ForgerList.ByAccountID(ctx, in.AccountID)
	if err != nil {
		return httpx.Page[profileQuery.ForgerProfileItemVO]{}, err
	}
	return httpx.NewPage(items, int64(len(items))), nil
}

func (s *Service) SearchCommunityProfiles(
	ctx context.Context,
	in SearchCommunityProfilesInput,
) (httpx.Page[profileQuery.ForgerProfileItemVO], error) {
	in.Query.Normalize()
	return s.profiles.ForgerList.Search(ctx, in.Query)
}
