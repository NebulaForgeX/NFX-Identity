package apps

import (
	"context"
	appResult "nfxid/modules/clients/application/apps/results"

	"github.com/google/uuid"
)

// GetApp 根据ID获取应用
func (s *Service) GetApp(ctx context.Context, appID uuid.UUID) (appResult.AppRO, error) {
	domainEntity, err := s.appRepo.Get.ByID(ctx, appID)
	if err != nil {
		return appResult.AppRO{}, err
	}
	return appResult.AppMapper(domainEntity), nil
}

// GetAppByAppID 根据AppID获取应用
func (s *Service) GetAppByAppID(ctx context.Context, appID string) (appResult.AppRO, error) {
	domainEntity, err := s.appRepo.Get.ByAppID(ctx, appID)
	if err != nil {
		return appResult.AppRO{}, err
	}
	return appResult.AppMapper(domainEntity), nil
}
