package groups

import (
	"context"
	groupResult "nfxid/modules/tenants/application/groups/results"

	"github.com/google/uuid"
)

// GetGroup 根据ID获取组
func (s *Service) GetGroup(ctx context.Context, groupID uuid.UUID) (groupResult.GroupRO, error) {
	domainEntity, err := s.groupRepo.Get.ByID(ctx, groupID)
	if err != nil {
		return groupResult.GroupRO{}, err
	}
	return groupResult.GroupMapper(domainEntity), nil
}
