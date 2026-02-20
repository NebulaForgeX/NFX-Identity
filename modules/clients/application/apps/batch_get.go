package apps

import (
	"context"
	appResult "nfxid/modules/clients/application/apps/results"

	"github.com/google/uuid"
)

// BatchGetApps 根据 ID 列表批量获取应用
func (s *Service) BatchGetApps(ctx context.Context, ids []uuid.UUID) ([]appResult.AppRO, error) {
	domainEntities, err := s.appRepo.Get.ByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}
	results := make([]appResult.AppRO, len(domainEntities))
	for i, entity := range domainEntities {
		results[i] = appResult.AppMapper(entity)
	}
	return results, nil
}
