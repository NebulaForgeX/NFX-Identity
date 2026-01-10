package grants

import (
	"context"
	grantResult "nfxid/modules/access/application/grants/results"

	"github.com/google/uuid"
)

// GetGrant 根据ID获取授权
func (s *Service) GetGrant(ctx context.Context, grantID uuid.UUID) (grantResult.GrantRO, error) {
	domainEntity, err := s.grantRepo.Get.ByID(ctx, grantID)
	if err != nil {
		return grantResult.GrantRO{}, err
	}
	return grantResult.GrantMapper(domainEntity), nil
}
