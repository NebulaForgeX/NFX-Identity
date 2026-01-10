package api_keys

import (
	"context"
	apiKeyResult "nfxid/modules/clients/application/api_keys/results"

	"github.com/google/uuid"
)

// GetAPIKey 根据ID获取API密钥
func (s *Service) GetAPIKey(ctx context.Context, apiKeyID uuid.UUID) (apiKeyResult.APIKeyRO, error) {
	domainEntity, err := s.apiKeyRepo.Get.ByID(ctx, apiKeyID)
	if err != nil {
		return apiKeyResult.APIKeyRO{}, err
	}
	return apiKeyResult.APIKeyMapper(domainEntity), nil
}

// GetAPIKeyByKeyID 根据KeyID获取API密钥
func (s *Service) GetAPIKeyByKeyID(ctx context.Context, keyID string) (apiKeyResult.APIKeyRO, error) {
	domainEntity, err := s.apiKeyRepo.Get.ByKeyID(ctx, keyID)
	if err != nil {
		return apiKeyResult.APIKeyRO{}, err
	}
	return apiKeyResult.APIKeyMapper(domainEntity), nil
}
