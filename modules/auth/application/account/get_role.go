package account

import (
	"context"

	"nfxidentity/enums"
	"nfxidentity/errors/src/sys"
	accountinternal "nfxidentity/modules/auth/application/account/internal"

	"github.com/google/uuid"
)

func (s *Service) ForgerRoles(ctx context.Context, profileID uuid.UUID) ([]string, error) {
	p, err := s.repoFactory.Profile(accountinternal.None()).Get.ForgerByProfileID(ctx, profileID)
	if err != nil {
		return nil, sys.ErrNotFound
	}
	return accountinternal.RoleStrings(p.ForgerRoles()), nil
}

func (s *Service) AuthorityRoles(ctx context.Context, profileID uuid.UUID) ([]string, error) {
	p, err := s.repoFactory.Profile(accountinternal.None()).Get.AuthorityByProfileID(ctx, profileID)
	if err != nil {
		return nil, sys.ErrNotFound
	}
	return accountinternal.RoleStrings(p.AuthorityRoles()), nil
}

func (s *Service) HasAuthorityRole(ctx context.Context, profileID, role string) (bool, error) {
	pid, err := uuid.Parse(profileID)
	if err != nil {
		return false, err
	}
	p, err := s.repoFactory.Profile(accountinternal.None()).Get.AuthorityByProfileID(ctx, pid)
	if err != nil {
		return false, nil
	}
	return p.HasRole(enums.AuthAuthorityRole(role)), nil
}

func (s *Service) HasForgerRole(ctx context.Context, profileID, role string) (bool, error) {
	pid, err := uuid.Parse(profileID)
	if err != nil {
		return false, err
	}
	p, err := s.repoFactory.Profile(accountinternal.None()).Get.ForgerByProfileID(ctx, pid)
	if err != nil {
		return false, nil
	}
	return p.HasRole(enums.AuthForgerRole(role)), nil
}
