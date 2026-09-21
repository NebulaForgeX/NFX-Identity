package account

import (
	"context"

	"github.com/google/uuid"

	"nfxidentity/enums"
	authErr "nfxidentity/errors/src/auth"
	accountinternal "nfxidentity/modules/auth/application/account/internal"
	profileQuery "nfxidentity/modules/auth/query/profile"
	"nfxidentity/pkgs/httpx"
	"nfxidentity/pkgs/query"
)

type ListOwnerForgerProfilesInput struct {
	AccountID uuid.UUID
	Query     string
	Limit     int
	Offset    int
}

func (s *Service) ListOwnerForgerProfiles(ctx context.Context, in ListOwnerForgerProfilesInput) (httpx.Page[profileQuery.ForgerProfileItemVO], error) {
	if err := s.RequireOwner(ctx, in.AccountID); err != nil {
		return httpx.Page[profileQuery.ForgerProfileItemVO]{}, err
	}
	q := profileQuery.ListQuery{DomainPagination: query.DomainPagination{Limit: in.Limit, Offset: in.Offset}}
	if in.Query != "" {
		qstr := in.Query
		q.Search = &qstr
	}
	return s.SearchCommunityProfiles(ctx, SearchCommunityProfilesInput{Query: q})
}

type ListOwnerAuthorityProfilesInput struct {
	AccountID uuid.UUID
	Query     string
	Limit     int
	Offset    int
}

func (s *Service) ListOwnerAuthorityProfiles(ctx context.Context, in ListOwnerAuthorityProfilesInput) (httpx.Page[profileQuery.AuthorityProfileItemVO], error) {
	if err := s.RequireOwner(ctx, in.AccountID); err != nil {
		return httpx.Page[profileQuery.AuthorityProfileItemVO]{}, err
	}
	q := profileQuery.ListQuery{DomainPagination: query.DomainPagination{Limit: in.Limit, Offset: in.Offset}}
	if in.Query != "" {
		qstr := in.Query
		q.Search = &qstr
	}
	return s.SearchAuthorityProfiles(ctx, SearchAuthorityProfilesInput{Query: q})
}

type UpdateAuthorityRolesInput struct {
	ActorAccountID uuid.UUID
	ProfileID      uuid.UUID
	Roles          []string
}

func (s *Service) UpdateAuthorityRoles(ctx context.Context, in UpdateAuthorityRolesInput) error {
	if err := s.RequireOwner(ctx, in.ActorAccountID); err != nil {
		return err
	}
	converted := make([]enums.AuthAuthorityRole, 0, len(in.Roles))
	for _, r := range in.Roles {
		if r == "owner" {
			return authErr.ErrOwnerRoleImmutable
		}
		converted = append(converted, enums.AuthAuthorityRole(r))
	}
	p, err := s.repoFactory.Profile(accountinternal.None()).Get.AuthorityByProfileID(ctx, in.ProfileID)
	if err != nil {
		return err
	}
	if err := p.SetAuthorityRoles(converted); err != nil {
		return err
	}
	return s.repoFactory.Profile(accountinternal.None()).Update.AuthorityGeneric(ctx, p)
}
